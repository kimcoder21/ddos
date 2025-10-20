# Advanced HTTP Testing Tool with Cloudflare Bypass

## ⚠️ WARNING
This tool is for **educational and testing purposes only**! Only use on systems you own or have explicit permission to test. Unauthorized use may violate laws and terms of service.

## Features

### 🚀 Core Features
- **High Performance**: Concurrent request handling with configurable concurrency
- **Real-time Statistics**: Live monitoring of requests, responses, and performance
- **Session Management**: Cookie jar support for maintaining sessions
- **Proxy Support**: Rotating proxy support (coming soon)
- **Rate Limiting**: Configurable request rate limiting

### 🛡️ Cloudflare Bypass Features
- **Automatic Detection**: Detects Cloudflare protection automatically
- **Multiple Bypass Modes**:
  - `auto`: Intelligent header selection based on user agent
  - `stealth`: Minimal headers to avoid detection
  - `aggressive`: Many headers to confuse detection systems
  - `custom`: Use custom headers from configuration
- **JavaScript Challenge Handling**: Automatic handling of CF challenges
- **Advanced Headers**: Modern browser headers including sec-ch-ua
- **Session Persistence**: Save and load sessions for consistent bypass

## Installation

### Prerequisites
- Go 1.19 or higher
- Windows/Linux/macOS

### Build
```bash
go build -o advanced_http_tool advanced_http_tool.go
```

## Usage

### Basic Usage
```bash
./advanced_http_tool -target=https://example.com
```

### Advanced Usage
```bash
./advanced_http_tool \
  -target=https://example.com \
  -concurrency=200 \
  -duration=300s \
  -bypass=stealth \
  -rate=100 \
  -session=session.json \
  -js-check=true
```

### Command Line Options

| Option | Description | Default |
|--------|-------------|---------|
| `-target` | Target URL (required) | - |
| `-concurrency` | Number of concurrent requests | 100 |
| `-duration` | Attack duration | 60s |
| `-bypass` | Bypass mode (auto/stealth/aggressive/custom) | auto |
| `-rate` | Requests per second (0 = unlimited) | 0 |
| `-timeout` | Request timeout | 30s |
| `-session` | Session file for cookies | - |
| `-js-check` | Enable JavaScript challenge detection | true |
| `-challenge-delay` | Delay between challenge attempts | 5s |
| `-useragent` | Custom user agent file | - |
| `-proxy` | Proxy file (one per line) | - |
| `-headers` | Custom headers file | - |
| `-tls` | Use TLS/HTTPS | true |
| `-skip-verify` | Skip TLS certificate verification | false |

## Bypass Modes

### Auto Mode (Default)
Intelligently selects headers based on the user agent. Automatically detects browser type and applies appropriate headers.

### Stealth Mode
Uses minimal headers to avoid detection. Best for avoiding basic bot detection.

### Aggressive Mode
Uses many headers to confuse detection systems. May be more effective against some protection systems.

### Custom Mode
Uses headers from a custom configuration file. Allows full control over request headers.

## Configuration Files

### User Agent File
Create a text file with one user agent per line:
```
Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36
Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36
```

### Headers File
Create a JSON file with custom headers:
```json
{
  "Accept": "text/html,application/xhtml+xml",
  "Accept-Language": "en-US,en;q=0.9",
  "Custom-Header": "value"
}
```

### Proxy File
Create a text file with one proxy per line:
```
http://proxy1:8080
http://proxy2:8080
socks5://proxy3:1080
```

## Examples

### Basic Load Testing
```bash
./advanced_http_tool -target=https://example.com -concurrency=50 -duration=60s
```

### Cloudflare Bypass Testing
```bash
./advanced_http_tool -target=https://protected-site.com -bypass=stealth -js-check=true
```

### High Performance Testing
```bash
./advanced_http_tool -target=https://example.com -concurrency=500 -rate=1000 -duration=300s
```

### Session-based Testing
```bash
./advanced_http_tool -target=https://example.com -session=my_session.json -bypass=auto
```

## Statistics

The tool provides real-time statistics including:
- Total requests sent
- Successful responses
- Error responses
- Request rate (requests per second)
- Total bytes received
- Average response size
- Success rate percentage

## Legal Notice

This tool is provided for educational and testing purposes only. Users are responsible for ensuring they have proper authorization before testing any systems. The authors are not responsible for any misuse of this tool.

## Contributing

Contributions are welcome! Please ensure any new features maintain the educational purpose of this tool.

## License

This project is for educational purposes only. Use at your own risk.
