package spec

//Name:          dummy
//Version:      1.0.0
//Release:        1%{?dist}
//Summary:       Dummy package
//
//License:       MIT
//URL:           https://example.com
//Source0:       https://example.com/%{name}-%{version}.tar.gz
//
//BuildArch:
//BuildRequires:  ruby ruby-devel
//Requires:       ruby(abi) = 1.8
//# If this package is mainly a ruby library, it should provide
//# whatever people have to require in their ruby scripts to use the library
//# For example, if people use this lib with "require 'foo'", it should provide
//# ruby(foo)
//Provides:       ruby(LIBNAME)
//
//%description
//
//
//%prep
//%setup -q
//
//
//%build
//export CFLAGS="$RPM_OPT_FLAGS"
//
//
//%install
//rm -rf $RPM_BUILD_ROOT
//
//
//%check
//
//
//%files
//%{!?_licensedir:%global license %%doc}
//%license add-license-file-here
//%doc add-docs-here
//# For noarch packages: ruby_sitelib
//%{ruby_sitelib}/*
//# For arch-specific packages: ruby_sitearch
//%{ruby_sitearch}/*
//
//
//%changelog

//func TestGenerateSpecContent(t *testing.T) {
//    p := &RpmSpec{
//        Name:          "dummy",
//        Version:       "1.0.0",
//        Summary:       "Dummy package",
//        License:       "MIT",
//        URL:           "https://example.com",
//        BuildArch:     Noarch,
//        Source0:       "https://example.com/%{name}-%{version}.tar.gz",
//        BuildRequires: []string{"ruby", "ruby-devel"},
//        Requires:      []string{"ruby(abi) = 1.8"},
//        Provides:      []string{"ruby(LIBNAME)"},
//        Description:   "Dummy package",
//    }
//
//	expected := `
//Name: dummy
//Version: 1.0.0
//Summary: Dummy package
//License: MIT
//URL: https://example.com
//Source0: https://example.com/%{name}-%{version}.tar.gz
//BuildRequires: ruby ruby-devel
//Requires: ruby(abi) = 1.8
//Provides: ruby(LIBNAME)
//
//%description
//Dummy package
//
//%prep
//%setup -q
//
//%build
//export CFLAGS="$RPM_OPT_FLAGS"
//
//%install
//rm -rf $RPM_BUILD_ROOT
//
//%check
//
//%files
//%license add-license-file-here
//%doc add-docs-here
//%{ruby_sitelib}/*
//%{ruby_sitearch}/*
//
//%changelog
//`
//
//	if content != expected {
//		t.Fatalf("expected: %s\nactual: %s", expected, content)
//	}
//}
