BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
COMMIT=$(shell git rev-parse --short HEAD)

# don't override user values
ifeq (,$(VERSION))
  VERSION := $(shell git describe --exact-match 2>/dev/null)
  # if VERSION is empty, then populate it with branch's name and raw commit hash
  ifeq (,$(VERSION))
    VERSION := $(BRANCH)-$(COMMIT)
  endif
endif

PACKAGES=$(shell go list ./... | grep -v '/simulation')
PACKAGES_SIMTEST=$(shell go list ./... | grep '/simulation')
LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::') # grab everything after the space in "github.com/tendermint/tendermint v0.34.7"
DOCKER := $(shell which docker)
BUILDDIR ?= $(CURDIR)/build

export GO111MODULE = on
export GOPROXY = https://goproxy.io,direct

# process build tags
build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
  ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

 # handle cleveldb
 ifeq (cleveldb,$(findstring cleveldb,$(M0_BUILD_OPTIONS)))
   BUILD_TAGS += cleveldb
 endif

# handle badgerdb
ifeq (badgerdb,$(findstring badgerdb,$(M0_BUILD_OPTIONS)))
  BUILD_TAGS += badgerdb
endif

# handle rocksdb
ifeq (rocksdb,$(findstring rocksdb,$(M0_BUILD_OPTIONS)))
  BUILD_TAGS += rocksdb
endif

# handle boltdb
ifeq (boltdb,$(findstring boltdb,$(M0_BUILD_OPTIONS)))
  BUILD_TAGS += boltdb
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags
ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=m0 \
		  -X github.com/cosmos/cosmos-sdk/version.AppName=m0d \
		  -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
		  -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
		  -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \
			-X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq (cleveldb,$(findstring cleveldb,$(BUILD_OPTIONS)))
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq (,$(findstring nostrip,$(BUILD_OPTIONS)))
  ldflags += -w -s
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'
# check for nostrip option
ifeq (,$(findstring nostrip,$(BUILD_OPTIONS)))
  BUILD_FLAGS += -trimpath
endif

# The below include contains the tools target.
#include contrib/devtools/Makefile

###############################################################################
###                              Documentation                              ###
###############################################################################

all: build wasm2c reserved

BUILD_TARGETS := build install

build: BUILD_ARGS=-o $(BUILDDIR)/

$(BUILD_TARGETS): go.sum $(BUILDDIR)/
	@go $@ -mod=readonly $(BUILD_FLAGS) $(BUILD_ARGS) ./cmd/...

$(BUILDDIR)/:
	mkdir -p $(BUILDDIR)/

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

vue:
	cd vue/wallet && npm install
	cd vue/browser && npm install
	cd vue/wallet && npm run build
	cd vue/browser && npm run build

reserved: build
	@./xdev.sh reserved

wasm2c: $(BUILDDIR)/
	@echo "Building wasm2c ..."
	@make -C x/wasm/xmodel/xvm/compile/wabt -j 4
	@cp x/wasm/xmodel/xvm/compile/wabt/build/wasm2c build

wasm2cclean:
	@make -C x/wasm/xmodel/xvm/compile/wabt clean

images/build-%: reserved
	@echo "Building Docker image tq_bc/$*"
	docker build --force-rm -f images/$*/Dockerfile \
		--build-arg GO_VER=$(GO_VER) \
		--build-arg ALPINE_VER=$(ALPINE_VER) \
		--build-arg GO_TAGS=${GO_TAGS} \
		-t ciyuntangquan/$*:${COMMIT} \
		.
	@echo "TAG=${COMMIT}" > samples/env
	docker tag ciyuntangquan/$*:${COMMIT} ciyuntangquan/$*:none
	docker login -u ciyuntangquan -p chaogaofeng
	docker push ciyuntangquan/$*:none


images/clean-%:
	-@for image in "$$(docker images --quiet --filter=reference='liubaninc/$*')"; do \
		[ -z "$$image" ] || docker rmi -f $$image; \
	done

GO_VER = 1.16
ALPINE_VER ?= 3.13
GO_TAGS ?=

m0-image: images/build-m0

m0-image-clean: images/clean-m0

clean: distclean
	rm -rf $(BUILDDIR)/

distclean: clean
	rm -rf vendor/

contractsdk:
	make -C x/wasm/xmodel/contractsdk/cpp build
	make -C x/wasm/xmodel/contractsdk/cpp test


###############################################################################
###                                Linting                                  ###
###############################################################################

lint:
	# golangci-lint run
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -d -s

format:
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs gofmt -w -s
	# find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs misspell -w
	@find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" | xargs goimports -w -local github.com/cosmos/cosmos-sdk


.PHONY: all build build-linux install format lint wasm2c reserved vue