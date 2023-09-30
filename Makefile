.PHONY: build vendor flash test tiny-version lint clean help

PRINTF_FORMAT := "\033[32m%-10s\033[33m %s\033[33m %s\033[0m\n"
BUILD_FOLDER := .build

all: build

vendor: clean ## Get dependencies according to go.mod
	go mod tidy
	go mod vendor

build: vendor tiny-version
	@mkdir -p ".build"
	tinygo build -target xiao -scheduler tasks -gc conservative -size full -o ".build\build.out" "."

flash: vendor tiny-version
	tinygo flash -target xiao -scheduler tasks -gc conservative -size full "."

test: tiny-version
	@tinygo test ./internal/lamp
	@tinygo test ./internal/support

lint: vendor
	@golangci-lint --version
	@golangci-lint run -v --disable-all -E asciicheck -E dogsled -E durationcheck -E exportloopref -E forcetypeassert -E goconst -E gocritic -E nilerr -E tagliatelle -E unconvert -E unparam -E whitespace

tiny-version:
	@tinygo version

clean: ## Remove vendor and artifacts
	@rm -rf vendor
	@rm -rf $(BUILD_FOLDER)

help: ## Display available commands
	@grep -E '^[%a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'