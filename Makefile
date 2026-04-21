COMMIT := $(shell git rev-parse HEAD | cut -c1-6)
TIME := $(shell date +"%Y-%m-%dT%H:%M:%S%z")
CONFIG_PKG := github.com/manboster/manboster/internal/config

LDFLAGS := -ldflags "-X '$(CONFIG_PKG).BuildCommit=$(COMMIT)' -X '$(CONFIG_PKG).BuildTime=$(TIME)'"

run:
	go run $(LDFLAGS) ./cmd/manboster/main.go version

build:
	go build $(LDFLAGS) -o bin/manboster ./cmd/manboster/main.go