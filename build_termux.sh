#!/bin/bash

# Termux Build Script for Advanced HTTP Tool
# Run this script in Termux to build the tool

echo "🚀 Building Advanced HTTP Tool for Termux..."

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Installing..."
    pkg update && pkg install -y golang
fi

# Check Go version
go version

# Create build directory
mkdir -p ~/http-tools
cd ~/http-tools

# Copy source files
echo "📁 Copying source files..."
cp advanced_http_tool.go ~/http-tools/

# Build the tool
echo "🔨 Building the tool..."
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o http_tool advanced_http_tool.go

# Check if build was successful
if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
    echo "📱 Tool location: ~/http-tools/http_tool"
    echo ""
    echo "🚀 Usage examples:"
    echo "  ./http_tool -target=https://example.com"
    echo "  ./http_tool -target=https://example.com -bypass=stealth -concurrency=20"
    echo "  ./http_tool -target=https://example.com -config=config.json"
    echo ""
    echo "📋 Available options:"
    echo "  -target: Target URL (required)"
    echo "  -concurrency: Number of concurrent requests (default: 100, max: 50 in Termux)"
    echo "  -duration: Attack duration (default: 60s)"
    echo "  -bypass: Bypass mode (auto/stealth/aggressive/custom)"
    echo "  -rate: Requests per second (0 = unlimited)"
    echo "  -verbose: Verbose output"
    echo "  -log: Log file path"
    echo "  -config: Configuration file (JSON)"
    echo ""
    echo "⚠️  Remember: This tool is for educational purposes only!"
else
    echo "❌ Build failed!"
    exit 1
fi
