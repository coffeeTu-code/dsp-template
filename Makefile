COMMIT_HASH=$(shell git rev-parse --verify HEAD | cut -c 1-8)
BUILD_DATE=$(shell date +%Y-%m-%d_%H:%M:%S%z)
GIT_TAG=$(shell git describe --tags)
GIT_AUTHOR=$(shell git show -s --format=%an)
SHELL:=/bin/bash

all: codegen codeline

# ----------- codeline
.PHONY: codeline
codeline:
	@echo -e "./api      \t : 功能代码 $(shell find ./api -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -adx      \t : 功能代码 $(shell find ./api/adx -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/adx -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -backend  \t : 功能代码 $(shell find ./api/backend -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/backend -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -base     \t : 功能代码 $(shell find ./api/base -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/base -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -dbstruct \t : 功能代码 $(shell find ./api/dbstruct -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/dbstruct -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -dmp      \t : 功能代码 $(shell find ./api/dmp -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/dmp -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -dsp      \t : 功能代码 $(shell find ./api/dsp -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/dsp -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -enum     \t : 功能代码 $(shell find ./api/enum -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/enum -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -juno     \t : 功能代码 $(shell find ./api/juno -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/juno -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -polaris  \t : 功能代码 $(shell find ./api/polaris -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/polaris -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -rank     \t : 功能代码 $(shell find ./api/rank -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/rank -name '*.go' | grep 'test.go' | xargs cat | wc -l)"
	@echo -e " -render   \t : 功能代码 $(shell find ./api/render -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./api/render -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

	@echo -e "./cmd      \t : 功能代码 $(shell find ./cmd -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./cmd -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

	@echo -e "./examples \t : 功能代码 $(shell find ./examples -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./examples -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

	@echo -e "./internal \t : 功能代码 $(shell find ./internal -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./internal -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

	@echo -e "./pkg      \t : 功能代码 $(shell find ./pkg -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find ./pkg -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

	@echo -e "total      \t : 功能代码 $(shell find . -name '*.go' | grep -v 'test.go' | xargs cat | wc -l) | 测试代码 $(shell find . -name '*.go' | grep 'test.go' | xargs cat | wc -l)"

# ----------- codeline end

# ----------- codegen
.PHONY: codegen
codegen:
	@echo -e "codegen    \t : $(shell cd ./api && pwd && sh codegen.sh && cd ../)"

# ----------- codegen end
