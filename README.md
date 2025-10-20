# 💀 NamDoS Pro v2.0

[![GitHub stars](https://img.shields.io/github/stars/kimcoder21/ddos.svg)](https://github.com/kimcoder21/ddos/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/kimcoder21/ddos.svg)](https://github.com/kimcoder21/ddos/network)
[![GitHub issues](https://img.shields.io/github/issues/kimcoder21/ddos.svg)](https://github.com/kimcoder21/ddos/issues)
[![GitHub license](https://img.shields.io/github/license/kimcoder21/ddos.svg)](https://github.com/kimcoder21/ddos/blob/main/LICENSE)

**Advanced DDoS Attack Tool written in Golang for Termux Android**

![NamDoS Pro Screenshot](https://user-images.githubusercontent.com/98259155/225797849-45cb3dc6-f3fa-45eb-abfd-92db50816d8f.jpg)

## ⚠️ WARNING

**This tool is for educational and testing purposes only!**

Use only on:
- Your own servers
- Test/staging environments  
- With explicit permission

**Unauthorized use is illegal and unethical!**

## 🚀 Quick Start

### One-Command Installation
```bash
curl -sSL https://raw.githubusercontent.com/kimcoder21/ddos/main/install_namdos_pro.sh | bash
```

### Manual Installation
```bash
# Update packages
pkg update -y

# Install Golang
pkg install golang -y

# Install Git
pkg install git -y

# Download NamDoS Pro
curl -O https://raw.githubusercontent.com/kimcoder21/ddos/main/namdos_pro.go

# Make executable
chmod +x namdos_pro.go

# Run the tool
go run namdos_pro.go
```

## 📊 Advanced Features

### 🎯 Attack Powers
- 🟢 **Light Attack**: 1,000 requests, 50 threads
- 🟡 **Medium Attack**: 5,000 requests, 100 threads
- 🟠 **Heavy Attack**: 15,000 requests, 200 threads
- 🔴 **Extreme Attack**: 30,000 requests, 400 threads
- 💀 **Nuclear Attack**: 50,000 requests, 600 threads
- ☢️ **Apocalypse Attack**: 100,000 requests, 1,000 threads
- ⚙️ **Custom Attack**: Your own settings

### 🎛️ Attack Types
1. **HTTP Flood Attack** - Multiple endpoints
2. **Resource Exhaustion** - CPU, Memory, Disk I/O
3. **Bandwidth Saturation** - High bandwidth usage
4. **Mixed Attack** - All types combined
5. **Slowloris Attack** - Slow connection attack
6. **POST Flood Attack** - POST request flood
7. **GET Flood Attack** - GET request flood

### 🔧 Advanced Features
- ✅ **Real-time Statistics** - Live attack monitoring
- ✅ **Multiple Attack Vectors** - Various attack types
- ✅ **Thread Pool Optimization** - Efficient concurrent requests
- ✅ **Random User Agents** - Bypass basic protection
- ✅ **Realistic Browser Headers** - Advanced header spoofing
- ✅ **Progress Tracking** - Real-time progress display
- ✅ **ETA Calculation** - Estimated time to completion
- ✅ **Quick Test Mode** - Test target before attack
- ✅ **Custom Parameters** - Adjustable settings
- ✅ **Connection Pooling** - Optimized connections
- ✅ **SSL/TLS Support** - Secure connections
- ✅ **Memory Optimization** - Efficient memory usage
- ✅ **CPU Optimization** - Multi-core utilization
- ✅ **Signal Handling** - Graceful shutdown
- ✅ **Bandwidth Monitoring** - Data usage tracking

## 📱 Requirements

- **Android**: 7.0+ (API level 24+)
- **Termux**: Latest version
- **Golang**: 1.21+ (auto-installed)
- **Storage**: 500MB+ free space
- **RAM**: 4GB+ recommended
- **Internet**: Stable connection

## 🎯 Usage Examples

### Interactive Mode
```bash
go run namdos_pro.go
# Select attack power and follow prompts
```

### Command Line Mode
```bash
# Basic attack
go run namdos_pro.go -site https://example.com -threads 100 -duration 60

# Nuclear attack
go run namdos_pro.go -site https://example.com -threads 1000 -duration 300 -type mixed_attack

# Quick test
go run namdos_pro.go -test -site https://example.com
```

## 📁 Repository Structure

```
ddos/
├── namdos_pro.go              # Advanced NamDoS Pro tool
├── install_namdos_pro.sh      # Installation script
├── README.md                  # Main documentation
├── README_NAMDoS_PRO.md       # Advanced documentation
├── CONTRIBUTING.md            # Contributing guidelines
├── CHANGELOG.md               # Version history
├── SECURITY.md                # Security policy
├── Makefile                   # Build automation
├── go.mod                     # Go modules
├── Dockerfile                 # Container support
├── docker-compose.yml         # Multi-container setup
└── .github/                   # GitHub configuration
    ├── workflows/build.yml    # CI/CD pipeline
    ├── ISSUE_TEMPLATE/        # Issue templates
    ├── CODEOWNERS            # Code ownership
    └── dependabot.yml        # Dependency updates
```

## 🔧 Command Line Options

| Option | Description | Default |
|--------|-------------|---------|
| `-site` | Target URL to attack | Required |
| `-threads` | Number of threads | 100 |
| `-duration` | Attack duration in seconds (0 = infinite) | 0 |
| `-delay` | Delay between requests in milliseconds | 0 |
| `-timeout` | Request timeout in seconds | 30 |
| `-type` | Attack type | mixed_attack |
| `-test` | Quick test mode | false |

## 📋 Troubleshooting

### Common Issues
1. **"go: command not found"**: `pkg install golang -y`
2. **"Permission denied"**: `chmod +x namdos_pro.go`
3. **"Connection failed"**: Check internet connection
4. **"Target unreachable"**: Verify target URL

### Performance Tips
- Use fewer threads on slow devices
- Increase delay for stable connections
- Monitor system resources
- Stop attack if system becomes unresponsive

## 🔄 Updates

```bash
# Download latest version
curl -O https://raw.githubusercontent.com/kimcoder21/ddos/main/namdos_pro.go

# Make executable
chmod +x namdos_pro.go
```

## 🛡️ Legal Notice

This tool is provided for educational purposes only. The user is responsible for:
- Obtaining proper authorization before testing
- Complying with all applicable laws and regulations
- Using the tool ethically and responsibly

## 🚨 Disclaimer

The developers are not responsible for any misuse of this tool. Use at your own risk.

## 📞 Support

- 📖 Check the [documentation](https://github.com/kimcoder21/ddos/blob/main/README_NAMDoS_PRO.md)
- 🐛 Report [issues](https://github.com/kimcoder21/ddos/issues)
- 💬 Start a [discussion](https://github.com/kimcoder21/ddos/discussions)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/kimcoder21/ddos/blob/main/LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=kimcoder21/ddos&type=Date)](https://star-history.com/#kimcoder21/ddos&Date)

## 📊 Repository Stats

![GitHub repo size](https://img.shields.io/github/repo-size/kimcoder21/ddos)
![GitHub language count](https://img.shields.io/github/languages/count/kimcoder21/ddos)
![GitHub top language](https://img.shields.io/github/languages/top/kimcoder21/ddos)

---

**Remember: With great power comes great responsibility!** 💀

**⭐ If you find this project helpful, please give it a star!**
