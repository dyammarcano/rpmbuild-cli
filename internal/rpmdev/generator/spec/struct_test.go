package spec

import (
	"fmt"
	"testing"
)

func TestTemplete(t *testing.T) {
	p := NewRpmSpec()

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
	p.Define = []string{"libname %(echo %{k} | sed -e 's/^ocaml-//')"}
	p.Install = []string{"mkdir -p %{buildroot}/opt", "touch %{buildroot}/opt/file"}
	p.Prefix = "/opt"
	p.Files = []string{"/opt/file", "/opt/file2"}
	p.Pre = []string{"%setup -q -n %{name}-%{version}"}

	content, err := p.GenerateTemplate()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("content: %s", content)
}
