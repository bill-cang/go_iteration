/*
@Time   : 2021/12/20 14:17
@Author : ckx0709
@Remark :
*/
package unit1

import (
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/packages/packagestest"
	"testing"
)

func TestClos0(t *testing.T) {
	Clos0()
}

func TestClos1(t *testing.T) {
	Clos1()
}

func TestClos2(t *testing.T) {
	Clos2()
}

func TestClos02(t *testing.T) {
	t.Parallel()

	exported := packagestest.Export(t, packagestest.GOPATH, []packagestest.Module{{
		Name: "golang.org/fake",
		Files: map[string]interface{}{
			"vendor/vendor.com/foo/foo.go": `package foo; const X = "hi"`,
			"user/user.go":                 `package user`,
		},
	}})
	defer exported.Cleanup()

	exported.Config.Mode = packages.LoadAllSyntax
	exported.Config.Logf = t.Logf
	exported.Config.Overlay = map[string][]byte{
		exported.File("golang.org/fake", "user/user.go"): []byte(`package user; import "vendor.com/foo"; var x = foo.X`),
	}
	initial, err := packages.Load(exported.Config, "golang.org/fake/user")
	if err != nil {
		t.Fatal(err)
	}
	user := initial[0]
	if len(user.Imports) != 1 {
		t.Fatal("no imports for user")
	}
	if user.Imports["vendor.com/foo"].Name != "foo" {
		t.Errorf("failed to load vendored package foo, imports: %#v", user.Imports["vendor.com/foo"])
	}
}
