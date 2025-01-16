# XIU2/CloudflareSpeedTest

[![Go Version](https://img.shields.io/github/go-mod/go-version/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Go&color=00ADD8&logo=go)](https://github.com/XIU2/CloudflareSpeedTest/)
[![Release Version](https://img.shields.io/github/v/release/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Release&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/releases/latest)
[![GitHub license](https://img.shields.io/github/license/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=License&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Star](https://img.shields.io/github/stars/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Star&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)
[![GitHub Fork](https://img.shields.io/github/forks/XIU2/CloudflareSpeedTest.svg?style=flat-square&label=Fork&color=00ADD8&logo=github)](https://github.com/XIU2/CloudflareSpeedTest/)

Many foreign websites use Cloudflare CDN, but the IPs assigned to visitors from mainland China are not very friendly (high latency, high packet loss, slow speed).
Although Cloudflare has made all [IP ranges](https://www.cloudflare.com/zh-cn/ips/) public, finding the right one among so many IPs could be exhausting, which is why this software was created.

**"Choose Optimal IP" to test Cloudflare CDN latency and speed, and get the fastest IP (IPv4+IPv6)!** If you find it useful, **please give it a `⭐` to encourage!**

> _Share my other open-source projects: [**TrackersList.com** - A list of popular BT trackers on the web! Effectively improves BT download speed~](https://github.com/XIU2/TrackersListCollection) <img src="https://img.shields.io/github/stars/XIU2/TrackersListCollection.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_
> _[**UserScript** - 🐵 Github high-speed download, Zhihu enhancements, automatic seamless paging, eye protection mode, and more than ten **Tampermonkey scripts**~](https://github.com/XIU2/UserScript) <img src="https://img.shields.io/github/stars/XIU2/UserScript.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_
> _[**SNIProxy** - 🧷 Simple SNI Proxy for personal use (supports all platforms, all systems, front proxy, simple configuration, etc.~](https://github.com/XIU2/SNIProxy) <img src="https://img.shields.io/github/stars/XIU2/SNIProxy.svg?style=flat-square&label=Star&color=4285dd&logo=github" height="16px" />_

> This project also supports latency testing for **other CDN / website IPs** (e.g., [CloudFront](https://github.com/XIU2/CloudflareSpeedTest/discussions/304), [Gcore](https://github.com/XIU2/CloudflareSpeedTest/discussions/303) CDN), but for download speed testing, you need to find the addresses yourself.

> [!IMPORTANT]
> Cloudflare CDN has **banned proxy usage**. For **proxy-based CDN** use, please bear the risks yourself, and avoid excessive reliance on [#382](https://github.com/XIU2/CloudflareSpeedTest/discussions/382) [#383](https://github.com/XIU2/CloudflareSpeedTest/discussions/383).

****

## \# Quick Usage

### Download and Run

1. Download the compiled executable file ([Github Releases](https://github.com/XIU2/CloudflareSpeedTest/releases) / [LanZou Cloud](https://pan.lanpw.com/b0742hkxe)) and extract it.
2. Double-click to run the `CloudflareST.exe` file (for Windows), and wait for the speed test to complete...

<details>
<summary><code><strong>「 Click to view the usage example for Linux systems 」</strong></code></summary>

****

The following commands are only examples. Please visit [**Releases**](https://github.com/XIU2/CloudflareSpeedTest/releases) for the version numbers and file names.

``` yaml
# If this is your first time using it, it is recommended to create a new folder (skip this step during future updates)
mkdir CloudflareST

# Go into the folder (for future updates, just repeat the download and extraction commands from here)
cd CloudflareST

# Download the CloudflareST archive (replace the [version] and [filename] in the URL according to your needs)
wget -N https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# If you are downloading from within China, use one of the following mirrors for faster speeds:
# wget -N https://ghp.ci/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://ghproxy.cc/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://ghproxy.net/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# wget -N https://gh-proxy.com/https://github.com/XIU2/CloudflareSpeedTest/releases/download/v2.2.5/CloudflareST_linux_amd64.tar.gz
# If the download fails, try removing the -N option (if updating, remember to delete the old archive with rm CloudflareST_linux_amd64.tar.gz first)

# Extract the archive (no need to delete old files, they will be directly overwritten, replace the filename as needed)
tar -zxf CloudflareST_linux_amd64.tar.gz

# Grant execute permissions
chmod +x CloudflareST

# Run (without parameters)
./CloudflareST

# Run (with parameters example)
./CloudflareST -dd -tll 90
```

> If the **average latency is very low** (e.g., 0.xx), it means the CloudflareST **used a proxy during the test**. Please turn off the proxy software before testing again.
> If running on a **router**, it is recommended to disable the proxy in the router (or exclude it), as the test results may be **inaccurate/cannot be used**.

</details>

****

> _A simple tutorial on running CloudflareST independently on **mobile**: **[Android](https://github.com/XIU2/CloudflareSpeedTest/discussions/61)**, **[Android App](https://github.com/xianshenglu/cloudflare-ip-tester-app)**, **[iOS](https://github.com/XIU2/CloudflareSpeedTest/discussions/321)**_

> [!NOTE]
> Note! This software is only applicable for websites and **does not support selecting optimal IPs for Cloudflare WARP using UDP protocol**. See: [#392](https://github.com/XIU2/CloudflareSpeedTest/discussions/392)

### Result Example

After the test is complete, the **fastest 10 IPs** will be displayed by default, for example:

``` bash
IP Address         Sent  Received  Loss Rate  Average Latency  Download Speed (MB/s)
104.27.200.69     4     4         0.00       146.23           28.64
172.67.60.78      4     4         0.00       139.82           15.02
104.25.140.153    4     4         0.00       146.49           14.90
104.27.192.65     4     4         0.00       140.28           14.07
172.67.62.214     4     4         0.00       139.29           12.71
104.27.207.5      4     4         0.00       145.92           11.95
172.67.54.193     4     4         0.00       146.71           11.55
104.22.66.8       4     4         0.00       147.42           11.11
104.27.197.63     4     4         0.00       131.29           10.26
172.67.58.91      4     4         0.00       140.19           9.14
...

# If the average latency is very low (like 0.xx), it means CloudflareST was using a proxy during testing. Please turn off the proxy software before testing again.
# If running on a router, please disable the proxy in the router (or exclude it), as the test results may be inaccurate/cannot be used.

# Since each test randomly picks an IP from each range, the results may differ every time, which is normal!

# Note! I found that the first test after the computer starts tends to have a noticeably higher latency (even for manual TCPing), but subsequent tests work fine.
# Therefore, it is recommended to do a few test runs after booting up, before doing the first formal speed test (no need to wait for the latency test to finish, just wait for the progress bar to move and then you can stop).

# The entire process with the default settings typically follows these steps:
# 1. Latency test (default TCPing mode, HTTPing mode requires manual parameters)
# 2. Latency sorting (sorting latency from low to high, filtering by conditions, so there may be some low-latency but packet-losing IPs at the back)
# 3. Download test (testing starts from the lowest-latency IP, stopping after 10 IPs are tested)
# 4. Speed sorting (sorting speeds from high to low)
# 5. Output results (controlled by parameters to either output to the command line (-p 0) or to a file (-o ""))

# Note: The output result file result.csv may show garbled text in Microsoft Excel, but other spreadsheet software or text editors display it fine.
```

The first line of the test result shows the **fastest IP with the highest download speed and lowest average latency**!

The complete results are saved in the `result.csv` file in the current directory. Open it using **Notepad/spreadsheet software**, and the format will be as follows:

```csv
IP Address, Sent, Received, Loss Rate, Average Latency, Download Speed (MB/s)
104.27.200.69, 4, 4, 0.00, 146.23, 28.64
```

> _You can filter and process the complete results **further** according to your needs or explore advanced usage with **specific filtering conditions**!_

****

## Advanced Usage

By default, the software runs with standard parameters. If you want more comprehensive or customized results, you can specify custom parameters.

```Dart
C:\>CloudflareST.exe -h

CloudflareSpeedTest vX.X.X
Test the latency and speed of all Cloudflare CDN IPs and get the fastest IPs (IPv4+IPv6)!
https://github.com/XIU2/CloudflareSpeedTest

Parameters:
    -n 200
        Latency test threads; the more threads, the faster the latency test, but do not set too high for devices with lower performance (e.g., routers); (default is 200, max 1000)
    -t 4
        Latency test attempts per IP; (default is 4 attempts)
    -dn 10
        Number of download tests; after sorting by latency, this many IPs will be used for download speed tests (default is 10 IPs)
    -dt 10
        Duration of each download test per IP; cannot be too short; (default is 10 seconds)
    -tp 443
        Specify test port for latency and download tests (default is port 443)
    -url https://cf.xiu2.xyz/url
        Specify the test URL; used for HTTPing latency and download tests; the default URL may not be reliable, it's recommended to self-host it.

    -httping
        Switch to HTTPing test mode for latency testing; the address used will be from the [-url] parameter; (default is TCPing)
        Note: HTTPing is a type of network scan, so if running on a server, reduce concurrency (-n) to avoid service suspension by some providers.
    -httping-code 200
        Valid HTTP status code for HTTPing latency tests; (default is 200, 301, 302)
    -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
        Match specific regions by three-letter airport codes, separated by commas (lowercase supported), only available in HTTPing mode; (default is all regions)

    -tl 200
        Latency upper limit; only output IPs with latency below the specified value; (default is 9999 ms)
    -tll 40
        Latency lower limit; only output IPs with latency above the specified value; (default is 0 ms)
    -tlr 0.2
        Packet loss upper limit; only output IPs with packet loss below or equal to the specified rate, range from 0.00 to 1.00; (default is 1.00)
    -sl 5
        Minimum download speed; only output IPs with download speed above the specified value, download tests will stop once the specified number is reached [-dn]; (default is 0.00 MB/s)

    -p 10
        Display result count; directly display the specified number of results after testing; if 0, it exits without showing results; (default is 10 IPs)
    -f ip.txt
        IP segment data file; if the path contains spaces, use quotes; supports other CDN IP segments; (default is ip.txt)
    -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
        Specify IP segments to test; separated by commas; (default is empty)
    -o result.csv
        Write results to a file; if the path contains spaces, use quotes; if empty, no file is written [-o ""]; (default is result.csv)

    -dd
        Disable download speed tests; results will be sorted by latency (default is by download speed)
    -allip
        Test all IPs; perform tests for every IP in the range (only supports IPv4)

    -v
        Print the program version and check for updates
    -h
        Print help documentation
```

### Interface Explanation

To avoid misunderstandings regarding **output content during the testing process** (available IPs, queues, interrupted or stuck downloads, etc.), here’s an explanation.

<details>
<summary><code><strong>「 Click to expand and view details 」</strong></code></summary>

****

> The following example includes common parameters: `-tll 40 -tl 150 -sl 1 -dn 5`, and the final output will look like this:

```python
# XIU2/CloudflareSpeedTest vX.X.X

Starting latency test (mode: TCP, port: 443, range: 40 ~ 150 ms, packet loss: 1.00)
321 / 321 [-----------------------------------------------------------] Available: 30
Starting download test (minimum speed: 1.00 MB/s, quantity: 5, queue: 10)
3 / 5 [-----------------------------------------↗--------------------]
IP Address           Sent  Received  Loss Rate  Average Latency  Download Speed (MB/s)
XXX.XXX.XXX.XXX   4       4      0.00    83.32    3.66
XXX.XXX.XXX.XXX   4       4      0.00    107.81   2.49
XXX.XXX.XXX.XXX   4       3      0.25    149.59   1.04

Complete test results have been written to the result.csv file, you can view it using Notepad or spreadsheet software.
Press Enter or Ctrl+C to exit.
```

****

> New users of CloudflareST might be confused when seeing **30 available IPs** after the latency test, but why does the final result only show 3?
> What does "queue" in download tests mean? Do I have to wait in line for download tests?

CloudflareST first performs latency testing. During this phase, the available IP count (shown on the right) represents the number of IPs that **did not time out**. After the latency tests, the software filters out IPs based on the specified **latency and packet loss** conditions, leaving only those that meet the criteria (in this case, 10). These IPs are then queued for download testing (shown as `queue: 10`).

As the download test is a single-threaded process, the IPs are tested one at a time, so the number of IPs waiting in the queue is displayed as `queue`.

****

> You might have noticed that, even though you specified that you want to find 5 IPs that meet the download speed requirement, why did the test "stop" after finding just 3?

In the download test progress bar, `3 / 5` means that **3** IPs met the download speed condition (i.e., the download speed is higher than `1 MB/s`), and the **5** refers to your request to find **5** IPs that meet the download speed condition (`-dn 5`).

> Additionally, if you specify `-dn` as greater than the number of IPs remaining in the download test queue, for example, after latency testing, only **4** IPs remain, then the progress bar will show **4** IPs in the queue, not 5 as you originally requested.

After completing the download test on these 10 IPs, the software only found 3 IPs with download speeds above `1 MB/s`. The remaining 7 IPs did not meet the requirement.

Therefore, this does not mean **"the test always stops before finding 5 IPs"**. Rather, the software tested all the IPs, but only found 3 that met the required conditions.

****

Here’s the English translation of the new situation you’ve described:

---

There’s another scenario when there are a lot of available IPs (hundreds or thousands), and you’ve set a download speed condition. In this case, you might encounter the situation where: **the download test progress bar seems stuck at `X / 5`**.

This isn’t actually stuck. The progress bar only increases when an IP that meets the condition is found. If no such IPs are found, the software will continue testing. This results in the appearance of the progress bar being stuck, but it's essentially the software telling you that the download speed condition you set is too high, and you need to adjust your expectations.

---

If you want to avoid such situations where you don’t find many IPs that meet your conditions, you should **lower the download speed upper limit parameter `-sl`** or remove it altogether.

As long as you specify the `-sl` parameter, the software will keep testing until it finds enough IPs to meet the `-dn` quantity (default is 10). If you remove `-sl` and set `-dn 20` instead, it will test the 20 lowest-latency IPs and stop once completed, saving you time.

---

Additionally, if all the IPs in the queue have been tested but none meet the download speed requirement, the software will **output the download test results for all IPs in the queue**. This lets you see the download speeds for those IPs, so you can assess the situation and **adjust the `-sl` parameter accordingly**.

Similarly, for latency testing, the values `可用: 30` (available IPs) and `队列：10` (IPs in queue) will help you assess whether the latency conditions you set are too strict. If there are many available IPs but only 2 or 3 remain after filtering, it’s a clear sign that you should **lower your expectations for latency or packet loss conditions**.

These two mechanisms help inform you whether your **latency/packet loss conditions** and **download speed conditions** are appropriate for your needs.

</details>

****

### Usage Example

To specify parameters on Windows, run the command in CMD or add the parameters to the shortcut target.

> [!TIP]
> - All parameters have **default values**, and those using the default values can be omitted (you can **select based on your needs**). The order of parameters **doesn’t matter**.
> - On Windows **PowerShell**, just replace `CloudflareST.exe` in the commands with `.\CloudflareST.exe`.
> - On Linux, replace `CloudflareST.exe` with `./CloudflareST`.

---

#### \# Running CloudflareST with Parameters in CMD

For users unfamiliar with command-line programs, here’s a brief explanation.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

---

Many people encounter errors when running CloudflareST with the **absolute path** in CMD because the default `-f ip.txt` parameter is a relative path. You need to specify the absolute path of `ip.txt`, but this can be cumbersome. It’s better to navigate to the CloudflareST program directory and run it using **relative paths**:

**Method 1**:
1. Open the folder where CloudflareST is located.
2. Right-click in an empty space and select **[Open command window here]** to open CMD, with the current directory set to CloudflareST.
3. Enter a command with parameters like: `CloudflareST.exe -tll 50 -tl 200` to run.

**Method 2**:
1. Open the folder where CloudflareST is located.
2. In the address bar, type `cmd` and press Enter to open CMD with the current directory set to CloudflareST.
3. Enter a command with parameters like: `CloudflareST.exe -tll 50 -tl 200` to run.

> Of course, you can also open any CMD window and enter something like `cd /d "D:\Program Files\CloudflareST"` to go to the program directory.

> **Tip**: If you are using **PowerShell**, just replace `CloudflareST.exe` in the command with `.\CloudflareST.exe`.

</details>

---

#### \# Running CloudflareST with Parameters via Windows Shortcut

If you don’t frequently change the run parameters (e.g., typically just double-click to run), using a shortcut is more convenient.

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

---

Right-click on `CloudflareST.exe` and select **[Create shortcut]**, then right-click the shortcut and select **[Properties]** to modify the **Target**:

``` bash
# If you don’t want to output a result file, add -o " " (the empty space in quotes ensures the parameter is not omitted).
D:\ABC\CloudflareST\CloudflareST.exe -n 500 -t 4 -dn 20 -dt 5 -o " "

# If the file path contains quotes, make sure the launch parameters are outside the quotes, and there should be a space between the quotes and the `-` parameter.
"D:\Program Files\CloudflareST\CloudflareST.exe" -n 500 -t 4 -dn 20 -dt 5 -o " "

# Note: The "Start in" field of the shortcut cannot be empty, otherwise, the ip.txt file won’t be found due to the absolute path.
```

</details>

---

#### \# IPv4/IPv6

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

---

``` bash
# Use the built-in IPv4 data file to test these IPv4 addresses (by default, the -f parameter is ip.txt, so you can omit this).
CloudflareST.exe -f ip.txt

# Use the built-in IPv6 data file to test these IPv6 addresses.
# Since version v2.1.0, IPv4+IPv6 mixed testing is supported, and the `-ipv6` parameter has been removed, so a file can contain both IPv4 and IPv6 addresses.
CloudflareST.exe -f ipv6.txt

# Alternatively, specify IP addresses directly.
CloudflareST.exe -ip 1.1.1.1,2606:4700::/32
```

> When testing IPv6, you might notice the number of IPs tested varies each time. This is because there are billions of IPv6 addresses, and most of them are unused, so only a portion of available IPv6 ranges are written in the `ipv6.txt` file. You can scan and modify this file as needed, with ASN data sourced from: [bgp.he.net](https://bgp.he.net/AS13335#_prefixes6)

</details>

---

#### \# HTTPing

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

---

Currently, there are two latency testing modes: **TCP protocol** and **HTTP protocol**.
The TCP protocol is faster and consumes fewer resources, with a timeout of 1 second (default mode).
The HTTP protocol is suitable for quickly testing if a domain can be accessed by an IP, with a timeout of 2 seconds.
Generally, the latency for an IP is: **ICMP < TCP < HTTP**, with the HTTP protocol being more sensitive to packet loss and network fluctuations.

> Note: HTTPing is essentially a **network scanning** action. If you run it on a server, you may need to **lower the concurrency** (`-n`), otherwise, some services might suspend your access. If you encounter a situation where the number of usable IPs decreases after an initial HTTPing test but then recovers after a while, it’s likely due to **temporary restrictions** triggered by network scanning. In such cases, lowering the concurrency (`-n`) can help reduce the occurrence of this issue.

> Additionally, the software’s HTTPing only fetches the **response headers**, not the body content (file size doesn’t affect the HTTPing test, but a large file is necessary for the download speed test).

``` bash
# Add the -httping parameter to switch to HTTP latency testing mode.
CloudflareST.exe -httping

# The software will evaluate the validity of the response based on the HTTP status code (by default, status codes 200, 301, and 302 are considered valid, but you can specify a different code).
CloudflareST.exe -httping -httping-code 200

# Use the -url parameter to specify the HTTPing test address (this can be any webpage URL, not limited to file addresses).
CloudflareST.exe -httping -url https://cf.xiu2.xyz/url

# Note: If the test address uses HTTP, make sure to add -tp 80 (this affects the latency and download test ports).
# Similarly, for testing port 80, specify an HTTP address using the -url parameter. For non-standard ports, ensure the download test address supports access via that port.
CloudflareST.exe -httping -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

---

#### \# Match a Specific Region (colo Airport Three-letter Code)

<details>
<summary><code><strong>「 Click to Expand and View Content 」</strong></code></summary>

---

Cloudflare CDN’s node IPs are Anycast IPs, meaning the corresponding server node and region for each IP are not fixed but dynamic. The server node assigned to the same IP can vary depending on the region, ISP, and time. For example, a US user may be assigned to a US server, while a Japanese user might be routed to a server in Japan. In mainland China, however, you may be routed to servers in other countries due to network restrictions.

> **Note**: Although Cloudflare CDN has many nodes in Asia, this doesn’t mean you’ll always be able to access them. For instance, in Singapore, you may find many nearby nodes, but when testing the same IP, you might not encounter any Cloudflare node nearby. This is due to Anycast IP routing being dynamic.

You can find the actual node regions assigned to an IP by manually accessing `http://CloudflareIP/cdn-cgi/trace`.

> This feature supports both Cloudflare CDN and AWS CloudFront CDN, with common airport codes for both.

``` bash
# Specify region names, and the latency test results will show only IPs from the selected regions (no further download tests will be performed unless specified by -dd).
# The region names are airport three-letter codes, and you can specify multiple regions separated by commas (lowercase is supported after v2.2.3).

CloudflareST.exe -httping -cfcolo HKG,KHH,NRT,LAX,SEA,SJC,FRA,MAD
```

> Common airport codes for both CDNs can be found here: [Cloudflare status](https://www.cloudflarestatus.com/)

</details>

---

This detailed guide explains how to use CloudflareST on various systems and provides examples of different parameters and configurations for testing IP addresses and servers effectively.

****

#### # File Relative/Absolute Path

<details>
<summary><code><strong>「 Click to expand and view content 」</strong></code></summary>

****

``` bash
# Specify IPv4 data file, no result shown, directly exit, output result to file (-p value is 0)
CloudflareST.exe -f 1.txt -p 0 -dd

# Specify IPv4 data file, no output result to file, directly show result (-p value is 10 entries, -o value is empty but must include quotes)
CloudflareST.exe -f 2.txt -o "" -p 10 -dd

# Specify IPv4 data file and output result to file (relative path, i.e., in the current directory, add quotes if it contains spaces)
CloudflareST.exe -f 3.txt -o result.txt -dd


# Specify IPv4 data file and output result to file (relative path, i.e., in the abc folder in the current directory, add quotes if it contains spaces)
# Linux (in the abc folder under the CloudflareST program directory)
./CloudflareST -f abc/3.txt -o abc/result.txt -dd

# Windows (note the backslash)
CloudflareST.exe -f abc\3.txt -o abc\result.txt -dd


# Specify IPv4 data file and output result to file (absolute path, i.e., in the C:\abc\ directory, add quotes if it contains spaces)
# Linux (in the /abc/ directory)
./CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows (note the backslash)
CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd


# If you want to run CloudflareST with the absolute path, then the file names in the -f / -o parameters must also be absolute paths, otherwise it will return an error that the file is not found!
# Linux (in the /abc/ directory)
/abc/CloudflareST -f /abc/4.txt -o /abc/result.csv -dd

# Windows (note the backslash)
C:\abc\CloudflareST.exe -f C:\abc\4.txt -o C:\abc\result.csv -dd
```
</details>

****

#### # Speed Test for Other Ports

<details>
<summary><code><strong>「 Click to expand and view content 」</strong></code></summary>

****

``` bash
# If you want to test a non-default port 443, you need to specify it with the -tp parameter (this parameter will affect the latency test/download test used port)

# If you want to perform a latency test on port 80 and download test (if -dd disables download test, then not needed), you also need to specify an HTTP:// protocol download address (and this address will not be forcibly redirected to HTTPS, as that would use port 443)
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm

# If it is a non-80 or 443 port, you need to confirm that the download test address supports accessing that non-standard port.
```

</details>

****

#### # Custom Speed Test Address

<details>
<summary><code><strong>「 Click to expand and view content 」</strong></code></summary>

****

``` bash
# This parameter is suitable for download testing and latency testing using HTTP protocol. For the latter, the address can be any webpage URL (not limited to specific file addresses)

# Address requirements: can be directly downloaded, file size over 200MB, uses Cloudflare CDN
CloudflareST.exe -url https://cf.xiu2.xyz/url

# Note: If the test address is HTTP protocol (the address cannot be forcibly redirected to HTTPS), remember to add -tp 80 (this parameter will affect latency testing/download testing port). If it is a non-80 or 443 port, then confirm whether the download test address supports access through that port.
CloudflareST.exe -tp 80 -url http://cdn.cloudflare.steamstatic.com/steam/apps/5952/movie_max.webm
```

</details>

****

#### # Custom Test Conditions (Specify Latency/Packet Loss/Download Speed Target Range)

<details>
<summary><code><strong>「 Click to expand and view content 」</strong></code></summary>

****

> Note: The **available number** next to the latency test progress bar only refers to the **number of IPs that did not timeout** during the latency test and is unrelated to the latency upper and lower limits.

- Only specify **[average latency upper limit]** condition

``` bash
# Average latency upper limit: 200 ms, download speed lower limit: 0 MB/s
# That is, find IPs with an average latency lower than 200 ms, then perform 10 download tests sorted by latency from low to high
CloudflareST.exe -tl 200
```

> If **no IP meets the latency** condition, no content will be output.

****

- Only specify **[average latency upper limit]** condition and **only latency test, no download test**

``` bash
# Average latency upper limit: 200 ms, download speed lower limit: 0 MB/s, quantity: unknown number
# That is, only output IPs lower than 200ms, and no download test will be performed (because download test is not performed, the -dn parameter is invalid)
CloudflareST.exe -tl 200 -dd
```

- Only specify **[packet loss probability upper limit]** condition

``` bash
# Packet loss probability upper limit: 0.25
# That is, find IPs with packet loss rate less than or equal to 0.25, range from 0.00 to 1.00. If -tlr 0, it means filtering out any IP with packet loss
CloudflareST.exe -tlr 0.25
```

****

- Only specify **[download speed lower limit]** condition

``` bash
# Average latency upper limit: 9999 ms, download speed lower limit: 5 MB/s, quantity: 10 (optional)
# That is, find 10 IPs with an average latency lower than 9999 ms and a download speed higher than 5 MB/s, and then stop testing
CloudflareST.exe -sl 5 -dn 10
```

> If **no IP meets the speed** condition, it will **ignore the condition and output all IP test results** (for you to adjust the condition next time).
> If no average latency upper limit is specified, if the required number of IPs is not found, it will **continue testing** until enough are found.
> Therefore, it is recommended to **specify both [download speed lower limit] + [average latency upper limit]** at the same time to stop testing once the desired latency is reached even if the quantity is insufficient.

****

- Specify both **[average latency upper limit] + [download speed lower limit]** conditions at the same time

``` bash
# Both average latency upper limit and download speed lower limit support decimals (e.g., -sl 0.5)
# Average latency upper limit: 200 ms, download speed lower limit: 5.6 MB/s, quantity: 10 (optional)
# That is, find 10 IPs with an average latency lower than 200 ms and download speed higher than 5.6 MB/s to stop testing
CloudflareST.exe -tl 200 -sl 5.6 -dn 10
```

> If **no IP meets the latency** condition, no content will be output.
> If **no IP meets the speed** condition, it will ignore the condition and output all IP test results (for you to adjust the condition next time).
> Therefore, it is recommended to first test without conditions to see the range of average latency and download speed to avoid setting conditions that are **too low/high**!

> Since the publicly available Cloudflare IP segments are **origin IP + Anycast IP**, and **origin IPs** cannot be used, download testing will be 0.00.
> You can add `-sl 0.01` (download speed lower limit) during runtime to filter out **origin IPs** (results with download speed less than 0.01MB/s).

</details>

****

#### # Speed Test for a Single or Multiple IPs

<details>
<summary><code><strong>「 Click to expand and view content 」</strong></code></summary>

****

**Method 1**:
Directly specify the IP segment data to test.

``` bash
# Enter the CloudflareST directory, then run:
# Windows system (run in CMD)
CloudflareST.exe -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32

# Linux system
./CloudflareST -ip 1.1.1.1,2.2.2.2/24,2606:4700::/32
```

****

**Method 2**:
Or write these IPs in any text file, such as `1.txt`

```
1.1.1.1
1.1.1.200
1.0.0.1/24
2606:4700::/32
```

> For a single IP, you can omit the `/32` subnet mask (i.e., `1.1.1.1` is equivalent to `1.1.1.1/32`).
> A subnet mask `/24` refers to the last segment of the IP, i.e., `1.0.0.1~1.0.0.255`.

Then run CloudflareST with the `-f 1.txt` parameter to specify the IP segment data file.

``` bash
# Enter the CloudflareST directory, then run:
# Windows system (run in CMD)
CloudflareST.exe -f 1.txt

# Linux system
./CloudflareST -f 1.txt

# For IP segments like 1.0.0.1/24, only the last segment will be randomly selected (1.0.0.1~255).
# To test all IPs in the segment, add the -allip parameter.
```

</details>

****

#### # One-Click Acceleration for All Cloudflare CDN Websites (No Need to Add Domains to Hosts One by One)

As I mentioned before, the purpose of developing this software project is to accelerate access to websites using Cloudflare CDN by **modifying the Hosts** file.

However, as stated in [**#8**](https://github.com/XIU2/CloudflareSpeedTest/issues/8), adding domains one by one to Hosts is **too cumbersome**. So I found a **one-click solution**. You can check out [**How to Accelerate All Cloudflare CDN Websites Locally in One Click!**](https://github.com/XIU2/CloudflareSpeedTest/discussions/71) and another tutorial on [modifying domain resolution IP with local DNS](https://github.com/XIU2/CloudflareSpeedTest/discussions/317).

****

#### # Automatic Hosts Update

Considering that many people need to replace IPs in the Hosts file after obtaining the fastest Cloudflare CDN IPs, you can check out this [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/discussions/312) for **Windows/Linux automatic Hosts update scripts**!

****

## Feedback

If you encounter any issues, please first check [**Issues**](https://github.com/XIU2/CloudflareSpeedTest/issues), [Discussions](https://github.com/XIU2/CloudflareSpeedTest/discussions) to see if someone else has already asked about it (remember to check [**Closed**](https://github.com/XIU2/CloudflareSpeedTest/issues?q=is%3Aissue+is%3Aclosed)).

If you cannot find a similar issue, please open a new [**Issue**](https://github.com/XIU2/CloudflareSpeedTest/issues/new) to inform me!

> [!NOTE]
> **Note!** _For issues that are not related to the CloudflareST feedback or function suggestions, please head to the forum (top `💬 Discussions`)._

****

## License

The GPL-3.0 License.
