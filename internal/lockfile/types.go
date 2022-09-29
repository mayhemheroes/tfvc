package lockfile

import (
	"errors"
	"fmt"

	"github.com/apparentlymart/go-versions/versions"
	"github.com/apparentlymart/go-versions/versions/constraints"
	tfaddr "github.com/hashicorp/terraform-registry-address"
	svchost "github.com/hashicorp/terraform-svchost"
)

type Provider = tfaddr.Provider
type Version = versions.Version
type VersionConstraints = constraints.IntersectionSpec
type Locks struct {
	Providers           map[Provider]*ProviderLock
	OverriddenProviders map[Provider]struct{}
	Sources             map[string][]byte
}

type ProviderLock struct {
	// addr is the address of the provider this lock applies to.
	Addr               Provider
	Version            Version
	VersionConstraints VersionConstraints
}

type ParserError struct {
	Summary string
	Detail  string
}

func (pe *ParserError) Error() string {
	return fmt.Sprintf("%s: %s", pe.Summary, pe.Detail)
}

const DefaultProviderRegistryHost = svchost.Hostname("registry.terraform.io")
const UnknownProviderNamespace = "?"
const LegacyProviderNamespace = "-"

func ParseVersion(str string) (Version, error) {
	return versions.ParseVersion(str)
}

func ParseVersionConstraints(str string) (VersionConstraints, error) {
	return constraints.ParseRubyStyleMulti(str)
}

var ErrParseProviderPart = errors.New("error parsing provider parts")
