# GoVersion    ![GitHub release (latest SemVer)](https://img.shields.io/github/v/tag/neaas/go-version?display_name=tag&label=%20&sort=semver)  ![GitHub Workflow Status](https://img.shields.io/github/actions/workflow/status/neaas/go-version/checks.yml?label=%20&logo=github)
[![Go Reference](https://pkg.go.dev/badge/github.com/neaas/go-version.svg)](https://pkg.go.dev/github.com/neaas/go-version)


A (very) simple go package for handling semver compliant application versions.

---

## Features

  - Immutable application version, set via linker flags.
  - Values from `runtime/debug` available.
  - Check for updates using public GitHub release/tag information.

## Usage

Somewhere in your application, be sure that this package is imported:

```go
import (
  _ "github.com/neaas/go-version"
)
```

## Configuration

Set each configuration value using the linker flags as described below:

```bash
go build -ldflags "-X github.com/neaas/go-version.VARIABLE=VALUE" .
```

|     Variable      |                                         Description                                         |    Default Value     |
| :---------------: | :-----------------------------------------------------------------------------------------: | :------------------: |
|     `version`     |                The semver compliant version of the application being built.                 | `v0.0.0+unversioned` |
| `repositoryOwner` |                   The owner of the GitHub repository for the application.                   |       `neaas`        |
| `repositoryName`  |                      The name of the application's GitHub repository.                       |     `go-version`     |
|   `versionType`   | The method in which versions are defined for the repository <br> (`release`, `tag`, `none`) |        `tag`         |