#!/bin/bash

echo "╔══════════════════════════════════════════════════════════════╗"
echo "║                    💀 NamDoS Pro v2.0 💀                    ║"
echo "║                Advanced DDoS Attack Tool                    ║"
echo "║                    Installation Script                      ║"
echo "╚══════════════════════════════════════════════════════════════╝"
echo ""

# Check if running on Termux
if [ ! -d "/data/data/com.termux" ]; then
    echo "⚠️ Warning: This tool is optimized for Termux Android"
    echo "It may not work properly on other systems"
    read -p "Press Enter to continue..."
fi

# Update package lists
echo "📦 Updating package lists..."
pkg update -y

# Install Golang
echo "🐹 Installing Golang..."
pkg install golang -y

# Install Git
echo "📥 Installing Git..."
pkg install git -y

# Install additional tools
echo "🔧 Installing additional tools..."
pkg install curl -y
pkg install wget -y

# Set up Go environment
echo "🌍 Setting up Go environment..."
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc

# Create project directory
echo "📁 Creating project directory..."
mkdir -p ~/namdos-pro
cd ~/namdos-pro

# Download NamDoS Pro
echo "📥 Downloading NamDoS Pro..."
curl -O https://raw.githubusercontent.com/kimcoder21/ddos/main/namdos_pro.go

# Make executable
echo "🔧 Making executable..."
chmod +x namdos_pro.go

# Create run script
echo "📝 Creating run script..."
cat > run_namdos_pro.sh << 'EOF'
#!/bin/bash
cd ~/namdos-pro
go run namdos_pro.go "$@"
EOF

chmod +x run_namdos_pro.sh

# Create desktop script
echo "📝 Creating desktop script..."
cat > ~/start_namdos_pro.sh << 'EOF'
#!/bin/bash
echo "╔══════════════════════════════════════════════════════════════╗"
echo "║                    💀 NamDoS Pro v2.0 💀                    ║"
echo "║                Advanced DDoS Attack Tool                    ║"
echo "╚══════════════════════════════════════════════════════════════╝"
echo ""
echo "🚀 Starting NamDoS Pro..."
echo ""
cd ~/namdos-pro
go run namdos_pro.go
EOF

chmod +x ~/start_namdos_pro.sh

# Create quick test script
echo "📝 Creating quick test script..."
cat > ~/test_namdos_pro.sh << 'EOF'
#!/bin/bash
if [ -z "$1" ]; then
    echo "Usage: $0 <target_url>"
    echo "Example: $0 https://example.com"
    exit 1
fi
cd ~/namdos-pro
go run namdos_pro.go -test -site "$1"
EOF

chmod +x ~/test_namdos_pro.sh

echo ""
echo "✅ Installation completed!"
echo ""
echo "🚀 Usage:"
echo "   Interactive mode: ~/start_namdos_pro.sh"
echo "   Command line:    ~/namdos-pro/run_namdos_pro.sh -site https://example.com"
echo "   Quick test:      ~/test_namdos_pro.sh https://example.com"
echo ""
echo "📁 Files installed:"
echo "   ~/namdos-pro/namdos_pro.go     - Main program"
echo "   ~/start_namdos_pro.sh          - Interactive launcher"
echo "   ~/test_namdos_pro.sh           - Quick test launcher"
echo "   ~/namdos-pro/run_namdos_pro.sh - Command line launcher"
echo ""
echo "⚠️ WARNING: Use only on test servers or with permission!"
echo ""

# Ask if user wants to run the tool
read -p "🎯 Do you want to run NamDoS Pro now? (y/n): " choice
if [ "$choice" = "y" ] || [ "$choice" = "Y" ]; then
    echo "🚀 Starting NamDoS Pro..."
    ~/start_namdos_pro.sh
fi
