# NamDoS Pro v2.0 Makefile
# Advanced DDoS Attack Tool for Termux

# Variables
BINARY_NAME=namdos_pro
BINARY_UNIX=$(BINARY_NAME)_unix
BINARY_WINDOWS=$(BINARY_NAME).exe
BINARY_ANDROID=$(BINARY_NAME)_android
GO_FILES=namdos_pro.go
VERSION=2.0.0
BUILD_TIME=$(shell date +%Y-%m-%d_%H:%M:%S)
GIT_COMMIT=$(shell git rev-parse --short HEAD)

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
LDFLAGS=-ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)"
BUILD_FLAGS=-a -installsuffix cgo

# Default target
.PHONY: all
all: clean deps build

# Build for current platform
.PHONY: build
build:
	@echo "🔨 Building NamDoS Pro for current platform..."
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_NAME) $(GO_FILES)
	@echo "✅ Build completed: $(BINARY_NAME)"

# Build for Unix/Linux
.PHONY: build-unix
build-unix:
	@echo "🔨 Building NamDoS Pro for Unix/Linux..."
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) $(BUILD_FLAGS) -o $(BINARY_UNIX) $(GO_FILES)
	@echo "✅ Build completed: $(BINARY_UNIX)"

# Build for Windows
.PHONY: build-windows
build-windows:
	@echo "🔨 Building NamDoS Pro for Windows..."
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) $(BUILD_FLAGS) -o $(BINARY_WINDOWS) $(GO_FILES)
	@echo "✅ Build completed: $(BINARY_WINDOWS)"

# Build for Android (Termux)
.PHONY: build-android
build-android:
	@echo "🔨 Building NamDoS Pro for Android (Termux)..."
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) $(BUILD_FLAGS) -o $(BINARY_ANDROID) $(GO_FILES)
	@echo "✅ Build completed: $(BINARY_ANDROID)"

# Build all platforms
.PHONY: build-all
build-all: build-unix build-windows build-android
	@echo "✅ All builds completed"

# Install dependencies
.PHONY: deps
deps:
	@echo "📦 Installing dependencies..."
	$(GOMOD) init namdos_pro
	$(GOMOD) tidy
	@echo "✅ Dependencies installed"

# Run the application
.PHONY: run
run: build
	@echo "🚀 Running NamDoS Pro..."
	./$(BINARY_NAME)

# Run with test mode
.PHONY: test
test: build
	@echo "🧪 Running NamDoS Pro in test mode..."
	./$(BINARY_NAME) -test -site https://httpbin.org/get

# Run with custom parameters
.PHONY: run-custom
run-custom: build
	@echo "🚀 Running NamDoS Pro with custom parameters..."
	./$(BINARY_NAME) -site https://example.com -threads 100 -duration 60 -type mixed_attack

# Clean build artifacts
.PHONY: clean
clean:
	@echo "🧹 Cleaning build artifacts..."
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(BINARY_UNIX) $(BINARY_WINDOWS) $(BINARY_ANDROID)
	@echo "✅ Clean completed"

# Run tests
.PHONY: test-unit
test-unit:
	@echo "🧪 Running unit tests..."
	$(GOTEST) -v ./...

# Format code
.PHONY: fmt
fmt:
	@echo "🎨 Formatting code..."
	$(GOCMD) fmt ./...
	@echo "✅ Code formatted"

# Lint code
.PHONY: lint
lint:
	@echo "🔍 Linting code..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "⚠️ golangci-lint not installed, skipping linting"; \
	fi

# Install development tools
.PHONY: install-tools
install-tools:
	@echo "🛠️ Installing development tools..."
	$(GOGET) -u github.com/golangci/golangci-lint/cmd/golangci-lint
	@echo "✅ Development tools installed"

# Create release package
.PHONY: release
release: clean build-all
	@echo "📦 Creating release package..."
	mkdir -p release
	cp $(BINARY_UNIX) release/
	cp $(BINARY_WINDOWS) release/
	cp $(BINARY_ANDROID) release/
	cp README_NAMDoS_PRO.md release/
	cp LICENSE release/
	@echo "✅ Release package created in release/ directory"

# Install to system
.PHONY: install
install: build
	@echo "📥 Installing NamDoS Pro to system..."
	@if [ -d "/data/data/com.termux" ]; then \
		cp $(BINARY_NAME) ~/../usr/bin/; \
		echo "✅ Installed to Termux system"; \
	else \
		sudo cp $(BINARY_NAME) /usr/local/bin/; \
		echo "✅ Installed to system"; \
	fi

# Uninstall from system
.PHONY: uninstall
uninstall:
	@echo "🗑️ Uninstalling NamDoS Pro from system..."
	@if [ -d "/data/data/com.termux" ]; then \
		rm -f ~/../usr/bin/$(BINARY_NAME); \
		echo "✅ Uninstalled from Termux system"; \
	else \
		sudo rm -f /usr/local/bin/$(BINARY_NAME); \
		echo "✅ Uninstalled from system"; \
	fi

# Show help
.PHONY: help
help:
	@echo "💀 NamDoS Pro v2.0 - Advanced DDoS Attack Tool"
	@echo ""
	@echo "Available targets:"
	@echo "  build          - Build for current platform"
	@echo "  build-unix     - Build for Unix/Linux"
	@echo "  build-windows  - Build for Windows"
	@echo "  build-android  - Build for Android (Termux)"
	@echo "  build-all      - Build for all platforms"
	@echo "  deps           - Install dependencies"
	@echo "  run            - Run the application"
	@echo "  test           - Run in test mode"
	@echo "  run-custom     - Run with custom parameters"
	@echo "  clean          - Clean build artifacts"
	@echo "  test-unit      - Run unit tests"
	@echo "  fmt            - Format code"
	@echo "  lint           - Lint code"
	@echo "  install-tools  - Install development tools"
	@echo "  release        - Create release package"
	@echo "  install        - Install to system"
	@echo "  uninstall      - Uninstall from system"
	@echo "  help           - Show this help"
	@echo ""
	@echo "Examples:"
	@echo "  make build-android    # Build for Termux"
	@echo "  make run              # Run interactively"
	@echo "  make test             # Test with httpbin.org"
	@echo "  make install          # Install to system"

# Show version info
.PHONY: version
version:
	@echo "💀 NamDoS Pro v$(VERSION)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Git Commit: $(GIT_COMMIT)"
	@echo "Go Version: $(shell go version)"
