# Version

A (very) simple go package for handling semver compliant application versions.

---

## Features

  - Immutable application version, set via linker flags.
  - Values from `runtime/debug` available.
  - Check for updates using public GitHub release information.

## Usage

Somewhere in your application, be sure that this package is imported:

```go
import (
  _ "github.com/neaas/go-version"
)
```

When building your application, set the version via linker flags... For example:

```bash
go build -ldflags "-X github.com/neaas/go-version.version=v0.1.0" -o myapp .
```
