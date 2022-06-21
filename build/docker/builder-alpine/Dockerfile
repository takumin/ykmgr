FROM golang:alpine3.15 AS builder
RUN apk --no-cache --update add build-base git p7zip pcsc-lite-dev
WORKDIR /src/github.com/takumin/ykmgr
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY tools tools
RUN cat tools/tools.go | grep _ | awk -F'"' '{print $2}' | xargs -tI % go install %
COPY . .
RUN make archive OS=linux-musl

FROM alpine:3.16 AS distribution
RUN apk --no-cache --update add pcsc-lite-libs
COPY --from=builder /src/github.com/takumin/ykmgr/dist /opt/ykmgr-musl
