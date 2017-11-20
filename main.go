package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/golang/dep"
	"github.com/golang/dep/gps"
	"github.com/golang/glog"
)

var (
	inputFile = flag.String("input", "", ``+
		`Input Godeps.json file to be converted into a Gopkg.toml`)
	outputFile = flag.String("output", "Gopkg.toml", ``+
		`The name of the output Gopkg.toml file`)
	releaseMapping = flag.String("rewrite-map", "", ``+
		`A comma separated list of repo=branchName mappings. This can be used `+
		`to specify custom branch overrides for a particular repository.`)
)

func main() {
	flag.Parse()
	var input Godeps

	// parse the release rewrite map
	releaseMap, err := parseReleaseMap(*releaseMapping)
	if err != nil {
		glog.Fatalf("Error parsing release rewrite map: %s", err.Error())
	}

	// read the Godeps.json file
	data, err := ioutil.ReadFile(*inputFile)
	if err != nil {
		glog.Fatalf("Error reading input Godeps file %q: %s", *inputFile, err)
	}

	// decode the Godeps.json file
	err = json.Unmarshal(data, &input)
	if err != nil {
		glog.Fatalf("Error unmarshaling Godeps file %q: %s", *inputFile, err)
	}

	// construct a dep source manager in order to resolve package names to
	// their project roots
	sourceMgr, err := gps.NewSourceManager(gps.SourceManagerConfig{
		Cachedir: filepath.Join(defaultGOPATH(), "pkg", "dep"),
	})
	if err != nil {
		glog.Fatalf("Error constructing Dep source manager: %s", err)
	}

	// we flatten dependencies here as Dep requires that only project roots are
	// specified in the Gopkg.toml file
	flattenedDeps, err := flattenDepsToRoot(sourceMgr, input.Deps)
	if err != nil {
		glog.Fatalf("Error flattening project dependencies: %s", err)
	}

	manifest := dep.Manifest{
		Constraints: rewriteDepsWithPrefix(flattenedDeps, releaseMap),
	}
	tomlBytes, err := manifest.MarshalTOML()
	if err != nil {
		glog.Fatalf("Error marshaling Dep manifest to TOML: %s", err)
	}
	fmt.Print(string(tomlBytes))
}

func parseReleaseMap(input string) (map[string]string, error) {
	releaseMap := make(map[string]string)
	entries := strings.Split(input, ",")
	for _, e := range entries {
		branchConstraint := strings.Split(e, "=")
		if len(branchConstraint) != 2 {
			return nil, fmt.Errorf("invalid release branch constraint: %q", e)
		}
		project := branchConstraint[0]
		branch := branchConstraint[1]
		releaseMap[project] = branch
	}
	return releaseMap, nil
}

// flattenDepsToRoot will 'flatten' all dependencies given in deps down to the
// root packages they are contained within.
// For example, it will convert 'k8s.io/api/core/v1' and 'k8s.io/api/extensions'
// into 'k8s.io/api'.
func flattenDepsToRoot(manager gps.SourceManager, deps []Dependency) (map[string]string, error) {
	depMap := make(map[string]string)
	for _, d := range deps {
		root, err := manager.DeduceProjectRoot(d.ImportPath)
		if err != nil {
			return nil, err
		}
		depMap[string(root)] = d.Rev
	}
	return depMap, nil
}

func rewriteDepsWithPrefix(deps map[string]string, releaseMap map[string]string) gps.ProjectConstraints {
	constraints := make(gps.ProjectConstraints)
Outer:
	for pkg, rev := range deps {
		// search for a release rewrite rule for this package. If one exists,
		// use the specified branch as the pinned dependency
		for prefix, branch := range releaseMap {
			if strings.HasPrefix(pkg, prefix) {
				constraints[gps.ProjectRoot(pkg)] = gps.ProjectProperties{
					Constraint: gps.NewBranch(branch),
				}
				continue Outer
			}
		}

		// if there is release rewrite rule for this repository, we pin to the
		// revision specified in the Godeps.json file
		constraints[gps.ProjectRoot(pkg)] = gps.ProjectProperties{
			Constraint: gps.Revision(rev),
		}
	}
	return constraints
}

// defaultGOPATH gets the default GOPATH that was added in 1.8
// copied from go/build/build.go
func defaultGOPATH() string {
	env := "HOME"
	if runtime.GOOS == "windows" {
		env = "USERPROFILE"
	} else if runtime.GOOS == "plan9" {
		env = "home"
	}
	if home := os.Getenv(env); home != "" {
		def := filepath.Join(home, "go")
		if def == runtime.GOROOT() {
			// Don't set the default GOPATH to GOROOT,
			// as that will trigger warnings from the go tool.
			return ""
		}
		return def
	}
	return ""
}
