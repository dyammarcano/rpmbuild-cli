package spec

import (
	"reflect"
)

// https://docs.fedoraproject.org/en-US/epel/epel-packaging/

const (
	Noarch BuildArch = "noarch"
	I386   BuildArch = "i386"
	X64    BuildArch = "x86_64"
	ARM    BuildArch = "arm"
	ARM64  BuildArch = "aarch64"
)

type (
	field struct {
		name     string
		_default string
		key      string
		//tag       bool
		omitEmpty bool
		required  bool
		index     []int
		typ       reflect.Type
	}

	BuildArch string

	RpmSpec struct {
		Name               string
		Version            string
		Release            string
		Summary            string
		License            string
		URL                string
		Source0            string
		Group              string
		Epoch              string
		BuildArch          BuildArch
		BuildRoot          string
		BuildRequires      []string
		Requires           []string
		Prefixes           []string
		Provides           []string
		Obsoletes          []string
		Conflicts          []string
		Supplements        []string
		Recommends         []string
		Suggests           []string
		Enhances           []string
		Description        string
		Prep               []string
		Build              []string
		Install            []string
		Check              []string
		Files              []string
		Changelog          []string
		PrepScriptlet      []string
		BuildScriptlet     []string
		InstallScriptlet   []string
		Bcond              []string
		Pre                []string
		Post               []string
		PostRun            []string
		Clean              []string
		Bindir             string
		BuildArchitectures string
		BuildDir           string
		Docdir             string
		Doc                string
		Attr               []string
		LibDir             string
		Setup              string
		AutoSetup          string
		Configure          []string
		IncludeDir         string
		Global             []string
		Define             []string
		PrepScriptlets     []string
		Distribution       string

		//%{?_smp_mflags}
		//%{?_smp_flags}
		//%{?_smp_ncpus}
		//%{?_smp_build_cpus}
		//%{?_smp_strflags}
		//%{?_smp_strflags_nodist}
		//%{?ldconfig_scriptlets}
		//%{!?_smp_mflags: %global _smp_mflags %{?_smp_mflags_nodist}}
	}
)

//TODO

func NewRpmSpec() *RpmSpec {
	return &RpmSpec{
		Release:   "1%{?dist}",
		BuildArch: Noarch,
		License:   "MIT",
		Version:   "0.0.1",
	}
}

//
//func NewRpmFromSpec(specFilepath string) (*RpmSpec, error) {
//	return nil, nil
//}
//
//func NewRpmFromJSON(jsonFilepath string) (*RpmSpec, error) {
//	content, err := os.ReadFile(jsonFilepath)
//	if err != nil {
//		return nil, errors.WithStack(err)
//	}
//
//	payload := &RpmSpec{}
//	if err := json.Unmarshal(content, payload); err != nil {
//		return nil, errors.WithStack(err)
//	}
//
//	return nil, nil
//}

//
//func (p *RpmSpec) GenerateSpecContent() (string, error) {
//
//	// using reflection to generate spec file
//	valueOf, typeOf, err := util.ReflectStruct(p)
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//
//	contentBuilder := &strings.Builder{}
//
//	for i := 0; i < typeOf.NumField(); i++ {
//		tag := typeOf.Field(i).Tag.Get("spec")
//		if tag == "-" {
//			continue
//		}
//
//		name, opts := parseTag(tag)
//		if !isValidTag(name) {
//			name = ""
//		}
//
//		field := field{
//			name:     name,
//			_default: typeOf.Field(i).Tag.Get("default"),
//			key:      typeOf.Field(i).Tag.Get("key"),
//			//tag:       opts.Contains("tag"),
//			omitEmpty: opts.Contains("omitempty"),
//			required:  opts.Contains("required"),
//			index:     typeOf.Field(i).Index,
//			typ:       typeOf.Field(i).Type,
//		}
//
//		if field.name == "" {
//			continue
//		}
//
//		v := valueOf.Field(i).Interface()
//
//		if field.name == "Sources" {
//			if len(p.Sources) == 1 {
//				insertRow(contentBuilder, "Source0", p.Sources[0])
//			} else {
//				insertSources(contentBuilder, p.Sources)
//			}
//			continue
//		}
//
//		if field.omitEmpty {
//			continue
//		}
//
//		switch v.(type) {
//		case string:
//			if v.(string) == "" {
//				if field.required {
//					return "", errors.Errorf("field %s is required", field.name)
//				}
//
//				if field._default != "" {
//					v = field._default
//				}
//			}
//		}
//
//		if field.key != "" {
//			insertRow(contentBuilder, field.key, v.(string))
//		}
//
//		if field.name != "" {
//			insertRow(contentBuilder, field.name, v.(string))
//		}
//	}
//
//	return contentBuilder.String(), nil
//}

//func (p *RpmSpec) AddName(name string) {
//	p.Name = name
//}
//
//func (p *RpmSpec) AddVersion(version string) {
//	p.Version = version
//}
//
//func (p *RpmSpec) AddSummary(summary string) {
//	p.Summary = summary
//}
//
//func (p *RpmSpec) AddLicense(license string) {
//	p.License = license
//}
//
//func (p *RpmSpec) AddURL(url string) {
//	p.URL = url
//}
//
//func (p *RpmSpec) AddSource(source string) {
//	p.Source0 = source
//}
//
//func (p *RpmSpec) AddBuildRequires(buildRequires []string) {
//	p.BuildRequires = buildRequires
//}
//
//func (p *RpmSpec) AddRequires(requires []string) {
//	p.Requires = requires
//}
//
//func (p *RpmSpec) AddProvides(provides []string) {
//	p.Provides = provides
//}
//
//func (p *RpmSpec) AddConflicts(conflicts []string) {
//	p.Conflicts = conflicts
//}
//
//func (p *RpmSpec) AddDescription(description string) {
//	p.Description = description
//}
//
//func (p *RpmSpec) AddPrep(prep []string) {
//	p.Prep = prep
//}
//
//func (p *RpmSpec) AddBuild(build []string) {
//	p.Build = build
//}
//
//func (p *RpmSpec) AddInstall(install []string) {
//	p.Install = install
//}
//
//func (p *RpmSpec) AddCheck(check []string) {
//	p.Check = check
//}
//
//func insertRow(contentBuilder *strings.Builder, key string, value string) {
//	contentBuilder.WriteString(fmt.Sprintf("%s:\t\t%s\n", key, value))
//}
//
//func insertSources(contentBuilder *strings.Builder, values []string) {
//	cb := strings.Builder{}
//	item := 0
//
//	for _, v := range values {
//		cb.WriteString(fmt.Sprintf("Source%d:\t\t%s\n", item, v))
//		item++
//	}
//
//	contentBuilder.WriteString(cb.String())
//}
