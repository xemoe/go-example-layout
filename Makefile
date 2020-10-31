BINDIR      	:= $(CURDIR)/bin
BINNAME     	?= go-example
INSTALL_PATH 	?= /usr/local/bin

GOBIN     		= $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN      		= $(shell go env GOPATH)/bin
endif
GOX       		= $(GOBIN)/gox
GOIMPORTS  		= $(GOBIN)/goimports
ARCH       		= $(shell uname -p)

#
# go option
#
PKG        := ./...
TESTS      := .
TESTFLAGS  :=
GOFLAGS    :=

# Required for globs to work correctly
SHELL      		= /usr/bin/env bash

# ------------------------------------------------------------------------------
#  default target

all: build

# ------------------------------------------------------------------------------
#  build

build: $(BINDIR)/$(BINNAME)

GO_APP_CMD	= ./cmd/go-example
$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) $(GO_APP_CMD)

# ------------------------------------------------------------------------------
#  install

.PHONY: install
install: build
	@install "$(BINDIR)/$(BINNAME)" "$(INSTALL_PATH)/$(BINNAME)"

# ------------------------------------------------------------------------------
#  test

.PHONY: test
test: build
test: TESTFLAGS += -race -v
test: test-unit

.PHONY: test-unit
test-unit:
	@echo
	@echo "==> Running unit tests <=="
	GO111MODULE=on go test $(GOFLAGS) -run $(TESTS) $(PKG) $(TESTFLAGS)

# ------------------------------------------------------------------------------
#  clean

.PHONY: clean
clean:
	@rm -rf '$(BINDIR)'
