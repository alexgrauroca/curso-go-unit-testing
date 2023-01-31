# go-testing-course
Project to learn unit testing in Golang

## Docker
- Build: `docker build . -t go-unit-testing`
- Run: `docker run -p 8082:8080 go-unit-testing`

## Testing
- Specific folder: `go test ./{folder} -v`
- All folders and files: `go test ./... -v`
- Coverage: `go test ./{folder} -coverprofile=coverage.out`
- Render coverage result: `go tool cover -html=coverage.out -o coverage.html`
- VSCode: test file & run package tests.

## Benchmarking
- Run: `go test ./{folder} -run={file} -bench= >bench.log`