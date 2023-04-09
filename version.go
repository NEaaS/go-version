package goversion

import (
	"fmt"
	"runtime/debug"
	"strings"

	"golang.org/x/mod/semver"
)

const (
	prefixPrerelease string = "-"
	prefixBuild      string = "+"
	prefixGoVersion  string = "go"

	unversioned string = "v0.0.0+unversioned"
)

var (
	version   string = unversioned
	buildInfo *debug.BuildInfo
)

// IsValid returns true if the version is in a state that is compatible with
// this package. This can be used to prevent invalid versions from being set.
func IsValid() bool {
	return semver.Canonical(version) != ""
}

// IsUnversioned returns true if the version is the default unversioned string.
// This is likely to be true in development builds.
func IsUnversioned() bool {
	return version == unversioned
}

// Version returns the canonically formatted semver compliant version string
// (with the 'v' prefix) or the default unversioned string.
func Version() string {
	return semver.Canonical(version)
}

// Prerelease returns the prerelease value of the version if any. The '-' prefix
// is stripped.
func Prerelease() string {
	return strings.TrimPrefix(semver.Prerelease(version), prefixPrerelease)
}

// Build returns the build metadata value of the version if any. The '+' prefix
// is stripped.
func Build() string {
	return strings.TrimPrefix(semver.Build(version), prefixBuild)
}

// GoVersion returns the golang version that was used to build the application.
// This does not guarantee a semver compliant response.
func GoVersion() string {
	return strings.TrimPrefix(buildInfo.GoVersion, prefixGoVersion)
}

// DepVersion returns the version of a given dependency used. If the given
// dependency was not used, an empty string is returned. This does not guarantee
// a semver compliant response.
func DepVersion(modulePath string) string {
	for _, d := range buildInfo.Deps {
		if strings.EqualFold(modulePath, d.Path) {
			return d.Version
		}
	}
	return ""
}

// VCSCommit returns the version control system commit associated with the
// application's binary.
func VCSCommit() string {
	for _, setting := range buildInfo.Settings {
		if strings.EqualFold("vcs.revision", setting.Key) {
			return setting.Value
		}
	}
	return ""
}

// GreaterThan returns true if the application version is greater than v.
func GreaterThan(v string) bool {
	return semver.Compare(version, v) > 0
}

// GreaterThanEqual returns true if the application version is greater than or
// equal to version v.
func GreaterThanEqual(v string) bool {
	return semver.Compare(version, v) >= 0
}

// LessThan returns true if the application version is less than version v.
func LessThan(v string) bool {
	return semver.Compare(version, v) < 0
}

// LessThanEqual returns true if the application version is less than or equal
// to version v.
func LessThanEqual(v string) bool {
	return semver.Compare(version, v) <= 0
}

// Equal returns true if the application version is equal to version v.
func Equal(v string) bool {
	return semver.Compare(version, v) == 0
}

// init ensures that only valid semver can be used as the application version
// and parses the build info data from the binary.
func init() {
	if !IsValid() {
		panic(fmt.Errorf("application version '%s' is not semver compliant", version))
	}
	if bi, ok := debug.ReadBuildInfo(); !ok {
		panic("could not read application build info")
	} else {
		buildInfo = bi
	}
}
