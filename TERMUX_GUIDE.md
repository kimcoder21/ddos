# 📱 Termux Installation & Usage Guide

## 🚀 Quick Start

### 1. Install Termux
```bash
# Download from F-Droid (Recommended)
# Or from GitHub releases
```

### 2. Install Dependencies
```bash
# Update packages
pkg update && pkg upgrade

# Install Go
pkg install golang

# Install additional tools
pkg install git curl wget nano
```

### 3. Build the Tool
```bash
# Make build script executable
chmod +x build_termux.sh

# Run build script
./build_termux.sh
```

### 4. Run the Tool
```bash
cd ~/http-tools
./http_tool -target=https://example.com
```

## 🔧 Advanced Usage

### Configuration File
```bash
# Use configuration file
./http_tool -config=termux_config.json

# Custom configuration
./http_tool -target=https://example.com -concurrency=20 -bypass=stealth -verbose
```

### Different Bypass Modes
```bash
# Stealth mode (recommended for mobile)
./http_tool -target=https://example.com -bypass=stealth

# Aggressive mode
./http_tool -target=https://example.com -bypass=aggressive

# Auto mode (default)
./http_tool -target=https://example.com -bypass=auto
```

### Logging
```bash
# Enable logging
./http_tool -target=https://example.com -log=~/attack.log -verbose

# View logs
cat ~/attack.log
```

## 📊 Performance Optimization

### Termux-Specific Settings
- **Max Concurrency**: 50 (automatically limited)
- **Rate Limiting**: Recommended 10-20 req/s
- **Memory Usage**: Optimized for mobile devices
- **Battery**: Consider charging during long tests

### Recommended Settings
```json
{
  "concurrency": 20,
  "rate_limit": 10,
  "timeout": "30s",
  "bypass_mode": "stealth"
}
```

## 🛠️ Troubleshooting

### Common Issues

#### 1. Build Errors
```bash
# Clean and rebuild
rm -f http_tool
go clean
./build_termux.sh
```

#### 2. Permission Errors
```bash
# Make executable
chmod +x http_tool
chmod +x build_termux.sh
```

#### 3. Network Issues
```bash
# Check network
ping google.com

# Test with simple target
./http_tool -target=http://httpbin.org/get -concurrency=5
```

#### 4. Memory Issues
```bash
# Reduce concurrency
./http_tool -target=https://example.com -concurrency=10

# Check memory usage
top
```

## 📱 Mobile-Specific Features

### Battery Optimization
- Use lower concurrency (10-20)
- Enable rate limiting
- Monitor battery usage
- Consider charging during tests

### Network Considerations
- Mobile data usage
- WiFi vs Mobile data
- Network stability
- Data limits

### Display Optimization
- Compact statistics display
- Mobile-friendly output
- Log file for detailed info

## 🔒 Security & Legal

### Important Notes
- ⚠️ **Educational purposes only**
- ⚠️ **Only test systems you own**
- ⚠️ **Respect terms of service**
- ⚠️ **Follow local laws**

### Best Practices
- Use on your own networks
- Test with permission
- Monitor resource usage
- Keep logs for analysis

## 📋 Command Reference

### Basic Commands
```bash
# Basic test
./http_tool -target=https://example.com

# With custom settings
./http_tool -target=https://example.com -concurrency=20 -duration=30s

# Stealth mode
./http_tool -target=https://example.com -bypass=stealth -verbose

# With logging
./http_tool -target=https://example.com -log=~/test.log -verbose
```

### Configuration Commands
```bash
# Use config file
./http_tool -config=termux_config.json

# Override config
./http_tool -config=termux_config.json -target=https://different.com

# Verbose output
./http_tool -target=https://example.com -verbose
```

## 🎯 Examples

### Example 1: Basic Load Test
```bash
./http_tool -target=https://httpbin.org/get -concurrency=10 -duration=60s
```

### Example 2: Cloudflare Bypass Test
```bash
./http_tool -target=https://protected-site.com -bypass=stealth -js-check=true
```

### Example 3: High Performance Test
```bash
./http_tool -target=https://example.com -concurrency=30 -rate=20 -duration=300s
```

### Example 4: Configuration-Based Test
```bash
# Edit termux_config.json first
nano termux_config.json

# Then run
./http_tool -config=termux_config.json
```

## 📞 Support

If you encounter issues:
1. Check the logs: `cat ~/http_tool.log`
2. Verify network connectivity
3. Check Go installation: `go version`
4. Try with lower concurrency
5. Check available memory: `free -h`

## 🔄 Updates

To update the tool:
```bash
# Pull latest changes
git pull origin main

# Rebuild
./build_termux.sh
```

---

**Remember: This tool is for educational and testing purposes only. Use responsibly!**
