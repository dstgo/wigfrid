# basic info
app := wigfrid
module := github.com/dstgo/wigfrid/cmd/$(app)
output := $(shell pwd)/bin
# meta info
build_time := $(shell date +"%Y/%m/%dT%H:%M:%SZ%z")
git_version := $(shell git describe --tags --always)
# build info
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)
os := $(host_os)
arch := $(host_arch)

ifeq ($(os), windows)
	exe := .exe
endif


.PHONY: build
build:
	# go lint
	go vet ./...

	# prepare target environment $(os)/$(arch)
	go env -w GOOS=$(os)
	go env -w GOARCH=$(arch)

	# build go module
	go build -trimpath \
		-ldflags="-X main.AppName=$(app) -X main.Version=$(git_version) -X main.BuildTime=$(build_time)" \
		-o $(output)/$(app)$(exe) \
		$(module)

	# resume host environment $(host_os)/$(host_arch)
	go env -w GOOS=$(host_os)
	go env -w GOARCH=$(host_arch)

# target protobuf files will be generated at api directory
target_proto_files := $(shell find ./proto/api/ -name *.proto)

.PHONY: proto
proto:
ifeq ($(wildcard pb), )
	@mkdir pb
endif

	@protoc --proto_path=./proto/api/ \
		   --proto_path=./proto/third_party/ \
		   --go_out=paths=source_relative:./pb \
		   --go-grpc_out=paths=source_relative:./pb \
		   --validate_out=paths=source_relative,lang=go:./pb \
		   $(target_proto_files)

wire_out := ./server/

.PHONY: wire
wire:
	# generate app dependencies injection file
	wire gen $(wire_out)