ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

COMMIT := $(shell git rev-parse HEAD | cut -c1-6)
TIME := $(shell date +"%Y-%m-%dT%H:%M:%S%z")
CONFIG_PKG := github.com/manboster/manboster/internal/config
VERSION ?= canary

LDFLAGS := -ldflags "-s -w -X '$(CONFIG_PKG).BuildCommit=$(COMMIT)' -X '$(CONFIG_PKG).BuildTime=$(TIME)' -X '$(CONFIG_PKG).CurrentVersion=$(VERSION)'"

.PHONY: run build build-linux-amd64 build-linux-i386 build-linux-arm64 build-mac build-mac-intel build-win-amd64 build-win-i386 build-win-arm64 build-win build-linux
run :
	@echo "=> Running: go run ./cmd/manboster/main.go $(RUN_ARGS)"
	go run -race $(LDFLAGS) ./cmd/manboster/main.go $(RUN_ARGS)

build:
	go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux-amd64 build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux-i386:
	GOOS=linux GOARCH="386" go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux-arm64:
    GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-mac build-mac-arm64:
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-mac-intel:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-win-amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go

build-win-i386:
	GOOS=windows GOARCH="386" go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go

build-win-arm64 build-win:
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go
