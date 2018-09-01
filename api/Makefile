tools:
	go get golang.org/x/tools/cmd/goimports
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

fmt:
	go fmt ./...

imports:
	goimports -w -l $(shell find . -name "*.go" | grep -v /vendor/)

vet:
	go vet ./...

metalint:
	gometalinter $(shell go list ./... | grep -v /vendor/)

test:
	go test -p 6 -race -cover ./...