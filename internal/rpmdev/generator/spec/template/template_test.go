package template

import (
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/rpmdev/generator/spec"
	"testing"
)

func TestTemplete(t *testing.T) {
	p := spec.NewRpmSpec()

	p.Name = "dummy"
	p.Summary = "Dummy package"
	p.URL = "https://example.com"
	p.Source0 = "https://example.com/%{name}-%{version}.tar.gz"
	p.BuildRequires = []string{"ruby", "ruby-devel"}
	p.Requires = []string{"ruby(abi) = 1.8"}
	p.Provides = []string{"ruby(LIBNAME)"}
	p.Description = "Dummy package"
	p.BuildRequires = []string{"ocaml >= 3.10.0", "ocaml-findlib-devel", "ocaml-ocamldoc", "chrpath"}
	p.Global = []string{"opt %(test -x %{_bindir}/ocamlopt && echo 1 || echo 0)", "debug_package %{nil}", "_use_internal_dependency_generator 0", "__find_requires /usr/lib/rpm/ocaml-find-requires.sh", "__find_provides /usr/lib/rpm/ocaml-find-provides.sh"}
	p.Define = []string{"libname %(echo %{name} | sed -e 's/^ocaml-//')"}

	content, err := treatTemplate(p)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("content: %s", content)
}
