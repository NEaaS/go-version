package goversion

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v45/github"
	"golang.org/x/mod/semver"
)

const (
	versionTypeRelease string = "release"
	versionTypeTag     string = "tag"
	versionTypeNone    string = "none"
)

var (
	ghClient *github.Client = github.NewClient(nil)

	repositoryOwner string = "neaas"
	repositoryName  string = "go-version"
	versionType     string = versionTypeTag
)

// Versions fetches all of the releases/tags from the application's associated
// GitHub repository and reports back all the semver compliant tag names coupled
// with them. This only errors if the release or tag list could not be fetched.
func Versions() ([]string, error) {
	return RepoVersions(repositoryOwner, repositoryName, versionType)
}

// RepoVersions fetches all of the releases/tags from the GitHub repository with
// the owner and name provided and reports back all the semver compliant tag
// names coupled with them. This only errors if the release or tag list could
// not be fetched.
func RepoVersions(owner, name, versionType string) ([]string, error) {
	versions := make([]string, 0)
	switch versionType {
	case versionTypeRelease:
		if relV, err := releaseVersions(owner, name); err != nil {
			return versions, err
		} else {
			versions = relV
		}
	case versionTypeTag:
		if tagV, err := tagVersions(owner, name); err != nil {
			return versions, err
		} else {
			versions = tagV
		}
	case versionTypeNone:
		return versions, nil
	default:
		return versions, fmt.Errorf("version type '%s' is not supported", versionType)
	}
	semver.Sort(versions)
	return versions, nil
}

// LatestVersion determines the latest version that is valid semver and not
// prerelease from a given set of version strings. If no string provided fits
// this criteria, an empty string is returned. There is no requirement for the
// versions to be sorted prior to calling. Can be used in conjunction with
// Versions and RepoVersions.
func LatestVersion(versions []string) string {
	semver.Sort(versions)
	for i := len(versions) - 1; i >= 0; i-- {
		if !semver.IsValid(versions[i]) {
			continue // only consider valid semver versions
		}
		if strings.TrimPrefix(semver.Prerelease(versions[i]), prefixPrerelease) != "" {
			continue // do not consider pre-release versions
		}
		return versions[i]
	}
	return ""
}

// Update checks if the current application's version is lower than that of the
// version provided (only if the provided version is semver compliant). If the
// application's version is lower, Update will return true. If the application's
// version is equal to or greater than the given version, Update will return
// false. If the given version is not semver compliant, Update will return false.
func Update(latest string) bool {
	if !semver.IsValid(latest) {
		return false
	}
	return LessThan(latest)
}

func releaseVersions(owner, name string) ([]string, error) {
	releaseTags := make([]string, 0)
	releases, _, err := ghClient.Repositories.ListReleases(context.Background(), owner, name, nil)
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

func tagVersions(owner, name string) ([]string, error) {
	tags := make([]string, 0)
	ghTags, _, err := ghClient.Repositories.ListTags(context.Background(), owner, name, nil)
	if err != nil {
		return tags, fmt.Errorf("failed to get tag list from github: %w", err)
	}
	for _, t := range ghTags {
		if semver.IsValid(*t.Name) {
			tags = append(tags, *t.Name)
		}
	}
	return tags, nil
}
