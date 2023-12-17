.PHONY: all

# if you build this project in windows, use make tool in git bash.
app_name := wig
app_pkg := github.com/dstgo/wigfrid/cmd/daemon
user := $(shell git config user.name)
git_version := $(shell git describe --tags --always)
build_time := $(shell date +"%Y%m%d%H%M%S")
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)

resume:
	go env -w GOOS=$(host_os)
	go env -w GOARCH=$(host_arch)

build:
	go env -w GOOS="linux"
	go env -w GOARCH="amd64"
	go build -trimpath \
    			-ldflags="-X main.Author=$(user) -X main.BuildTime=$(build_time)  -X main.Version=$(git_version)" \
    			-o ./bin/$(app_name).exe $(app_pkg)
	make resume

# only for test, wigfird only run for the linux platform.
build_win:
	go env -w GOOS="windows"
	go env -w GOARCH="amd64"
	go build -trimpath \
			-ldflags="-X main.Author=$(user) -X main.BuildTime=$(build_time)  -X main.Version=$(git_version)" \
			-o ./bin/$(app_name).exe $(app_pkg)
	make resume