### NAMES AND LOCS ############################
APPNAME  = bot
PACKAGE  = github.com/dev-schueppchen/$(APPNAME)
LDPAKAGE = static
CONFIG   = $(CURDIR)/config/private.config.yml
###############################################

### EXECUTABLES ###############################
GO     = go
DEP    = dep
GOLINT = golint
GREP   = grep
###############################################

# ---------------------------------------------

BIN    = $(CURDIR)/bin/$(APPNAME)
TAG    = $(shell git describe --tags)
COMMIT = $(shell git rev-parse HEAD)


ifneq ($(GOOS),)
	BIN := $(BIN)_$(GOOS)
endif

ifneq ($(GOARCH),)
	BIN := $(BIN)_$(GOARCH)
endif

ifneq ($(TAG),)
	BIN := $(BIN)_$(TAG)
endif

ifeq ($(OS),Windows_NT)
	BIN := $(BIN).exe
endif


PHONY = _make
_make: deps $(BIN) cleanup

PHONY += deps
deps:
	$(DEP) ensure -v

$(BIN):
	$(GO) build  \
		-v -o $@ -ldflags "\
			-X $(PACKAGE)/$(LDPAKAGE).AppVersion=$(TAG) \
			-X $(PACKAGE)/$(LDPAKAGE).AppCommit=$(COMMIT) \
			-X $(PACKAGE)/$(LDPAKAGE).Release=TRUE" \
		$(CURDIR)/cmd/server

PHONY += test
test:
	$(GO) test -v -cover ./...

PHONY += lint
lint:
	$(GOLINT) ./... | $(GREP) -v vendor || true

PHONY += run
run:
	$(GO) run -v \
		$(CURDIR)/cmd/server

PHONY += fmt
fmt:
	$(GO) fmt ./...

PHONY += cleanup
cleanup:

PHONY += help
help:
	@echo "Available targets:"
	@echo "  #        - creates binary in ./bin"
	@echo "  cleanup  - tidy up temporary stuff created by build or scripts"
	@echo "  deps     - ensure dependencies are installed"
	@echo "  fmt      - format go code with go fmt"
	@echo "  lint     - run linters (golint)"
	@echo "  run      - debug run app (go run) with test config"
	@echo "  test     - run tests (go test)"
	@echo ""
	@echo "Cross Compiling:"
	@echo "  (env GOOS=linux GOARCH=arm make)"
	@echo ""
	@echo "Use different configs for run:"
	@echo "  make CONF=./myCustomConfig.yml run"
	@echo ""


.PHONY: $(PHONY)
