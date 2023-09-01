package spec

import (
	"bytes"
	"fmt"
	"github.com/dyammarcano/utils/orderedmap"
	"github.com/dyammarcano/utils/util"
	"reflect"
	"strings"
	"text/template"
)

// https://docs.fedoraproject.org/en-US/epel/epel-packaging/

const (
	Noarch BuildArch = "noarch"
	I386   BuildArch = "i386"
	X64    BuildArch = "x86_64"
	ARM    BuildArch = "arm"
	ARM64  BuildArch = "aarch64"

	Error                  RpmParseState = -1
	None                   RpmParseState = 0
	Preamble               RpmParseState = 11
	Prep                   RpmParseState = 12
	Build                  RpmParseState = 13
	Install                RpmParseState = 14
	Check                  RpmParseState = 15
	Clean                  RpmParseState = 16
	Files                  RpmParseState = 17
	Pre                    RpmParseState = 18
	Post                   RpmParseState = 19
	Preun                  RpmParseState = 20
	Postun                 RpmParseState = 21
	Pretrans               RpmParseState = 22
	Posttrans              RpmParseState = 23
	Description            RpmParseState = 24
	Changelog              RpmParseState = 25
	Triggerin              RpmParseState = 26
	Triggerun              RpmParseState = 27
	Verifyscript           RpmParseState = 28
	Buildarchitectures     RpmParseState = 29
	Triggerpostun          RpmParseState = 30
	Triggerprein           RpmParseState = 31
	Policies               RpmParseState = 32
	Filetriggerin          RpmParseState = 33
	Filetriggerun          RpmParseState = 34
	Filetriggerpostun      RpmParseState = 35
	Transfiletriggerin     RpmParseState = 36
	Transfiletriggerun     RpmParseState = 37
	Transfiletriggerpostun RpmParseState = 38
	Empty                  RpmParseState = 39
	Patchlist              RpmParseState = 40
	Sourcelist             RpmParseState = 41
	Buildrequires          RpmParseState = 42
	Conf                   RpmParseState = 43
	Last                   RpmParseState = 44
)

