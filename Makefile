# Here you can reformat, check or build the binary.

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -v --bench=. ./...

cover:
	@go test -cover ./...

cover-out:
	@go test -coverprofile cover.out ./...
	@go tool cover --html=cover.out

lint:
	@golangci-lint run ./...

doc:
	@godoc -play=true -goroot=/usr/local/go -http=:6060