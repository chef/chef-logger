# Building the API service

## generate the API spec (should be moved to GHA)
1b. add to spec in code
> update api/specification/docs.go (header info)
2b. regenerate spec.json
> swagger generate spec -o ./api/specification/swagger.yaml --scan-models
> swagger serve -F=swagger ./api/specification/swagger.yaml

## build the docs
go get golang.org/x/tools/cmd/godoc 
~/go/bin/godoc 
http://localhost:6060

## build and run the sample
go mod vendor
go build -o bin/arch-sample-service
./bin/arch-sample-service

## test the sample
### go test

### curl or Postman

## adding go-releaser
go install github.com/goreleaser/goreleaser@latest
goreleaser --version
goreleaser init (creates the .yaml file)
goreleaser release --snapshot --clean (local build test, needs a git tag or defaults to v0.0.0, puts in /dist)

### tag and release (make sure all files are pushed up, and dist is empty, else --clean)
git tag -a v0.1.0 -m "First release"
git push origin v0.1.0
goreleaser release

