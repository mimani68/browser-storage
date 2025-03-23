GOOS := js
GOARCH := wasm
GO := go

WASM_FILE := main.wasm
GO_FILES := main.go

.PHONY: all build clean serve install-dependencies test

all: build serve

build:
	GOOS=${GOOS} GOARCH=${GOARCH} $(GO) build -o $(WASM_FILE) $(GO_FILES)
	mv $(WASM_FILE) ./examples/simple-caching

clean:
	rm -f $(WASM_FILE)

serve:
	python3 -m http.server 3700 --directory ./examples

install-dependencies:
	$(GO) get github.com/blevesearch/bleve/v2

compress:
	gzip -kf ./examples/simple-caching/$(WASM_FILE)

#Example of running tests if you had them.
test:
	$(GO) test ./...