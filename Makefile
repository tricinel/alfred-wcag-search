BINARY_NAME=wcag-search
WORKFLOW_NAME=alfred-wcag-search.alfredworkflow
SOURCE_FILE=./main.go

GOLINT := $(shell which golangci-lint)

VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
LDFLAGS=-X 'main.Version=$(VERSION)' -s -w

ASSETS_DIR=./assets
DIST_DIR=./dist
PKG_DIR=./package

.PHONY: build assemble package
.PHONY: clean
.PHONY: prebuild fmt lint test

# Build universal binary for Apple Silicon and Intel
build: clean prebuild
	@echo "Building version $(VERSION)..."
	@mkdir -p $(DIST_DIR)
	GOOS=darwin GOARCH=arm64 go build -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(BINARY_NAME)-arm64 $(SOURCE_FILE)
	GOOS=darwin GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(BINARY_NAME)-amd64 $(SOURCE_FILE)
	lipo -create -output $(DIST_DIR)/$(BINARY_NAME) $(DIST_DIR)/$(BINARY_NAME)-arm64 $(DIST_DIR)/$(BINARY_NAME)-amd64
	rm $(DIST_DIR)/$(BINARY_NAME)-arm64 $(DIST_DIR)/$(BINARY_NAME)-amd64

# Prepare the distribution folder
assemble: build
	@echo "Assembling assets..."
	cp $(ASSETS_DIR)/info.plist $(DIST_DIR)/
	cp $(ASSETS_DIR)/*.png $(DIST_DIR)/
	@plutil -replace version -string "$(VERSION)" $(DIST_DIR)/info.plist || (echo "Failed to set version"; exit 1)
	@plutil -replace readme -string "$$(cat alfred-description.md)" $(DIST_DIR)/info.plist || (echo "Failed to set alfred-description. Check your file"; exit 1)
	chmod +x $(DIST_DIR)/$(BINARY_NAME)

# Create the final .alfredworkflow file
package: assemble
	@echo "Creating final package in $(PKG_DIR)..."
	@mkdir -p $(PKG_DIR)
	rm -f $(PKG_DIR)/$(WORKFLOW_NAME)
	cd $(DIST_DIR) && zip -q -r ../$(PKG_DIR)/$(WORKFLOW_NAME) . -x "*.DS_Store"
	@echo "✅ Success! File is at $(PKG_DIR)/$(WORKFLOW_NAME)"

clean:
	rm -rf $(DIST_DIR)
	rm -rf $(PKG_DIR)

prebuild: fmt lint test

fmt:
	@echo "Formatting code..."
	go fmt ./...

lint:
	@echo "Linting code..."
	go vet ./...
	@command -v golangci-lint >/dev/null 2>&1 && golangci-lint run ./... || echo "golangci-lint not found."

test:
	@echo "Running tests..."
	go test ./...

