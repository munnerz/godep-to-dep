package main

import (
	"go/build"
	"golang.org/x/tools/go/vcs"
)

// These definitions are taken from github.com/tools/godep as we cannot
// directly import them due to all files being in 'package main'

// Godeps describes what a package needs to be rebuilt reproducibly.
// It's the same information stored in file Godeps.
type Godeps struct {
	ImportPath   string
	GoVersion    string
	GodepVersion string
	Packages     []string `json:",omitempty"` // Arguments to save, if any.
	Deps         []Dependency
	isOldFile    bool
}

// A Dependency is a specific revision of a package.
type Dependency struct {
	ImportPath string
	Comment    string `json:",omitempty"` // Description of commit, if present.
	Rev        string // VCS-specific commit ID.

	// used by command save & update
	ws   string // workspace
	root string // import path to repo root
	dir  string // full path to package

	// used by command update
	matched bool // selected for update by command line
	pkg     *Package
	missing bool // packages is missing

	// used by command go
	vcs *VCS
}

// VCS represents a version control system.
type VCS struct {
	vcs *vcs.Cmd

	IdentifyCmd string
	DescribeCmd string
	DiffCmd     string
	ListCmd     string
	RootCmd     string

	// run in sandbox repos
	ExistsCmd string
}

// Package represents a Go package.
type Package struct {
	Dir        string
	Root       string
	ImportPath string
	Deps       []string
	Standard   bool
	Processed  bool

	GoFiles        []string
	CgoFiles       []string
	IgnoredGoFiles []string

	TestGoFiles  []string
	TestImports  []string
	XTestGoFiles []string
	XTestImports []string

	Error struct {
		Err string
	}

	// --- New stuff for now
	Imports      []string
	Dependencies []build.Package
}
