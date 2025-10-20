#!/bin/bash

# Quick Test Script for Termux
# Tests the HTTP tool with safe targets

echo "🧪 Testing Advanced HTTP Tool in Termux..."

# Check if tool exists
if [ ! -f "~/http-tools/http_tool" ]; then
    echo "❌ Tool not found. Please run build_termux.sh first"
    exit 1
fi

cd ~/http-tools

echo "🔍 Running basic connectivity test..."
./http_tool -target=http://httpbin.org/get -concurrency=5 -duration=10s -verbose

echo ""
echo "🔍 Running stealth mode test..."
./http_tool -target=http://httpbin.org/headers -concurrency=5 -duration=10s -bypass=stealth -verbose

echo ""
echo "🔍 Running configuration test..."
./http_tool -config=termux_config.json -duration=10s -verbose

echo ""
echo "✅ All tests completed!"
echo "📊 Check ~/http_tool.log for detailed logs"
