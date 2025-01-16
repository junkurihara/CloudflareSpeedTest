package task

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

const defaultInputFile = "ip.txt"

var (
	// TestAll specifies whether to test all IPs
	TestAll = false
	// IPFile is the filename containing IP ranges
	IPFile = defaultInputFile
	IPText string
)

// InitRandSeed initializes the random seed
func InitRandSeed() {
	rand.Seed(time.Now().UnixNano())
}

// isIPv4 checks whether an IP is IPv4
func isIPv4(ip string) bool {
	return strings.Contains(ip, ".")
}

// randIPEndWith generates a random byte value for the last segment of an IP
func randIPEndWith(num byte) byte {
	if num == 0 { // For /32, a single IP
		return byte(0)
	}
	return byte(rand.Intn(int(num)))
}

type IPRanges struct {
	ips     []*net.IPAddr
	mask    string
	firstIP net.IP
	ipNet   *net.IPNet
}

func newIPRanges() *IPRanges {
	return &IPRanges{
		ips: make([]*net.IPAddr, 0),
	}
}

// If it's a single IP, append a subnet mask; otherwise, get the subnet mask (r.mask)
func (r *IPRanges) fixIP(ip string) string {
	// If '/' is not present, it's a single IP, so add /32 or /128 as a subnet mask
	if i := strings.IndexByte(ip, '/'); i < 0 {
		if isIPv4(ip) {
			r.mask = "/32"
		} else {
			r.mask = "/128"
		}
		ip += r.mask
	} else {
		r.mask = ip[i:]
	}
	return ip
}

// Parse IP range to obtain IP, IP range, and subnet mask
func (r *IPRanges) parseCIDR(ip string) {
	var err error
	if r.firstIP, r.ipNet, err = net.ParseCIDR(r.fixIP(ip)); err != nil {
		log.Fatalln("ParseCIDR error", err)
	}
}

func (r *IPRanges) appendIPv4(d byte) {
	r.appendIP(net.IPv4(r.firstIP[12], r.firstIP[13], r.firstIP[14], d))
}

func (r *IPRanges) appendIP(ip net.IP) {
	r.ips = append(r.ips, &net.IPAddr{IP: ip})
}

// Returns the minimum value of the fourth IP segment and the number of usable hosts
func (r *IPRanges) getIPRange() (minIP, hosts byte) {
	minIP = r.firstIP[15] & r.ipNet.Mask[3] // Minimum value of the fourth IP segment

	// Calculate the number of hosts based on the subnet mask
	m := net.IPv4Mask(255, 255, 255, 255)
	for i, v := range r.ipNet.Mask {
		m[i] ^= v
	}
	total, _ := strconv.ParseInt(m.String(), 16, 32) // Total number of usable IPs
	if total > 255 {                                 // Adjust the number of usable IPs in the fourth segment
		hosts = 255
		return
	}
	hosts = byte(total)
	return
}

func (r *IPRanges) chooseIPv4() {
	if r.mask == "/32" { // Single IP, no need for random selection, just add it
		r.appendIP(r.firstIP)
	} else {
		minIP, hosts := r.getIPRange()    // Get the minimum value of the fourth IP segment and the number of usable hosts
		for r.ipNet.Contains(r.firstIP) { // Loop as long as the IP is within the range
			if TestAll { // Test all IPs
				for i := 0; i <= int(hosts); i++ { // Traverse from minimum to maximum value of the fourth segment
					r.appendIPv4(byte(i) + minIP)
				}
			} else { // Randomly select the last segment of the IP
				r.appendIPv4(minIP + randIPEndWith(hosts))
			}
			r.firstIP[14]++ // Increment the third segment
			if r.firstIP[14] == 0 {
				r.firstIP[13]++ // Increment the second segment
				if r.firstIP[13] == 0 {
					r.firstIP[12]++ // Increment the first segment
				}
			}
		}
	}
}

func (r *IPRanges) chooseIPv6() {
	if r.mask == "/128" { // Single IP, no need for random selection, just add it
		r.appendIP(r.firstIP)
	} else {
		var tempIP uint8                  // Temporary variable to store previous value
		for r.ipNet.Contains(r.firstIP) { // Loop as long as the IP is within the range
			r.firstIP[15] = randIPEndWith(255) // Randomly select the last segment
			r.firstIP[14] = randIPEndWith(255)

			targetIP := make([]byte, len(r.firstIP))
			copy(targetIP, r.firstIP)
			r.appendIP(targetIP) // Add IP to the address pool

			for i := 13; i >= 0; i-- { // Randomly adjust the third-to-last segment and earlier
				tempIP = r.firstIP[i]
				r.firstIP[i] += randIPEndWith(255)
				if r.firstIP[i] >= tempIP { // Exit if the value increases
					break
				}
			}
		}
	}
}

// Load IP ranges from file or inline text
func loadIPRanges() []*net.IPAddr {
	ranges := newIPRanges()
	if IPText != "" { // Get IP range data from parameters
		IPs := strings.Split(IPText, ",") // Split by commas and iterate
		for _, IP := range IPs {
			IP = strings.TrimSpace(IP) // Trim leading and trailing spaces
			if IP == "" {              // Skip empty entries
				continue
			}
			ranges.parseCIDR(IP) // Parse IP range to get IP and subnet mask
			if isIPv4(IP) {      // Generate IPv4 or IPv6 addresses for testing (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	} else { // Load IP range data from a file
		if IPFile == "" {
			IPFile = defaultInputFile
		}
		file, err := os.Open(IPFile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() { // Iterate through each line in the file
			line := strings.TrimSpace(scanner.Text()) // Trim leading and trailing spaces
			if line == "" {                           // Skip empty lines
				continue
			}
			ranges.parseCIDR(line) // Parse IP range to get IP and subnet mask
			if isIPv4(line) {      // Generate IPv4 or IPv6 addresses for testing (single/random/all)
				ranges.chooseIPv4()
			} else {
				ranges.chooseIPv6()
			}
		}
	}
	return ranges.ips
}
