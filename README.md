# godep-to-dep

This is a simple tool for converting a Godeps.json file into a Gopkg.toml.

It was primarily designed for use in the Kubernetes project, when exporting
projects such as k8s.io/apimachinery, k8s.io/api, k8s.io/apiserver etc. from
the main monorepo.

It rewrites all dependencies specified in the Godeps.json into revision
pinned constraints in the Gopkg.toml file. Using the `-rewrite-map` argument,
it is possible to force certain repositories to be pinned to particular
branches too.

An example invocation:

```bash
➜  client-go git:(master) godep-to-dep -input Godeps/Godeps.json -rewrite-map k8s.io/api=release-1.8,k8s.io/apimachinery=release-1.8,k8s.io/client-go=release-5.0

[[constraint]]
  name = "cloud.google.com/go"
  revision = "3b1ae45394a234c385be014e9a488f2bb6eef821"

[[constraint]]
  name = "github.com/Azure/go-autorest"
  revision = "58f6f26e200fa5dfb40c9cd1c83f3e2c860d779d"

[[constraint]]
  name = "github.com/PuerkitoBio/purell"
  revision = "8a290539e2e8629dbc4e6bad948158f790ec31f4"

[[constraint]]
  name = "github.com/PuerkitoBio/urlesc"
  revision = "5bd2802263f21d8788851d5305584c82a5c75d7e"

[[constraint]]
  name = "github.com/coreos/go-oidc"
  revision = "a4973d9a4225417aecf5d450a9522f00c1f7130f"

[[constraint]]
  name = "github.com/coreos/pkg"
  revision = "fa29b1d70f0beaddd4c7021607cc3c3be8ce94b8"

[[constraint]]
  name = "github.com/davecgh/go-spew"
  revision = "782f4967f2dc4564575ca782fe2d04090b5faca8"

[[constraint]]
  name = "github.com/dgrijalva/jwt-go"
  revision = "01aeca54ebda6e0fbfafd0a524d234159c05ec20"

[[constraint]]
  name = "github.com/docker/spdystream"
  revision = "449fdfce4d962303d702fec724ef0ad181c92528"

[[constraint]]
  name = "github.com/emicklei/go-restful"
  revision = "ff4f55a206334ef123e4f79bbf348980da81ca46"

[[constraint]]
  name = "github.com/ghodss/yaml"
  revision = "73d445a93680fa1a78ae23a5839bad48f32ba1ee"

[[constraint]]
  name = "github.com/go-openapi/jsonpointer"
  revision = "46af16f9f7b149af66e5d1bd010e3574dc06de98"

[[constraint]]
  name = "github.com/go-openapi/jsonreference"
  revision = "13c6e3589ad90f49bd3e3bbe2c2cb3d7a4142272"

[[constraint]]
  name = "github.com/go-openapi/spec"
  revision = "7abd5745472fff5eb3685386d5fb8bf38683154d"

[[constraint]]
  name = "github.com/go-openapi/swag"
  revision = "f3f9494671f93fcff853e3c6e9e948b3eb71e590"

[[constraint]]
  name = "github.com/gogo/protobuf"
  revision = "c0656edd0d9eab7c66d1eb0c568f9039345796f7"

[[constraint]]
  name = "github.com/golang/glog"
  revision = "44145f04b68cf362d9c4df2182967c2275eaefed"

[[constraint]]
  name = "github.com/golang/groupcache"
  revision = "02826c3e79038b59d737d3b1c0a1d937f71a4433"

[[constraint]]
  name = "github.com/golang/protobuf"
  revision = "4bd1920723d7b7c925de087aa32e2187708897f7"

[[constraint]]
  name = "github.com/google/btree"
  revision = "7d79101e329e5a3adf994758c578dab82b90c017"

[[constraint]]
  name = "github.com/google/gofuzz"
  revision = "44d81051d367757e1c7c6a5a86423ece9afcf63c"

[[constraint]]
  name = "github.com/googleapis/gnostic"
  revision = "0c5108395e2debce0d731cf0287ddf7242066aba"

[[constraint]]
  name = "github.com/gophercloud/gophercloud"
  revision = "443743e88335413103dcf1997e46d401b264fbcd"

[[constraint]]
  name = "github.com/gregjones/httpcache"
  revision = "787624de3eb7bd915c329cba748687a3b22666a6"

[[constraint]]
  name = "github.com/hashicorp/golang-lru"
  revision = "a0d98a5f288019575c6d1f4bb1573fef2d1fcdc4"

[[constraint]]
  name = "github.com/howeyc/gopass"
  revision = "bf9dde6d0d2c004a008c27aaee91170c786f6db8"

[[constraint]]
  name = "github.com/imdario/mergo"
  revision = "6633656539c1639d9d78127b7d47c622b5d7b6dc"

[[constraint]]
  name = "github.com/jonboulle/clockwork"
  revision = "72f9bd7c4e0c2a40055ab3d0f09654f730cce982"

[[constraint]]
  name = "github.com/json-iterator/go"
  revision = "36b14963da70d11297d313183d7e6388c8510e1e"

[[constraint]]
  name = "github.com/juju/ratelimit"
  revision = "5b9ff866471762aa2ab2dced63c9fb6f53921342"

[[constraint]]
  name = "github.com/mailru/easyjson"
  revision = "2f5df55504ebc322e4d52d34df6a1f5b503bf26d"

[[constraint]]
  name = "github.com/peterbourgon/diskv"
  revision = "5f041e8faa004a95c88a202771f4cc3e991971e6"

[[constraint]]
  name = "github.com/pmezard/go-difflib"
  revision = "d8ed2627bdf02c080bf22230dbb337003b7aba2d"

[[constraint]]
  name = "github.com/spf13/pflag"
  revision = "9ff6c6923cfffbcd502984b8e0c80539a94968b7"

[[constraint]]
  name = "github.com/stretchr/testify"
  revision = "f6abca593680b2315d2075e0f5e2a9751e3f431a"

[[constraint]]
  name = "golang.org/x/crypto"
  revision = "81e90905daefcd6fd217b62423c0908922eadb30"

[[constraint]]
  name = "golang.org/x/net"
  revision = "1c05540f6879653db88113bc4a2b70aec4bd491f"

[[constraint]]
  name = "golang.org/x/oauth2"
  revision = "a6bd8cefa1811bd24b86f8902872e4e8225f74c4"

[[constraint]]
  name = "golang.org/x/sys"
  revision = "7ddbeae9ae08c6a06a59597f0c9edbc5ff2444ce"

[[constraint]]
  name = "golang.org/x/text"
  revision = "b19bf474d317b857955b12035d2c5acb57ce8b01"

[[constraint]]
  name = "gopkg.in/inf.v0"
  revision = "3887ee99ecf07df5b447e9b00d9c0b2adaa9f3e4"

[[constraint]]
  name = "gopkg.in/yaml.v2"
  revision = "53feefa2559fb8dfa8d81baad31be332c97d6c77"

[[constraint]]
  branch = "release-1.8"
  name = "k8s.io/api"

[[constraint]]
  branch = "release-1.8"
  name = "k8s.io/apimachinery"

[[constraint]]
  name = "k8s.io/kube-openapi"
  revision = "61b46af70dfed79c6d24530cd23b41440a7f22a5"
```
