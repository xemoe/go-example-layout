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

ROOT_DIR=${PWD}
DOCKER_GO_PATH=/usr/src/myapp
DOCKER_GO_IMAGE="my-go-layout"

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
#  build, build-all

PLATFORMS=darwin linux windows
ARCHITECTURES=amd64

build: $(BINDIR)/$(BINNAME)

GO_APP_CMD	= .
$(BINDIR)/$(BINNAME): $(SRC)
	GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -o '$(BINDIR)'/$(BINNAME) $(GO_APP_CMD)

.PHONY: build-all
build-all:
	$(foreach GOOS, $(PLATFORMS),\
	$(foreach GOARCH, $(ARCHITECTURES), $(shell export GOOS=$(GOOS); export GOARCH=$(GOARCH); GO111MODULE=on go build $(GOFLAGS) -tags '$(TAGS)' -ldflags '$(LDFLAGS)' -v -o '$(BINDIR)'/$(BINNAME)-$(GOOS)-$(GOARCH) $(GO_APP_CMD) )))

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

# ------------------------------------------------------------------------------
#  rebuild

.PHONY: rebuild
rebuild: clean build

# ------------------------------------------------------------------------------
#  docker functions
#

define docker_build
	@echo
	docker build -t ${DOCKER_GO_IMAGE} \
		--build-arg USER_ID=$(shell id -u) \
		--build-arg GROUP_ID=$(shell id -g) .
endef

define docker_run
	@echo
	docker run --rm \
		-v ${ROOT_DIR}:${DOCKER_GO_PATH} \
		-w ${DOCKER_GO_PATH} \
		${DOCKER_GO_IMAGE} \
		$1
endef

# ------------------------------------------------------------------------------
#  clean exited docker process
#
DOCKER_PS_EXITED= \
	$(shell test -x docker && docker ps -a -f status=exited -f ancestor=${DOCKER_GO_IMAGE} -q)

.PHONY: clean-docker-exited
clean-docker-exited:
ifneq "$(DOCKER_PS_EXITED)" ""
	@echo "Clean exited docker build process"
	docker rm $(DOCKER_PS_EXITED)
endif

# ------------------------------------------------------------------------------
#  build with docker
#

.PHONY: docker-build-image
docker-build-image:
	$(call docker_build)
	@echo "[x]  Go build image has been created!"

.PHONY: docker-build-app
docker-build-app: docker-build-image
	$(call docker_run,make clean)
	$(call docker_run,make build-all)
