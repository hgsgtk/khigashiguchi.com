tools:
	go get golang.org/x/tools/cmd/goimports

fmt:
	go fmt ./...

imports:
	goimports -w -l $(shell find . -name "*.go" | grep -v /vendor/)

vet:
	go vet ./...

test:
	go test -p 6 -race -cover ./...