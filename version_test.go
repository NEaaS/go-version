package goversion

import "testing"

func TestVersion(t *testing.T) {
	version = "v1.9.0-alpha+meta"
	if Version() != "v1.9.0-alpha" {
		t.Logf("version is not presented in correct canonical form ('%s')", Version())
		t.Fail()
	}
	version = "1.2.3.4-alpha"
	if Version() != "" {
		t.Logf("invlaid semver is not presented as empty")
		t.Fail()
	}
	version = "1.2-alpha"
	if Version() != "" {
		t.Logf("version should not be allowed a prerelease with no patch value")
		t.Fail()
	}
}

func TestPrerelease(t *testing.T) {
	version = "v1.9.0-alpha+meta"
	if Prerelease() != "alpha" {
		t.Logf("version prerelease is not being correctly determined")
		t.Fail()
	}
	version = "1.2.3-beta_test"
	if Prerelease() != "" {
		t.Logf("semver prerelease should not be allowed non alpha-numberic & hyphen characters")
		t.Fail()
	}
}

func TestBuild(t *testing.T) {
	version = "v1.9.0-alpha+meta"
	if Build() != "meta" {
		t.Logf("version build meta is not being correctly determined")
		t.Fail()
	}
}

func TestGoVersion(t *testing.T) {
	if GoVersion() == "" {
		t.Logf("golang version is not being reported")
		t.Fail()
	}
}

// TODO: More version tests
