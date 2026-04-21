ifeq (run,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif

COMMIT := $(shell git rev-parse HEAD | cut -c1-6)
TIME := $(shell date +"%Y-%m-%dT%H:%M:%S%z")
CONFIG_PKG := github.com/manboster/manboster/internal/config
VERSION ?= canary

LDFLAGS := -ldflags "-X '$(CONFIG_PKG).BuildCommit=$(COMMIT)' -X '$(CONFIG_PKG).BuildTime=$(TIME)' -X '$(CONFIG_PKG).CurrentVersion=$(VERSION)'"

.PHONY: run build
run :
	@echo "=> Running: go run ./cmd/manboster/main.go $(RUN_ARGS)"
	go run $(LDFLAGS) ./cmd/manboster/main.go $(RUN_ARGS)

build:
	go build $(LDFLAGS) -o bin/manboster ./cmd/manboster/main.go