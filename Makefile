PACKAGE = github.com/vechain/thor-go-sdk
MAJOR = $(shell go version | cut -d' ' -f3 | cut -b 3- | cut -d. -f1)
MINOR = $(shell go version | cut -d' ' -f3 | cut -b 3- | cut -d. -f2)
PACKAGES = `go list ./... | grep -v '/vendor/' | grep -v '/cmd/test/'`

help:
	@egrep -h '\s#@\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?#@ "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

go_version_check:
	@if test $(MAJOR) -lt 1; then \
		echo "Go 1.24 or higher required"; \
		exit 1; \
	else \
		if test $(MAJOR) -eq 1 -a $(MINOR) -lt 19; then \
			echo "Go 1.24 or higher required"; \
			exit 1; \
		fi \
	fi

test:| go_version_check #@ Run the tests
	@docker pull vechain/thor:latest
	@docker pull ghcr.io/vechain/thor:release-galactica-latest
	@go test -cover $(PACKAGES)

test-coverage:| go_version_check #@ Run the tests with coverage
	@go test -coverpkg=./... -race -coverprofile=coverage.out -covermode=atomic $(PACKAGES)
	@go tool cover -html=coverage.out

lint_command_check:
	@command -v golangci-lint || (echo "golangci-lint not found, please install it from https://golangci-lint.run/usage/install/" && exit 1)

lint: | go_version_check lint_command_check #@ Run 'golangci-lint'
	@echo "running golanci-lint..."
	@golangci-lint run --config .golangci.yml
	@echo "running modernize..."
	@go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@v0.18.1 ./...
	@echo "done."

lint-fix: | go_version_check lint_command_check #@ Attempt to fix linting issues
	@echo "running golanci-lint..."
	@golangci-lint run --config .golangci.yml --fix
	@echo "running modernize..."
	@go run golang.org/x/tools/gopls/internal/analysis/modernize/cmd/modernize@v0.18.1 --fix ./...
	@echo "done."
