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



## The evolution to a framework
- started with http/net
- moved to gorilla/mux (now unmaintained)
- moved to echo & go-swagger (alternatively could use gin-gonic & swaggo)

## moving to Echo
- https://echo.labstack.com/guide/
go get github.com/labstack/echo/v4

## getting Go-swagger 
On Mac, use Homebrew (2 commands) https://goswagger.io/install.html 

## TODO
### Replace the Sonar Token - fails on access to SQ server

### Spec-first
1. start from an OAS spec
2. generate code, merge with echo framework 

### code-first
1b. add to spec in code
> update api/specification/docs.go (header info)
2b. regenerate spec.json
> swagger generate spec -o ./api/specification/swagger.yaml --scan-models

- version of API support
- server security tokens

- add a sample local library (mymath from command-line)
- add specific remote modules for chef-utils (from analysis)
- add default command-line arguments shown below (parameterized to echo)
- add environment variables overlay for Viper
- env vars and cmd line flags support feature flags (defined in one spot, for licensing service or standalone, telemetry collector location or standalone, logging functions, database location, etc.)
- add unit tests (go test cover or go tool cover (mpre detail))
- update CONTRIBUTING.md (Automate's is different - https://github.com/chef/automate/blob/main/CONTRIBUTING.md)
- need to add TLS support - TLS like https://go.dev/src/net/http/example_test.go
- need to add method to list all keys (possibly defaultHandler)
- need to add health checks (in next version with echo)
- DONE ---- add Dockerization steps - https://docs.github.com/en/actions/publishing-packages/publishing-docker-images, https://github.com/marketplace/actions/publish-docker

- add logging utility module
-- convert to required fields (app name, time in UTC, JSON for production, text for non-prod environments)

- add os/exec utility module
- add file reader/writer for common types (JSON, CSV, TOML) with <T> conversion
- add Postman tests

- add ORM
- add gRPC (optional support) to match OAS

- update GoLang standard
-- API requests should follow /keyvalue/add/v1/{key} -- /{namespace}/endpoint/{APIversion}/{parameters}
-- API: we could also do the version in the header, if we have an easy mechanism to select which endpoint we're going to use...

add loggingMiddleware (tracing)
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
...

func main() { // Create a new mux router
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
    ...


## Command line flags
- need feature flags as well for service (send telemetry back)

The sample CLI provides (reserved):
- standard command line flags
-- read from config file -c --config
-- --verbose (same as -d --debug)
-- -h --help (all up or on a command)
-- -v --version
-- exec (passthrough to cmdline from file or single cmd line in quotes, to powershell, cmd.exe, bash, sh, or Habitat shell)
-- --ssh "commands to pass through" to a node
-- content update
-- --u update self/tool (downloadable file will uninstall previous versions, this one will check downlaods and see if update is availanle)
-- --uninstall (prompt for confirmation, does this tool)
-- list services
-- <service> <api hook> (could include install/upgrades of server products, featureflags, reconfig, restarts, service status; security operations - cert rotation, etc. through admin API)
-- -l <uid> <pwd> | <token> -env <where server is>, signin to service environment
-- -o outputfile instead of to stdout/console
-- basic licensing acceptance (EULA for this product)

https://docs.chef.io/inspec/cli/#shell
https://docs.chef.io/server/ctl_chef_server/#install
https://docs.chef.io/workstation/knife/
https://docs.chef.io/automate/cli_chef_automate/

Should not support process control (UNIX goal of simple tools)
-- <service> <hup | int | kill | once | restart | start | stop | tail \ term >

Libraries should be included for
- environment variable support (for legacy applications)
- ability to use common identity and .pem files
- ability to call web services
- ability to call other command lines and capture output
- use common Chef utility libraries (telemetry, licensing-slim)
- output to command-line for processing by caller
- cross-platform build annd installation
- standard folder structure
- security tools scanning

PLUGINS and COBUNDLED APPS (Ruby in workstation)
Calls to service API's are wrapped in a resilience layer, which is configurable for retries, timeouts, etc.

How to use the service is described in dev-docs