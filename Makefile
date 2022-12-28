# Here you can reformat, check or build the binary.

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	go test -v ./... -bench=.

coverage:
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060