GLIDE=$(shell if which glide > /dev/null 2>&1; then echo "glide"; fi)
GO_SOURCE_FILES=$(shell find . -type f -name "*.go" -not -name "bindata.go" -not -path "./vendor/*")
GO_PACKAGES=$(shell go list ./... | grep -v /vendor/)

# Install dependencies locally, optionally using go get
install:
ifdef GLIDE
	@$(GLIDE) install
else ifeq ($(FORCE), true)
	@go get
else
	@echo "Glide is necessary for installing project dependencies: http://glide.sh/ Run this command with FORCE=true to fall back to go get" 1>&2 && exit 1
endif

# Clean environment
clean:
	@rm -rf vendor/
	@go clean

# Run test suite
# Run tests
test:
ifeq ($(VERBOSE), true)
	@go test -v $(GO_PACKAGES)
else
	@go test $(GO_PACKAGES)
endif

# Check that all source files follow the Coding Style
cs:
	@gofmt -l $(GO_SOURCE_FILES) | read something && echo "Code differs from gofmt's style" 1>&2 && exit 1 || true

# Fix Coding Standard violations
csfix:
	@gofmt -l -w -s $(GO_SOURCE_FILES)

.PHONY: install clean test cs csfix
