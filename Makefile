ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

COMMIT := $(shell git rev-parse HEAD | cut -c1-6)
TIME := $(shell date +"%Y-%m-%dT%H:%M:%S%z")
CONFIG_PKG := github.com/manboster/manboster/internal/config
VERSION ?= canary
VER ?= 0.1.0

RELEASE := $(VER)-$(VERSION)-$(COMMIT)-

LDFLAGS := -ldflags "-s -w -X '$(CONFIG_PKG).BuildCommit=$(COMMIT)' -X '$(CONFIG_PKG).BuildTime=$(TIME)' -X '$(CONFIG_PKG).CurrentVersion=$(VERSION)'"

.PHONY: run build build-release build-release-all build-linux-amd64 build-linux-i386 build-linux-arm64 build-mac build-mac-intel build-mac-amd64 build-win-amd64 build-win-arm64 build-win build-linux
run :
	@echo "=> Running: go run ./cmd/manboster/main.go $(RUN_ARGS)"
	go run -race $(LDFLAGS) ./cmd/manboster/main.go $(RUN_ARGS)

build:
	go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux:
	GOOS=linux go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-mac:
	GOOS=darwin go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-mac-arm64:
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-mac-intel build-mac-amd64:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster ./cmd/manboster/main.go

build-win:
	GOOS=windows go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go

build-win-amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go

build-win-arm64:
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster.exe ./cmd/manboster/main.go

build-release:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)linux-amd64 ./cmd/manboster/main.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)linux-arm64 ./cmd/manboster/main.go
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)darwin-arm64 ./cmd/manboster/main.go
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)darwin-amd64 ./cmd/manboster/main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)win-amd64.exe ./cmd/manboster/main.go

build-release-all:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)linux-amd64 ./cmd/manboster/main.go
	GOOS=linux GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)linux-arm64 ./cmd/manboster/main.go
	GOOS=linux GOARCH=riscv64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)linux-riscv64 ./cmd/manboster/main.go
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)darwin-arm64 ./cmd/manboster/main.go
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)darwin-amd64 ./cmd/manboster/main.go
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)win-amd64.exe ./cmd/manboster/main.go
	GOOS=windows GOARCH=arm64 go build $(LDFLAGS) -trimpath -o build/manboster-$(RELEASE)win-arm64.exe ./cmd/manboster/main.go