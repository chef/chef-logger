# chef-logger platform component (Go)
The debug logging component for all platform services

Copyright © 2023 Progress Software Corporation

## Introduction
This component is designed to be included in multiple services - both new platform ones like node management and courier, as well as a replacement to capabilities in pre-existing ones such as Automate -, provide a standard log output format, and operate in a number of modes based on configuration such as writing only to stdout in containers, writing to fixed size files on disk-constrained agents, or forwarding debug information to a telemetry listener/external montitoring system.

## Usage
```code```

```code```

## Building and tagging new releases of this component
go mod tidy
go build .
go test -v .
go-releaser

## adding go-releaser
go install github.com/goreleaser/goreleaser@latest
goreleaser --version
goreleaser init (creates the .yaml file)
goreleaser release --snapshot --clean (local build test, needs a git tag or defaults to v0.0.0, puts in /dist)

### tag and release (make sure all files are pushed up, and dist is empty, else --clean)
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release

