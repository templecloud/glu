
# BUILD := $(shell git describe --always --dirty)
BUILD = 0.0.1
VERSION ?= $(BUILD)
GOOS ?= linux
ARCH ?= amd64
DIST ?= dist/
BINARY ?= glu

.PHONY: init
init:
	@mkdir -p $(DIST)

.PHONY: build
build: init
	@GOOS=$(GOOS) GOARCH=$(ARCH) go build    	\
	    -o dist/$(BINARY)  						\
	    ./cmd/$(BINARY)

.PHONY: run
run: build
	./$(DIST)/$(BINARY)

.PHONY: test
test:
	go test ./...

.PHONY: clean
clean:
	rm -Rf $(DIST)