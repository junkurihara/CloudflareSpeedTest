package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/XIU2/CloudflareSpeedTest/task"
	"github.com/XIU2/CloudflareSpeedTest/utils"
)

var (
	version, versionNew string
)

func init() {
	var printVersion bool
	var help = `
CloudflareSpeedTest ` + version + `
Test the latency and speed of all Cloudflare CDN IPs to find the fastest IPs (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Latency testing threads; the higher the number, the faster the latency test. Do not set too high on low-performance devices (e.g., routers); (default 200, max 1000)
    -t 4
        Number of latency tests; number of latency tests per IP; (default 4)
    -dn 10
        Number of download tests; after latency testing and sorting, download testing starts with the lowest latency IPs; (default 10)
    -dt 10
        Download test duration; maximum download test time per IP, should not be too short; (default 10 seconds)
    -tp 443
        Specify test port; used for latency/download testing; (default port 443)
    -url https://cf.xiu2.xyz/url
        Specify test URL; used for latency (HTTPing)/download testing. The default URL is not guaranteed to be available; it's recommended to set up your own;

    -httping
        Switch test mode; change latency test mode to HTTP protocol. The test URL is specified by the [-url] parameter; (default TCPing)
    -httping-code 200
        Valid status code; valid HTTP status code for HTTPing latency testing, only one allowed; (default 200, 301, 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specific regions; region names use the three-letter airport codes, separated by commas, only available in HTTPing mode; (default all regions)

    -tl 200
        Average latency upper limit; only output IPs with average latency below the specified value. Conditions can be combined; (default 9999 ms)
    -tll 40
        Average latency lower limit; only output IPs with average latency above the specified value; (default 0 ms)
    -tlr 0.2
        Packet loss rate upper limit; only output IPs with a loss rate equal to or below the specified value, range 0.00~1.00. 0 filters out any IPs with packet loss; (default 1.00)
    -sl 5
        Download speed lower limit; only output IPs with download speed above the specified value. Testing stops once the specified number [-dn] is met; (default 0.00 MB/s)

    -p 10
        Number of results to display; directly display the specified number of results after testing. Set to 0 to exit without displaying results; (default 10)
    -f ip.txt
        IP range data file; add quotes if the path contains spaces; supports other CDN IP ranges; (default ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP range data; directly specify IP range data for testing via parameters, separated by commas; (default empty)
    -o result.csv
        Write results to file; add quotes if the path contains spaces. Leave empty to avoid writing to file [-o ""]; (default result.csv)

    -dd
        Disable download testing; results will be sorted by latency (default by download speed); (default enabled)
    -allip
        Test all IPs; test every IP in the range (IPv4 only); (default randomly test one IP per /24 range)

    -v
        Print program version + check for updates
    -h
        Print help documentation
`
	var minDelay, maxDelay, downloadTime int
	var maxLossRate float64
	flag.IntVar(&task.Routines, "n", 200, "Latency test threads")
	flag.IntVar(&task.PingTimes, "t", 4, "Number of latency tests")
	flag.IntVar(&task.TestCount, "dn", 10, "Number of download tests")
	flag.IntVar(&downloadTime, "dt", 10, "Download test duration")
	flag.IntVar(&task.TCPPort, "tp", 443, "Specify test port")
	flag.StringVar(&task.URL, "url", "https://cf.xiu2.xyz/url", "Specify test URL")

	flag.BoolVar(&task.Httping, "httping", false, "Switch test mode")
	flag.IntVar(&task.HttpingStatusCode, "httping-code", 0, "Valid status code")
	flag.StringVar(&task.HttpingCFColo, "cfcolo", "", "Match specific regions")

	flag.IntVar(&maxDelay, "tl", 9999, "Average latency upper limit")
	flag.IntVar(&minDelay, "tll", 0, "Average latency lower limit")
	flag.Float64Var(&maxLossRate, "tlr", 1, "Packet loss rate upper limit")
	flag.Float64Var(&task.MinSpeed, "sl", 0, "Download speed lower limit")

	flag.IntVar(&utils.PrintNum, "p", 10, "Number of results to display")
	flag.StringVar(&task.IPFile, "f", "ip.txt", "IP range data file")
	flag.StringVar(&task.IPText, "ip", "", "Specify IP range data")
	flag.StringVar(&utils.Output, "o", "result.csv", "Output results to file")

	flag.BoolVar(&task.Disable, "dd", false, "Disable download testing")
	flag.BoolVar(&task.TestAll, "allip", false, "Test all IPs")

	flag.BoolVar(&printVersion, "v", false, "Print program version")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()

	if task.MinSpeed > 0 && time.Duration(maxDelay)*time.Millisecond == utils.InputMaxDelay {
		fmt.Println("[Tip] When using the [-sl] parameter, it's recommended to combine it with the [-tl] parameter to avoid continuous testing due to insufficient results for [-dn]...")
	}
	utils.InputMaxDelay = time.Duration(maxDelay) * time.Millisecond
	utils.InputMinDelay = time.Duration(minDelay) * time.Millisecond
	utils.InputMaxLossRate = float32(maxLossRate)
	task.Timeout = time.Duration(downloadTime) * time.Second
	task.HttpingCFColomap = task.MapColoMap()

	if printVersion {
		println(version)
		fmt.Println("Checking for updates...")
		checkUpdate()
		if versionNew != "" {
			fmt.Printf("*** New version [%s] found! Please visit [https://github.com/XIU2/CloudflareSpeedTest] to update! ***", versionNew)
		} else {
			fmt.Println("You're using the latest version [" + version + "]!")
		}
		os.Exit(0)
	}
}

func main() {
	task.InitRandSeed() // Initialize random seed

	fmt.Printf("# XIU2/CloudflareSpeedTest %s \n\n", version)

	// Start latency testing + filter by latency/packet loss
	pingData := task.NewPing().Run().FilterDelay().FilterLossRate()
	// Start download testing
	speedData := task.TestDownloadSpeed(pingData)
	utils.ExportCsv(speedData) // Output results to file
	speedData.Print()          // Print results

	if versionNew != "" {
		fmt.Printf("\n*** New version [%s] found! Please visit [https://github.com/XIU2/CloudflareSpeedTest] to update! ***\n", versionNew)
	}
	endPrint()
}

func endPrint() {
	if utils.NoPrintResult() {
		return
	}
	if runtime.GOOS == "windows" { // If running on Windows, press Enter or Ctrl+C to exit (prevents auto-close when double-clicked)
		fmt.Printf("Press Enter or Ctrl+C to exit.")
		fmt.Scanln()
	}
}

// Check for updates
func checkUpdate() {
	timeout := 10 * time.Second
	client := http.Client{Timeout: timeout}
	res, err := client.Get("https://api.xiu2.xyz/ver/cloudflarespeedtest.txt")
	if err != nil {
		return
	}
	// Read response data body: []byte
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	// Close response body stream
	defer res.Body.Close()
	if string(body) != version {
		versionNew = string(body)
	}
}
