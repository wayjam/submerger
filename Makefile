.PHONY : build clean

all: clean build

BUILD_TIME=`date +%FT%T%z`
VERSION=`git describe --tags --exact-match 2>/dev/null`
GIT_REV=`git rev-parse --short HEAD`
GO_VERSION=$(shell go version)

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-w -s \
-X 'github.com/wayjam/submerger/internal/config.Version=${VERSION}' \
-X 'github.com/wayjam/submerger/internal/config.GitRev=${GIT_REV}' \
-X 'github.com/wayjam/submerger/internal/config.BuildTime=${BUILD_TIME}' \
-X 'github.com/wayjam/submerger/internal/config.GoVersion=${GO_VERSION}' \
"

clean:
	rm -rf build/*

build:
	@go build -v ${LDFLAGS} -o build/submerger ./cmd/submerger/main.go

