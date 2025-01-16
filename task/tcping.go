package task

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/utils"
)

const (
	tcpConnectTimeout = time.Second * 1 // TCP connection timeout duration
	maxRoutine        = 1000            // Maximum number of concurrent routines
	defaultRoutines   = 200             // Default number of concurrent routines
	defaultPort       = 443             // Default port number for testing
	defaultPingTimes  = 4               // Default number of ping attempts
)

var (
	Routines      = defaultRoutines  // Number of routines to use
	TCPPort   int = defaultPort      // TCP port for connection
	PingTimes int = defaultPingTimes // Number of pings to send
)

type Ping struct {
	wg      *sync.WaitGroup    // WaitGroup for goroutine synchronization
	m       *sync.Mutex        // Mutex for data access synchronization
	ips     []*net.IPAddr      // List of IP addresses
	csv     utils.PingDelaySet // Data set of ping delays
	control chan bool          // Channel to control concurrency
	bar     *utils.Bar         // Progress bar
}

// Ensure default values for Routines, TCPPort, and PingTimes are valid
func checkPingDefault() {
	if Routines <= 0 {
		Routines = defaultRoutines
	}
	if TCPPort <= 0 || TCPPort >= 65535 {
		TCPPort = defaultPort
	}
	if PingTimes <= 0 {
		PingTimes = defaultPingTimes
	}
}

// Create a new Ping instance
func NewPing() *Ping {
	checkPingDefault()
	ips := loadIPRanges() // Load IP ranges
	return &Ping{
		wg:      &sync.WaitGroup{},
		m:       &sync.Mutex{},
		ips:     ips,
		csv:     make(utils.PingDelaySet, 0),
		control: make(chan bool, Routines),
		bar:     utils.NewBar(len(ips), "Available:", ""),
	}
}

// Run the ping test
func (p *Ping) Run() utils.PingDelaySet {
	if len(p.ips) == 0 {
		return p.csv
	}
	if Httping {
		fmt.Printf("Starting latency test (Mode: HTTP, Port: %d, Range: %v ~ %v ms, Packet Loss: %.2f%%)\n",
			TCPPort, utils.InputMinDelay.Milliseconds(), utils.InputMaxDelay.Milliseconds(), utils.InputMaxLossRate)
	} else {
		fmt.Printf("Starting latency test (Mode: TCP, Port: %d, Range: %v ~ %v ms, Packet Loss: %.2f%%)\n",
			TCPPort, utils.InputMinDelay.Milliseconds(), utils.InputMaxDelay.Milliseconds(), utils.InputMaxLossRate)
	}
	for _, ip := range p.ips {
		p.wg.Add(1)
		p.control <- false // Block if max concurrency is reached
		go p.start(ip)
	}
	p.wg.Wait()
	p.bar.Done() // Mark the progress bar as complete
	sort.Sort(p.csv)
	return p.csv
}

// Start a single IP address test
func (p *Ping) start(ip *net.IPAddr) {
	defer p.wg.Done()
	p.tcpingHandler(ip)
	<-p.control
}

// Test TCP connection
func (p *Ping) tcping(ip *net.IPAddr) (bool, time.Duration) {
	startTime := time.Now()
	var fullAddress string
	if isIPv4(ip.String()) {
		fullAddress = fmt.Sprintf("%s:%d", ip.String(), TCPPort)
	} else {
		fullAddress = fmt.Sprintf("[%s]:%d", ip.String(), TCPPort)
	}
	conn, err := net.DialTimeout("tcp", fullAddress, tcpConnectTimeout)
	if err != nil {
		return false, 0
	}
	defer conn.Close()
	duration := time.Since(startTime)
	return true, duration
}

// Check connection and calculate total delay
func (p *Ping) checkConnection(ip *net.IPAddr) (recv int, totalDelay time.Duration) {
	if Httping {
		recv, totalDelay = p.httping(ip)
		return
	}
	for i := 0; i < PingTimes; i++ {
		if ok, delay := p.tcping(ip); ok {
			recv++
			totalDelay += delay
		}
	}
	return
}

// Append data to the CSV record
func (p *Ping) appendIPData(data *utils.PingData) {
	p.m.Lock()
	defer p.m.Unlock()
	p.csv = append(p.csv, utils.CloudflareIPData{
		PingData: data,
	})
}

// Handle TCP ping test for an IP
func (p *Ping) tcpingHandler(ip *net.IPAddr) {
	recv, totalDlay := p.checkConnection(ip)
	nowAble := len(p.csv)
	if recv != 0 {
		nowAble++
	}
	p.bar.Grow(1, strconv.Itoa(nowAble))
	if recv == 0 {
		return
	}
	data := &utils.PingData{
		IP:       ip,
		Sended:   PingTimes,
		Received: recv,
		Delay:    totalDlay / time.Duration(recv),
	}
	p.appendIPData(data)
}
