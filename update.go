package goversion

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v45/github"
	"golang.org/x/mod/semver"
)

var (
	ghClient        *github.Client = github.NewClient(nil)
	repositoryOwner string
	repositoryName  string

	candidatelabel string = "-rc"
	betaLabel      string = "-beta"
	alphaLabel     string = "-alpha"
)

// Versions fetches all of the releases from the applications associated GitHub
// repository and reports back all the tag names coupled with them, provided the
// tag name is valid semver. This only errors if the release list could not be
// fetched.
func Versions() ([]string, error) {
	releaseTags := make([]string, 0)
	releases, _, err := ghClient.Repositories.ListReleases(context.Background(), repositoryOwner, repositoryName, nil)
	if err != nil {
		return releaseTags, fmt.Errorf("failed to get release list from github: %w", err)
	}
	for _, r := range releases {
		if semver.IsValid(*r.TagName) {
			releaseTags = append(releaseTags, *r.TagName)
		}
	}
	return releaseTags, nil
}

// LatestVersion gets the latest release for the application's repository from
// GitHub and reports the tag associated with it as the version, provided the
// tag is semver compliant. This errors if the latest release can not be found,
// or the tag is not valid semver.
func LatestVersion() (string, error) {
	release, _, err := ghClient.Repositories.GetLatestRelease(context.Background(), repositoryOwner, repositoryName)
	if err != nil {
		return "", fmt.Errorf("failed to get latest release from github: %w", err)
	}
	if semver.IsValid(*release.TagName) {
		return *release.TagName, nil
	}
	return "", fmt.Errorf("release tag '%s' is not semver compliant", *release.TagName)
}

// UpdateCheck uses the GitHub release information to determine if the current
// application version is not the latest suitable version. If candidate, beta or
// alpha builds can be used as update versions, this can be specified by the
// parameters of the same names. If an update is available, true is returned
// along with the version string of the update. If the GitHub release versions
// can not be obtained, an error is returned.
func UpdateCheck(candidate, beta, alpha bool) (bool, string, error) {
	versions, err := Versions()
	if err != nil {
		return false, "", err
	}
	if len(versions) == 0 {
		return false, "", nil
	}
	for _, v := range versions {
		if !candidate && strings.Contains(semver.Prerelease(v), candidatelabel) {
			continue
		}
		if !beta && strings.Contains(semver.Prerelease(v), betaLabel) {
			continue
		}
		if !alpha && strings.Contains(semver.Prerelease(v), alphaLabel) {
			continue
		}
		if LessThan(v) {
			return true, v, nil
		}
	}
	return false, "", nil
}
