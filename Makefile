PROJECT_NAME := "htbctl"
PKG := "gitlab.com/thesoy_sauce/$(PROJECT_NAME)
PKG_LIST := $(shell go list ./...)
GO_FILES := $(shell find . -name '*.go' | grep -v _test.go)

.PHONY: all build clean test mod lint

all: build

lint: ## Lint the go files
	@golint -set_exit_status ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

mod: ## Download dependencies
	@go mod download

build: ## Build the binaries
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -rf $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:. *?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
