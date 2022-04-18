FROM golang:bullseye AS crossbuild
RUN set -eux; \
    dpkg --add-architecture i386; \
    dpkg --add-architecture amd64; \
    dpkg --add-architecture armel; \
    dpkg --add-architecture armhf; \
    dpkg --add-architecture arm64; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        crossbuild-essential-i386 \
        crossbuild-essential-amd64 \
        crossbuild-essential-armel \
        crossbuild-essential-armhf \
        crossbuild-essential-arm64 \
        libpcsclite-dev:i386 \
        libpcsclite-dev:amd64 \
        libpcsclite-dev:armel \
        libpcsclite-dev:armhf \
        libpcsclite-dev:arm64 \
        p7zip-full \
    ; \
    rm -rf /var/lib/apt/lists/*

FROM crossbuild AS builder
WORKDIR /src/github.com/takumin/ykmgr
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make tools
RUN make generate

FROM builder AS builder-i386
ENV CC=i686-linux-gnu-gcc
ENV CXX=i686-linux-gnu-g++
ENV GOARCH=386
RUN make archive ARCH=i386

FROM builder AS builder-amd64
ENV CC=x86_64-linux-gnu-gcc
ENV CXX=x86_64-linux-gnu-g++
ENV GOARCH=amd64
RUN make archive ARCH=amd64

FROM builder AS builder-armel
ENV CC=arm-linux-gnueabi-gcc
ENV CXX=arm-linux-gnueabi-g++
ENV GOARCH=arm
ENV GOARM=6
RUN make archive ARCH=armel

FROM builder AS builder-armhf
ENV CC=arm-linux-gnueabihf-gcc
ENV CXX=arm-linux-gnueabihf-g++
ENV GOARCH=arm
ENV GOARM=7
RUN make archive ARCH=armhf

FROM builder AS builder-arm64
ENV CC=aarch64-linux-gnu-gcc
ENV CXX=aarch64-linux-gnu-g++
ENV GOARCH=arm64
RUN make archive ARCH=arm64

FROM debian:bullseye-slim AS distribution
RUN set -eux; \
    dpkg --add-architecture i386; \
    dpkg --add-architecture amd64; \
    dpkg --add-architecture armel; \
    dpkg --add-architecture armhf; \
    dpkg --add-architecture arm64; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        libpcsclite1:i386 \
        libpcsclite1:amd64 \
        libpcsclite1:armel \
        libpcsclite1:armhf \
        libpcsclite1:arm64 \
    ; \
    rm -rf /var/lib/apt/lists/*
COPY --from=builder-i386 /src/github.com/takumin/ykmgr/dist /opt/ykmgr-i386
COPY --from=builder-amd64 /src/github.com/takumin/ykmgr/dist /opt/ykmgr-amd64
COPY --from=builder-armel /src/github.com/takumin/ykmgr/dist /opt/ykmgr-armel
COPY --from=builder-armhf /src/github.com/takumin/ykmgr/dist /opt/ykmgr-armhf
COPY --from=builder-arm64 /src/github.com/takumin/ykmgr/dist /opt/ykmgr-arm64
