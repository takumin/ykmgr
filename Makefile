OS       := $(strip $(shell uname | tr A-Z a-z | sed -e 's/mingw64.*/windows/'))
ARCH     := $(strip $(shell uname -m | tr A-Z a-z))
APPNAME  := $(strip $(shell basename $(CURDIR)))
VERSION  := $(strip $(shell git describe --abbrev=0 --tags 2>/dev/null))
REVISION := $(strip $(shell git rev-parse HEAD))
SRCS     := $(shell find . -type f -name '*.go' -or -name '*.proto')

LDFLAGS_APPNAME  := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(if $(VERSION),$(VERSION),v0.0.0)"
LDFLAGS_REVISION := -X "main.Revision=$(if $(REVISION),$(REVISION),unknown)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION)'

CC ?= gcc
CXX ?= g++

ifeq ($(OS),darwin)
SHA256SUM := shasum -a 256
else
SHA256SUM := sha256sum
endif

BINNAME := $(APPNAME)-$(OS)-$(ARCH)

ifeq ($(RELEASE),true)
ARCHIVE_BASENAME := $(APPNAME)-$(VERSION)-$(OS)-$(ARCH)
else
ARCHIVE_BASENAME := $(APPNAME)-$(REVISION)-$(OS)-$(ARCH)
endif

.PHONY: all
all: bufbuild generate lint vet test build

.PHONY: tools
tools:
	go install github.com/bufbuild/buf/cmd/buf
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install honnef.co/go/tools/cmd/staticcheck

.PHONY: bufbuild
bufbuild:
	buf lint
	buf build
	buf generate

.PHONY: generate
generate:
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go generate $(LDFLAGS) ./...

.PHONY: lint
lint:
	staticcheck ./...

.PHONY: vet
vet:
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go vet -race ./...

.PHONY: test
test:
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go test -race ./...

.PHONY: build
build: bin/$(BINNAME)
bin/$(BINNAME): $(SRCS)
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go build $(LDFLAGS) -o $@

.PHONY: server
server: build
	bin/$(BINNAME) server

.PHONY: client
client: build
	bin/$(BINNAME) client

.PHONY: install
install: build
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go install $(LDFLAGS)

.PHONY: archive
archive: build
	mkdir -p dist/$(BINNAME)
	cp README.md dist/$(BINNAME)
	cp LICENSE dist/$(BINNAME)
	cp bin/$(BINNAME) dist/$(BINNAME)/$(APPNAME)
	cd dist/$(BINNAME) && 7z a ../$(ARCHIVE_BASENAME).zip *
	cd dist && $(SHA256SUM) *.zip | tee $(ARCHIVE_BASENAME).zip.sha256sum

.PHONY: release
release:
	gh release create --generate-notes

.PHONY: clean
clean:
	rm -rf bin
	rm -rf dist
