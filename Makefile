PROJECT_NAME := "doom-fire"
PKG := "github.com/juliendoutre/$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)

.PHONY: all dep build clean test race cover

all: build

test: dep ## Run unittests
	@go test -short ${PKG_LIST}

cover: ## Compute the project's test coverage
	@go test -cover ${PKG_LIST}

race: dep ## Run data race detector
	@go test -race ${PKG_LIST}

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/JoelOtter/termloop

build: dep ## Build the binary file
	@env GOOS=${GOOS} GOARCH=${GOARCH} go build -i -v $(PKG)

clean: ## Remove previous build
	@rm -f $(PROJECT_NAME)

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
