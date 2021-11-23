test:
	@echo "+ $@"
	go test -v -race ./...

fmt:
	@echo "+ $@"
	@test -z "$$(gofmt -s -l . 2>&1 | grep -v ^vendor/ | tee /dev/stderr)" || \
		(echo >&2 "+ please format Go code with 'gofmt -s'" && false)

vet:
	@echo "+ $@"
	@go vet ./...

GO_LDFLAGS ?= -w -s -X main.version=v1.3.2

build/linux:
	@echo "+ $@ $(FULL_VERSION)"
	@GOOS=linux GOARCH=amd64 go build -ldflags="$(GO_LDFLAGS)" -trimpath -o fooer ./main.go
