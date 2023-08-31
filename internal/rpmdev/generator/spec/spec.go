package spec

//import (
//	"fmt"
//	"github.com/pkg/errors"
//	"strings"
//)
//
//func (p *Package) preCheckString(contentBuilder strings.Builder, key string, value string) {
//	if value != "" {
//		contentBuilder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
//	}
//}
//
//func (p *Package) generateSpecContent(okVersion string, preRelease string, sourceDir string) (string, error) {
//	var contentBuilder strings.Builder
//
//	p.preCheckString(contentBuilder, "Name: %s\n", p.Name)
//	p.preCheckString(contentBuilder, "Version: %s\n", okVersion)
//	p.preCheckString(contentBuilder, "Release: %s\n", preRelease)
//	p.preCheckString(contentBuilder, "Group: %s\n", p.Group)
//	p.preCheckString(contentBuilder, "License: %s\n", p.License)
//	p.preCheckString(contentBuilder, "Url: %s\n", p.URL)
//	p.preCheckString(contentBuilder, "Summary: %s\n", p.Summary)
//	p.preCheckString(contentBuilder, "AutoReqProv: %s\n", p.AutoReqProv)
//	p.preCheckString(contentBuilder, "\nBuildRequires:%s\n", strings.Join(p.BuildRequires, ", "))
//	p.preCheckString(contentBuilder, "\nRequires:%s\n", strings.Join(p.Requires, ", "))
//	p.preCheckString(contentBuilder, "\nProvides:%s\n", strings.Join(p.Provides, ", "))
//	p.preCheckString(contentBuilder, "\nConflicts:%s\n", strings.Join(p.Conflicts, ", "))
//	p.preCheckString(contentBuilder, "\n%%description\n%s\n", p.Description)
//
//	contentBuilder.WriteString(fmt.Sprintf("\n%%prep\n"))
//	contentBuilder.WriteString(fmt.Sprintf("\n%%build\n"))
//	contentBuilder.WriteString(fmt.Sprintf("\n%%install\n"))
//
//	install, err := p.generateInstallSection(sourceDir)
//	if err != nil {
//		return "", errors.WithStack(err)
//	}
//
//	return contentBuilder.String(), nil
//}
