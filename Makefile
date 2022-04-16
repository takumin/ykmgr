OS ?= $(shell uname | tr A-Z a-z)
ARCH := $(shell uname -m | tr A-Z a-z)

APPNAME := $(shell basename $(CURDIR))

ifeq (,$(shell git describe --abbrev=0 --tags 2>/dev/null))
VERSION := v0.0.0
else
VERSION := $(shell git describe --abbrev=0 --tags)
endif

ifeq (,$(shell git rev-parse --short HEAD 2>/dev/null))
REVISION := unknown
else
REVISION := $(shell git rev-parse --short HEAD)
endif

CC ?= gcc
CXX ?= g++

LDFLAGS_APPNAME  := -X "main.AppName=$(APPNAME)"
LDFLAGS_VERSION  := -X "main.Version=$(VERSION)"
LDFLAGS_REVISION := -X "main.Revision=$(REVISION)"
LDFLAGS          := -ldflags '-s -w $(LDFLAGS_APPNAME) $(LDFLAGS_VERSION) $(LDFLAGS_REVISION)'

SRCS := $(shell find . -type f -name '*.go' -or -name '*.proto')

.PHONY: all
all: lint vet build

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

.PHONY: lint
lint:
	buf lint
	staticcheck ./...

.PHONY: vet
vet:
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go vet -race ./...

.PHONY: test
test:
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go test -race ./...

.PHONY: build
build: bin/$(APPNAME)
bin/$(APPNAME): $(SRCS)
	buf build
	buf generate
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go generate $(LDFLAGS) ./...
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go build $(LDFLAGS) -o $@

.PHONY: server
server: build
	bin/$(APPNAME) server

.PHONY: client
client: build
	bin/$(APPNAME) client

.PHONY: install
install: build
	CC=$(CC) CXX=$(CXX) CGO_ENABLED=1 go install $(LDFLAGS)

.PHONY: archive
archive: build
	mkdir -p dist/archive
	cp README.md dist/archive
	cp LICENSE dist/archive
	cp bin/$(APPNAME) dist/archive
ifeq ($(RELEASE),true)
	cd dist/archive && 7z a ../$(APPNAME)-$(VERSION)-$(OS)-$(ARCH).zip *
	cd dist && sha256sum *.zip | tee $(APPNAME)-$(VERSION)-$(OS)-$(ARCH).zip.sha256sum
else
	cd dist/archive && 7z a ../$(APPNAME)-$(VERSION)-$(REVISION)-$(OS)-$(ARCH).zip *
	cd dist && sha256sum *.zip | tee $(APPNAME)-$(VERSION)-$(REVISION)-$(OS)-$(ARCH).zip.sha256sum
endif

.PHONY: release
release:
	gh release create --generate-notes

.PHONY: clean
clean:
	rm -rf bin
	rm -rf dist
