GOCMD=go
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=hexcat

all: install

build:
	$(GOBUILD) -v

install:
	$(GOINSTALL) -v .

.PHONY: test
test:
	$(GOTEST) -v ./pkg/...

.PHONY: windows
windows:
	mkdir -p release
	GOOS=windows GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-windows-amd64

.PHONY: linux
linux:
	mkdir -p release
	GOOS=linux GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-linux-amd64

.PHONY: darwin
darwin:
	mkdir -p release
	GOOS=darwin GOARCH=amd64 go build -o release/$(BINARY)-v1.0.0-darwin-amd64

.PHONY: release
release: windows linux darwin


.PHONY: lint
lint:
	gometalinter -j 1 --vendor --exclude=libexec ./...

clean:
	# $(GOCLEAN)
	rm -f $(GOPATH)/bin/$(BINARY_NAME)
