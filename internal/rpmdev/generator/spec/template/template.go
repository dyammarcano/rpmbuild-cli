package template

import (
	"bytes"
	"fmt"
	"github.com/dyammarcano/rpmbuild-cli/internal/rpmdev/generator/spec"
	"github.com/dyammarcano/utils/orderedmap"
	"reflect"
	"strings"
	"text/template"
)

func createTemplateString(rpmSpec *spec.RpmSpec) string {
	templateStrings := orderedmap.NewOrderedMap()

	templateStrings.Set("Global", "{{if .Global}}%global {{.}}{{end}}\n")
	templateStrings.Set("Define", "{{if .Define}}%define {{.}}{{end}}\n")
	templateStrings.Set("Name", "{{if .Name}}Name:           {{.Name}}{{end}}\n")
	templateStrings.Set("Version", "{{if .Version}}Version:        {{.Version}}{{end}}\n")
	templateStrings.Set("Release", "{{if .Release}}Release:        {{.Release}}{{else}}Release:        1%{?dist}{{end}}\n")
	templateStrings.Set("Summary", "{{if .Summary}}Summary:        {{.Summary}}{{end}}\n")
	templateStrings.Set("Group", "{{if .Group}}Group:          {{.Group}}{{end}}\n")
	templateStrings.Set("Distribution", "{{if .Distribution}}Distribution:   {{.Distribution}}{{end}}\n")
	templateStrings.Set("License", "{{if .License}}License:        {{.License}}{{end}}\n")
	templateStrings.Set("BuildArch", "{{if .BuildArch}}BuildArch:      {{.BuildArch}}{{end}}\n")
	templateStrings.Set("BuildRequires", "{{if .BuildRequires}}BuildRequires:  {{.BuildRequires}}{{end}}\n")
	templateStrings.Set("Epoch", "{{if .Epoch}}Epoch:          {{.Epoch}}{{end}}\n")
	templateStrings.Set("BuildRoot", "{{if .BuildRoot}}BuildRoot:      {{.BuildRoot}}{{end}}\n")
	templateStrings.Set("Prefixes", "{{if .Prefixes}}Prefixes:       {{.Prefixes}}{{end}}\n")
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

	valueOf := reflect.ValueOf(rpmSpec).Elem()
	tmplStr := strings.Builder{}

	//TODO remove [] from the templateStrings items

	for pair := templateStrings.Oldest(); pair != nil; pair = pair.Next() {
		fieldVal := valueOf.FieldByName(pair.Key.(string))
		if !fieldVal.IsValid() || fieldVal.String() == "" {
			continue
		}

		if fieldVal.Kind() == reflect.Slice {
			fixItemList(&tmplStr, pair.Key.(string), fieldVal.Interface().([]string))
			continue
		}

		tmplStr.WriteString(pair.Value.(string))
	}
	return tmplStr.String()
}

func fixItemList(sb *strings.Builder, key string, items []string) {
	switch key {
	case "BuildRequires":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s:  %s\n", key, strings.TrimSpace(item)))
		}
	case "Requires", "Provides":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s:       %s\n", key, strings.TrimSpace(item)))
		}
	case "Global":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%%global: %s\n", strings.TrimSpace(item)))
		}
	case "Define":
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%%define: %s\n", strings.TrimSpace(item)))
		}
	}
}

func treatTemplate(rpmSpec *spec.RpmSpec) (string, error) {
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
