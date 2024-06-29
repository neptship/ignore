all: help

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  setup        Setup ignore cli"
	@echo "  lint      	  Scan the project using linters"
	@echo "  lint-fast    Fast scan the project using linters"
	@echo ""

setup:
	@go build && go install

lint:
	@golangci-lint run ./... --config=./.golangci.yml

lint-fast:
	@golangci-lint run ./... --fast --config=./.golangci.yml
