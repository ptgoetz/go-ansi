.PHONY: build test test-all dist release

build: ## Build the projeect and output distribution binaries
	echo "Build"

test: ## Run tests
	echo "Test"
	go test ./...

test-all:
	echo "Test All"

dist:
	echo "Dist"

release:
	echo "Release"