type (
	BuildArch     string
	RpmParseState int

	field struct {
		k string
		v reflect.Value
		t string
	}

	RpmSpec struct {
		Name               string    `spec:"Name"`
		Version            string    `spec:"Version"`
		Release            string    `spec:"Release"`
		Summary            string    `spec:"Summary"`
		License            string    `spec:"License"`
		URL                string    `spec:"URL"`
		Source0            string    `spec:"Source0"`
		Group              string    `spec:"Group"`
		Epoch              string    `spec:"Epoch"`
		BuildArch          BuildArch `spec:"BuildArch"`
		BuildRoot          string    `spec:"BuildRoot"`
		BuildRequires      []string  `spec:"BuildRequires"`
		Requires           []string  `spec:"Requires"`
		Prefix             string    `spec:"Prefix"`
		Provides           []string  `spec:"Provides"`
		Obsoletes          []string  `spec:"Obsoletes"`
		Conflicts          []string  `spec:"Conflicts"`
		Supplements        []string  `spec:"Supplements"`
		Recommends         []string  `spec:"Recommends"`
		Suggests           []string  `spec:"Suggests"`
		Enhances           []string  `spec:"Enhances"`
		Description        string    `spec:"%description"`
		Prep               []string  `spec:"%prep"`
		Build              []string  `spec:"%build"`
		Install            []string  `spec:"%install"`
		Check              []string  `spec:"%check"`
		Files              []string  `spec:"%files"`
		Changelog          []string  `spec:"%changelog"`
		PrepScriptlet      []string  `spec:"%prep"`
		BuildScriptlet     []string  `spec:"%build"`
		InstallScriptlet   []string  `spec:"%install"`
		Bcond              []string  `spec:"%bcond_with"`
		Pre                []string  `spec:"%pre"`
		Post               []string  `spec:"%post"`
		PostRun            []string  `spec:"%posttrans"`
		Clean              []string  `spec:"%clean"`
		Bindir             string    `spec:"%{_bindir}"`
		BuildArchitectures string    `spec:"%{?_build_architectures}"`
		BuildDir           string    `spec:"%{?_builddir}"`
		Docdir             string    `spec:"%{?_docdir_fmt}"`
		Doc                string    `spec:"%doc"`
		Attr               []string  `spec:"%attr"`
		LibDir             string    `spec:"%{_libdir}"`
		Setup              string    `spec:"%setup"`
		AutoSetup          string    `spec:"%autosetup"`
		Configure          []string  `spec:"%configure"`
		IncludeDir         string    `spec:"%{_includedir}"`
		Global             []string  `spec:"%global"`
		Define             []string  `spec:"%define"`
		PrepScriptlets     []string  `spec:"%prep_scriptlets"`
		Distribution       string    `spec:"%{?distribution}"`

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

func NewRpmSpec() *RpmSpec {
	return &RpmSpec{
		Release:   "1%{?dist}",
		BuildArch: Noarch,
		License:   "MIT",
		Version:   "0.0.1",
	}
}

func (rpmSpec *RpmSpec) GenerateTemplate() (string, error) {
	tmplStr := createTemplateString(rpmSpec)
	tmpl, err := template.New("spec").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var data bytes.Buffer

	if err = tmpl.Execute(&data, rpmSpec); err != nil {
		return "", err
	}

	return data.String(), nil
}

func createTemplateString(rpmSpec *RpmSpec) string {
	templateStrings := orderedmap.NewOrderedMap()

	templateStrings.Set("Global", "{{if .Global}}%global {{.}}{{end}}\n")
	templateStrings.Set("Define", "{{if .Define}}%define {{.}}{{end}}\n")
	templateStrings.Set("Name", "{{if .Name}}Name:           {{.Name}}{{end}}\n")
	templateStrings.Set("Version", "{{if .Version}}Version:        {{.Version}}{{end}}\n")
	templateStrings.Set("Release", "{{if .Release}}Release:        {{.Release}}{{end}}\n")
	templateStrings.Set("Summary", "{{if .Summary}}Summary:        {{.Summary}}{{end}}\n")
	templateStrings.Set("Group", "{{if .Group}}Group:          {{.Group}}{{end}}\n")
	templateStrings.Set("Distribution", "{{if .Distribution}}Distribution:   {{.Distribution}}{{end}}\n")
	templateStrings.Set("License", "{{if .License}}License:        {{.License}}{{end}}\n")
	templateStrings.Set("BuildArch", "{{if .BuildArch}}BuildArch:      {{.BuildArch}}{{end}}\n")
	templateStrings.Set("BuildRequires", "{{if .BuildRequires}}BuildRequires:  {{.BuildRequires}}{{end}}\n")
	templateStrings.Set("Epoch", "{{if .Epoch}}Epoch:          {{.Epoch}}{{end}}\n")
	templateStrings.Set("BuildRoot", "{{if .BuildRoot}}BuildRoot:      {{.BuildRoot}}{{end}}\n")
	templateStrings.Set("Prefix", "{{if .Prefix}}Prefix:         {{.Prefix}}{{end}}\n")
	templateStrings.Set("Provides", "{{if .Provides}}Provides:       {{.Provides}}{{end}}\n")
	templateStrings.Set("URL", "{{if .URL}}URL:            {{.URL}}{{end}}\n")
	templateStrings.Set("Source0", "{{if .Source0}}Source0:        {{.Source0}}{{end}}\n")
	templateStrings.Set("Requires", "{{if .Requires}}Requires:       {{.Requires}}{{end}}\n")
	templateStrings.Set("Obsoletes", "{{if .Obsoletes}}Obsoletes:      {{.Obsoletes}}{{end}}\n")
	templateStrings.Set("Conflicts", "{{if .Conflicts}}Conflicts:      {{.Conflicts}}{{end}}\n")
	templateStrings.Set("Supplements", "{{if .Supplements}}Supplements:    {{.Supplements}}{{end}}\n")
	templateStrings.Set("Recommends", "{{if .Recommends}}Recommends:     {{.Recommends}}{{end}}\n")
	templateStrings.Set("Suggests", "{{if .Suggests}}Suggests:       {{.Suggests}}{{end}}\n")
	templateStrings.Set("Enhances", "{{if .Enhances}}Enhances:       {{.Enhances}}{{end}}\n")
	templateStrings.Set("Description", "{{if .Description}}%description\n{{.Description}}\n{{end}}\n")
	templateStrings.Set("Prep", "{{if .Prep}}%prep\n{{.Prep}}\n{{end}}\n")
	templateStrings.Set("Build", "{{if .Build}}%build\n{{.Build}}\n{{end}}\n")
	templateStrings.Set("Install", "{{if .Install}}%install\n{{.Install}}\n{{end}}\n")
	templateStrings.Set("Check", "{{if .Check}}%check\n{{.Check}}\n{{end}}\n")
	templateStrings.Set("Files", "{{if .Files}}%files\n{{.Files}}\n{{end}}\n")
	templateStrings.Set("Changelog", "{{if .Changelog}}%changelog\n{{.Changelog}}\n{{end}}\n")

	valueOf, typeOf, err := util.ReflectStruct(rpmSpec)
	if err != nil {
		panic(err)
	}

	tmplStr := strings.Builder{}

	for pair := templateStrings.Oldest(); pair != nil; pair = pair.Next() {
		key := pair.Key.(string)
		value := valueOf.FieldByName(key)
		tp, ok := typeOf.FieldByName(key)
		if !ok {
			continue
		}

		fieldVal := field{
			k: key,
			v: value,
			t: tp.Tag.Get("spec"),
		}

		if isFieldSet(value) {
			if fieldVal.v.Kind() == reflect.Slice {
				fixItemList(&tmplStr, fieldVal)
				continue
			}

			tmplStr.WriteString(pair.Value.(string))
		}
	}

	return tmplStr.String()
}

func isFieldSet(value reflect.Value) bool {
	return !reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func fixItemList(sb *strings.Builder, f field) {
	items := f.v.Interface().([]string)

	switch f.k {
	case "BuildRequires":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s:  %s\n", f.t, strings.TrimSpace(item)))
		}
	case "Requires", "Provides":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s:       %s\n", f.t, strings.TrimSpace(item)))
		}
	case "Global", "Define":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s: %s\n", f.t, strings.TrimSpace(item)))
		}
	case "Attr":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%%attr(%s) %s\n", strings.TrimSpace(item), "NULL"))
		}
	case "Defattr":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%%defattr(%s)\n", strings.TrimSpace(item)))
		}
	case "Install", "Files", "Changelog":
		sb.WriteString(fmt.Sprintf("%s\n", f.t))
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s\n", strings.TrimSpace(item)))
		}
	}
}

func treatTemplate(rpmSpec *RpmSpec) (string, error) {
	tmplStr := createTemplateString(rpmSpec)
	tmpl, err := template.New("spec").Parse(tmplStr)
	if err != nil {
		return "", err
	}

	var data bytes.Buffer

	if err = tmpl.Execute(&data, rpmSpec); err != nil {
		return "", err
	}

	return data.String(), nil
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
