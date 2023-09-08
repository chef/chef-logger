# chef-log platform component (Go)
A debug logging component for all platform services

Copyright © 2023 Progress Software Corporation

## Introduction
This component is designed to be included in multiple services - both new platform ones like node management and courier, as well as a replacement to capabilities in pre-existing ones such as Automate -, provide a standard log output format, and operate in a number of modes based on configuration such as writing only to stdout in containers, writing to fixed size files on disk-constrained agents, or forwarding debug information to a telemetry listener/external montitoring system.

## Usage
```
defer l.Close()
l.Configure()
l.LogWarn()

// handle in-app panic by writing to log
defer func() {
if := recover(); err != nil {
     log.Println("Panic:", err)
}
}()

```

## Upgrading and releasing this module
See the README.md in the dev-dcos folder

## References
how configured in applications, services, CLI's, agents