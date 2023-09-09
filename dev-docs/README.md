# Building the chef-log module

## Building and tagging new releases of this component
```
go mod tidy
go build .
go test -v .
go-releaser
```

## adding go-releaser
go install github.com/goreleaser/goreleaser@latest
goreleaser --version
goreleaser init (creates the .yaml file)
goreleaser release --snapshot --clean (local build test, needs a git tag or defaults to v0.0.0, puts in /dist)

### tag and release (make sure all files are pushed after tagging, and dist is empty, else --clean)
(commit to main locally in GitHub Desktop - all file changes must be committed, then in Terminal do the following)
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release --clean
