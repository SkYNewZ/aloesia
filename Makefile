# Project variables
NAME        := aloesia
VENDOR      := skynewz
DESCRIPTION := Aloesia Management API
MAINTAINER  := Quentin Lemaire <quentin@lemairepro.fr>
URL         := https://github.com/$(VENDOR)/$(NAME)
LICENSE     := GPL-3

# Build variables
BUILD_DIR   := bin
COMMIT_HASH ?= $(shell git rev-parse --short HEAD 2>/dev/null)
BUILD_DATE  ?= $(shell date +%FT%T%z)
VERSION     ?= $(shell git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.0-$(COMMIT_HASH)")

# Go variables
GOCMD       := GO111MODULE=on go
GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
GOPKGS      ?= $(shell $(GOCMD) list $(MODVENDOR) ./... | grep -v /vendor)
MODVENDOR   := -mod=vendor

GOBUILD     ?= CGO_ENABLED=0 $(GOCMD) build $(MODVENDOR)

.PHONY: all
all: clean lint test build

#########################
## Development targets ##
#########################
.PHONY: clean
clean: ## Clean workspace
	@ $(MAKE) --no-print-directory log-$@
	rm -rf ./$(BUILD_DIR)

.PHONY: lint
lint: ## Run linter
	@ $(MAKE) --no-print-directory log-$@
	GO111MODULE=on golangci-lint run ./...

.PHONY: test
test: ## Run tests
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) test $(MODVENDOR) -v $(GOPKGS) -cover

.PHONY: vendor
vendor: ## Install 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) mod vendor

.PHONY: verify
verify: ## Verify 'vendor' dependencies
	@ $(MAKE) --no-print-directory log-$@
	$(GOCMD) mod verify

###################
## Build targets ##
###################
.PHONY: build
build: clean ## Build binary for current OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	$(GOBUILD) -o ./$(BUILD_DIR)/$(GOOS)-$(GOARCH)/$(NAME)

.PHONY: build-all
build-all: GOOS      = linux darwin windows freebsd
build-all: GOARCH    = amd64
build-all: clean gox ## Build binary for all OS/ARCH
	@ $(MAKE) --no-print-directory log-$@
	@ gox -arch="$(GOARCH)" -os="$(GOOS)" -output="./$(BUILD_DIR)/{{.Dir}}-$(VERSION)-{{.OS}}-{{.Arch}}"

####################
## Helper targets ##
####################
.PHONY: gox
gox: ## Installing gox for cross compile
	@ $(MAKE) --no-print-directory log-$@
	GO111MODULE=off go get -u github.com/mitchellh/gox

.PHONY: golangci
golangci: ## Installing golangci
	@ $(MAKE) --no-print-directory log-$@
	curl -sfL https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s  -- -b $(shell go env GOPATH)/bin $(GOLANGCI_VERSION)

########################################################################
## Self-Documenting Makefile Help                                     ##
## https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html ##
########################################################################
.PHONY: help
help:
	@ grep -h -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

log-%:
	@ grep -h -E '^$*:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m==> %s\033[0m\n", $$2}'
