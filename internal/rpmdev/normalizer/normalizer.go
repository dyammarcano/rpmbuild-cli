package normalizer

import (
	"github.com/dyammarcano/rpmbuild-cli/internal/rpmdev"
	"log"
	"strings"
)

type Normalizer struct {
	tokens          map[string]string
	packageInstance *rpmdev.Package
}

func NewNormalizer(p *rpmdev.Package) *Normalizer {
	return &Normalizer{
		tokens:          make(map[string]string),
		packageInstance: p,
	}
}

func (n *Normalizer) replaceAndLog(key, value string) string {
	result := n.replaceTokens(value, n.tokens)
	log.Printf("%s=%s\n", key, result)
	return result
}

func (n *Normalizer) replaceTokens(in string, tokens map[string]string) string {
	for token, v := range tokens {
		in = strings.Replace(in, token, v, -1)
	}
	return in
}

func (n *Normalizer) SetTokens(arch string, version string) {
	n.tokens["!version!"] = version
	n.tokens["!arch!"] = arch
	n.tokens["!name!"] = n.packageInstance.Name
}

func (n *Normalizer) NormalizeProperties() {
	n.packageInstance.Version = n.replaceAndLog("Version", n.packageInstance.Version)
	n.packageInstance.Arch = n.replaceAndLog("Arch", n.packageInstance.Arch)
	n.packageInstance.URL = n.replaceAndLog("URL", n.packageInstance.URL)
	n.packageInstance.Summary = n.replaceAndLog("Summary", n.packageInstance.Summary)
	n.packageInstance.Description = n.replaceAndLog("Description", n.packageInstance.Description)
	n.packageInstance.ChangelogFile = n.replaceAndLog("ChangelogFile", n.packageInstance.ChangelogFile)
	n.packageInstance.ChangelogCmd = n.replaceAndLog("ChangelogCmd", n.packageInstance.ChangelogCmd)
}
