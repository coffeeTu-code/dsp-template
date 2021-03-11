COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash

all: codeline

# ----------- codeline
.PHONY: codeline
codeline: code-api code-api-end code-cmd code-examples code-internal code-pkg
	@echo -e "total      \t : 功能代码 $(shell find . -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find . -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-api
code-api:
	@echo -e "./api      \t : 功能代码 $(shell find ./api -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-api-end
code-api-end: code-api-adx

.PHONY: code-api-adx
code-api-adx:
	@echo -e " -adx      \t : 功能代码 $(shell find ./api/adx -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/adx -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-cmd
code-cmd:
	@echo -e "./cmd      \t : 功能代码 $(shell find ./cmd -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./cmd -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-examples
code-examples:
	@echo -e "./examples \t : 功能代码 $(shell find ./examples -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./examples -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-internal
code-internal:
	@echo -e "./internal \t : 功能代码 $(shell find ./internal -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./internal -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

.PHONY: code-pkg
code-pkg:
	@echo -e "./pkg      \t : 功能代码 $(shell find ./pkg -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./pkg -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

# ----------- codeline end