package goversion

import "testing"

func TestVersions(t *testing.T) {
	repositoryOwner = "google"
	repositoryName = "go-github"
	versions, err := Versions()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	if len(versions) == 0 {
		t.Logf("failed to get any valid versions")
		t.FailNow()
	}
}

// TODO: More update tests
