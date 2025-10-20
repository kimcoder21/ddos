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
- **Golang**: 1.19+ (auto-installed)
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

# Custom attack
go run namdos_pro.go -site https://example.com -threads 500 -duration 0 -delay 10 -timeout 60 -type http_flood
```

### Using Launcher Scripts
```bash
# Interactive launcher
~/start_namdos_pro.sh

# Quick test
~/test_namdos_pro.sh https://example.com

# Command line launcher
~/namdos-pro/run_namdos_pro.sh -site https://example.com
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

## 📊 Performance Comparison

| Feature | Original NamDoS | NamDoS Pro v2.0 |
|---------|-----------------|-----------------|
| **Language** | Golang | Golang |
| **Max Threads** | 4096 | 1000+ |
| **Attack Types** | 1 (HTTP Flood) | 7 (Multiple) |
| **Real-time Stats** | ❌ | ✅ |
| **Progress Tracking** | ❌ | ✅ |
| **ETA Calculation** | ❌ | ✅ |
| **Memory Optimization** | ❌ | ✅ |
| **CPU Optimization** | ❌ | ✅ |
| **Signal Handling** | ❌ | ✅ |
| **Bandwidth Monitoring** | ❌ | ✅ |
| **Custom Headers** | ❌ | ✅ |
| **SSL/TLS Support** | ❌ | ✅ |
| **Connection Pooling** | ❌ | ✅ |
| **Quick Test Mode** | ❌ | ✅ |
| **Interactive Mode** | ❌ | ✅ |
| **Command Line Mode** | ✅ | ✅ |

## 📁 Repository Structure

```
ddos/
├── namdos_pro.go              # Advanced NamDoS Pro tool
├── install_namdos_pro.sh      # Installation script
├── README_NAMDoS_PRO.md       # This documentation
├── README.md                  # Main documentation
└── LICENSE                    # License file
```

## 🔧 Configuration

### Attack Types Explained

#### 1. HTTP Flood Attack
- Targets multiple endpoints
- Uses various HTTP methods
- Random user agents
- Realistic headers

#### 2. Resource Exhaustion
- Focuses on resource-heavy endpoints
- Targets uploads, themes, plugins
- Consumes CPU and memory
- Disk I/O intensive

#### 3. Bandwidth Saturation
- High bandwidth consumption
- Large file requests
- Media content targeting
- Network intensive

#### 4. Mixed Attack
- Combines all attack types
- Most effective approach
- Balanced resource usage
- Comprehensive coverage

#### 5. Slowloris Attack
- Slow connection establishment
- Keeps connections open
- Exhausts connection limits
- Server resource intensive

#### 6. POST Flood Attack
- POST request flooding
- Form submission simulation
- Database intensive
- Processing heavy

#### 7. GET Flood Attack
- GET request flooding
- Page request simulation
- Cache intensive
- Memory heavy

## 📋 Troubleshooting

### Common Issues

1. **"go: command not found"**
   ```bash
   pkg install golang -y
   ```

2. **"Permission denied"**
   ```bash
   chmod +x namdos_pro.go
   ```

3. **"Connection failed"**
   - Check internet connection
   - Verify target URL
   - Try different network

4. **"Out of memory"**
   - Reduce thread count
   - Close other apps
   - Restart Termux

5. **"Target unreachable"**
   - Verify URL format
   - Check if site is online
   - Try different target

### Performance Tips

#### For Low-End Devices:
- Use 50-100 threads
- Increase delay (10-50ms)
- Shorter duration
- Monitor memory usage

#### For High-End Devices:
- Use 500-1000 threads
- Minimal delay (0-10ms)
- Longer duration
- Multiple attack types

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
- Complying with local laws and regulations
- Using the tool ethically and responsibly

## 🚨 Disclaimer

The developers are not responsible for any misuse of this tool. Use at your own risk.

## 📞 Support

- 📖 Check the [documentation](https://github.com/naem021/ddos/blob/main/README_NAMDoS_PRO.md)
- 🐛 Report [issues](https://github.com/naem021/ddos/issues)
- 💬 Start a [discussion](https://github.com/naem021/ddos/discussions)

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/naem021/ddos/blob/main/LICENSE) file for details.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=naem021/ddos&type=Date)](https://star-history.com/#naem021/ddos&Date)

## 📊 Repository Stats

![GitHub repo size](https://img.shields.io/github/repo-size/naem021/ddos)
![GitHub language count](https://img.shields.io/github/languages/count/naem021/ddos)
![GitHub top language](https://img.shields.io/github/languages/top/naem021/ddos)

---

**Remember: With great power comes great responsibility!** 💀

**⭐ If you find this project helpful, please give it a star!**
