Analytix is a swiss army knife TUI built in Go with Bubbletea that brings together everything you'd want to know about your machine and network into a single, fast TUI dashboard — CPU, RAM, disk, temperatures, running processes, live logs, open connections, LAN device scanning, port scanning, MAC spoofing, and security tooling, all without leaving your terminal. 

# Tickets

## High Priority

- [ ] **System Monitor**
  - [x] CPU usage % with live sparkline graph
  - [x] RAM usage %
  - [x] Tick-based periodic refresh loop
  - [x] Responsive graph width adapts to terminal size
  - [ ] Per-core CPU breakdown instead of aggregate
  - [ ] Disk I/O monitor
    - [ ] Read speed per disk
    - [ ] Write speed per disk
    - [ ] Sparkline graph per disk
  - [ ] Temperature monitor
    - [ ] CPU temp via sensors
    - [ ] GPU temp via sensors
    - [ ] Warning indicator when over threshold
  - [ ] Memory breakdown
    - [ ] Show used / free / cached / buffers
    - [ ] Swap usage %
    - [ ] Swap sparkline graph

---

- [ ] **Network Bandwidth**
  - [x] Real-time download speed graph
  - [x] Real-time upload speed graph
  - [ ] Per-interface breakdown
    - [ ] List all active interfaces
    - [ ] Show download/upload per interface
    - [ ] Switch between interfaces in TUI
  - [ ] Bandwidth per-process
    - [ ] Read from `/proc/net`
    - [ ] Map usage to process names
    - [ ] Sort by highest usage
  - [ ] Packet loss tracker
    - [ ] Ping multiple hosts simultaneously
    - [ ] Show loss % per host over time
    - [ ] Sparkline graph per host
  - [ ] Total data usage counter
    - [ ] Track total MB/GB sent and received since boot
    - [ ] Reset counter manually

---

- [ ] **MAC Spoofing**
  - [ ] List all available network interfaces
  - [ ] Select interface from TUI
  - [ ] Enter new MAC address
  - [ ] Validate MAC address format
  - [ ] Apply MAC change (requires `sudo`)
  - [ ] Store original MAC address
  - [ ] Restore original MAC address
  - [ ] Generate random valid MAC address
  - [ ] Show success/failure feedback in TUI

---

- [ ] **Device Scanner**
  - [ ] ARP scan to discover all devices on LAN
  - [ ] Display IP address per device
  - [ ] Display MAC address per device
  - [ ] Resolve and display hostname per device
  - [ ] Identify device vendor from MAC OUI
  - [ ] ARP table viewer
    - [ ] Read local ARP cache
    - [ ] Display in TUI table
    - [ ] Manual refresh keybind
  - [ ] Auto-refresh device list on interval
  - [ ] Flag new devices since last scan
  - [ ] Ping a selected device from TUI
  - [ ] Port scan a selected device from TUI

---

- [ ] **Network Tools**
  - [ ] Ping / latency monitor
    - [ ] Continuous ping to a target host
    - [ ] Live sparkline of response times
    - [ ] Show min / avg / max / jitter
    - [ ] Multi-host ping side by side
  - [ ] Port scanner
    - [ ] TCP port scan on a target host
    - [ ] UDP port scan on a target host
    - [ ] Show service names for known ports
    - [ ] Configurable port range
    - [ ] Export results to file
  - [ ] DNS lookup
    - [ ] Forward lookup (hostname → IP)
    - [ ] Reverse lookup (IP → hostname)
    - [ ] Show TTL
    - [ ] Show all record types (A, AAAA, MX, TXT)
    - [ ] Switchable DNS servers
  - [ ] Traceroute
    - [ ] Run traceroute to a target host
    - [ ] Display each hop with latency
    - [ ] Flag high-latency hops
  - [ ] Wi-Fi signal monitor
    - [ ] Detect if on wireless interface
    - [ ] Track RSSI over time
    - [ ] Sparkline graph of signal strength
    - [ ] Show connected SSID and channel
  - [ ] Firewall rule viewer
    - [ ] Read `iptables` rules
    - [ ] Read `nftables` rules
    - [ ] Display in readable TUI table

---

- [ ] **Open Connection Monitor**
  - [ ] List all active TCP/UDP connections
  - [ ] Show local and remote address per connection
  - [ ] Show connection state (ESTABLISHED, LISTEN, TIME_WAIT, etc.)
  - [ ] Map connections to process names
  - [ ] IP geolocation for remote addresses
    - [ ] Look up country/region for remote IPs
    - [ ] Display flag or country code in table
    - [ ] Flag connections to unusual regions
  - [ ] Filter by protocol, state, or process
  - [ ] Auto-refresh on tick

---

- [ ] **Process Table**
  - [ ] List all running processes
  - [ ] Show CPU % per process
  - [ ] Show RAM usage per process
  - [ ] Sort by CPU, RAM, or PID
  - [ ] Kill a process from the TUI
  - [ ] Filter / search processes by name
  - [ ] Show process owner (user)
  - [ ] Show process uptime
  - [ ] Show open file descriptor count
  - [ ] Tree view — parent/child process relationships
  - [ ] Auto-refresh on tick
  - [ ] Watch a named process — alert if it disappears

---

- [ ] **Alerts & Notifications**
  - [ ] CPU usage threshold alert
    - [ ] Configurable threshold %
    - [ ] Visual highlight in TUI when exceeded
    - [ ] Optional desktop notification
  - [ ] RAM usage threshold alert
    - [ ] Configurable threshold %
    - [ ] Visual highlight in TUI
  - [ ] Network speed threshold alert
    - [ ] Alert when download or upload exceeds limit
  - [ ] Temperature threshold alert
    - [ ] Alert when CPU/GPU temp exceeds limit
  - [ ] New device on LAN alert
    - [ ] Notify when unknown MAC joins network
  - [ ] Process crash alert
    - [ ] Watch a named process
    - [ ] Alert if it disappears from process table
  - [ ] Disk space alert
    - [ ] Configurable threshold %
    - [ ] Alert when disk is nearly full

---

- [ ] **SSH & Remote**
  - [ ] Known hosts viewer
    - [ ] Parse and display `~/.ssh/known_hosts`
    - [ ] Remove stale entries from TUI
  - [ ] SSH connection launcher
    - [ ] List saved SSH targets from config
    - [ ] Launch SSH session from TUI
    - [ ] Add / remove targets
  - [ ] Remote monitoring mode
    - [ ] Connect to a remote analytix instance
    - [ ] View remote system stats locally
    - [ ] Multi-host dashboard view

---

- [ ] **Tests**
  - [ ] Unit tests
    - [ ] `system/` — CPU, RAM, disk, network reads
    - [ ] `ui/` — graph rendering, history append
    - [ ] MAC spoof logic
    - [ ] ARP scanner
    - [ ] Port scanner
    - [ ] DNS lookup
    - [ ] Connection monitor
  - [ ] Integration tests
    - [ ] Full tick cycle test
    - [ ] TUI render smoke test
  - [ ] Benchmarks
    - [ ] Render performance under load
    - [ ] Network scan speed

---

- [ ] **Documentation**
  - [ ] `README.md`
    - [ ] Description and full feature list
    - [ ] Install instructions (binary + build from source)
    - [ ] Screenshot / demo gif
    - [ ] Keybindings reference table
    - [ ] Permissions note (`sudo` requirements)
  - [ ] `CONTRIBUTING.md`
    - [ ] How to add a new panel or tab
    - [ ] How to run tests
    - [ ] PR guidelines
    - [ ] Code style notes
  - [ ] Inline code comments on all exported functions
  - [ ] `man` page

---

## Low Priority

- [ ] **TUI & Navigation**
  - [ ] Tabbed interface
    - [ ] System tab
    - [ ] Network tab
    - [ ] Devices tab
    - [ ] Tools tab
    - [ ] Connections tab
    - [ ] Tab switching with arrow keys or number keys
  - [ ] Help panel
    - [ ] Press `?` to open overlay
    - [ ] List all keybindings
    - [ ] Dismiss with `?` or `esc`
  - [ ] Color themes
    - [ ] Default theme
    - [ ] Dark theme
    - [ ] Light theme
    - [ ] High contrast theme
    - [ ] Switchable at runtime
  - [ ] Mouse support
    - [ ] Click to switch tabs
    - [ ] Scroll through tables
  - [ ] Configurable layout
    - [ ] Reorder panels
    - [ ] Show/hide individual panels
    - [ ] Resize panels

---

- [ ] **Data & Config**
  - [ ] Export snapshots
    - [ ] Press `s` to export
    - [ ] JSON format output
    - [ ] CSV format output
    - [ ] Timestamped filename
  - [ ] Metric history logging
    - [ ] Write metrics to file on interval
    - [ ] Configurable output path
    - [ ] Rotate log files by size or date
  - [ ] Config file
    - [ ] Create `~/.config/analytix/config.toml`
    - [ ] Configurable refresh interval
    - [ ] Configurable default interface
    - [ ] Configurable theme
    - [ ] Configurable alert thresholds
  - [ ] CLI flag support
    - [ ] `--interface` flag
    - [ ] `--interval` flag
    - [ ] `--theme` flag
    - [ ] `--config` flag for custom config path
    - [ ] `--no-color` flag

---

- [ ] **Distribution**
  - [ ] Releases
    - [ ] Linux binary (amd64)
    - [ ] Linux binary (arm64)
    - [ ] macOS binary (amd64)
    - [ ] macOS binary (arm64 / Apple Silicon)
    - [ ] Windows binary
  - [ ] Packaging
    - [ ] `.deb` package
    - [ ] `.AppImage` package
    - [ ] AUR package (Arch Linux)
    - [ ] Homebrew formula (macOS)
    - [ ] Scoop formula (Windows)
  - [ ] CI/CD
    - [ ] GitHub Actions build on push
    - [ ] Run tests on PR
    - [ ] Auto release binaries on tag
    - [ ] Auto-generate changelog from commits

---

- [ ] **Battery Monitor**
  - [ ] Detect if battery is present
  - [ ] Show charge % with live sparkline
  - [ ] Show charging / discharging / full status
  - [ ] Estimated time remaining
  - [ ] Battery health %
  - [ ] Charge cycle count
  - [ ] Low battery alert
    - [ ] Configurable threshold %
    - [ ] Visual warning in TUI

---

- [ ] **Storage**
  - [ ] Disk usage breakdown
    - [ ] Scan a path and show size per directory
    - [ ] Sort by largest first
    - [ ] Drill down into subdirectories
    - [ ] Like `ncdu` but built into analytix
  - [ ] Filesystem mount viewer
    - [ ] List all mounted filesystems
    - [ ] Show usage % per mount
    - [ ] Show inode usage %
    - [ ] Highlight nearly full filesystems
  - [ ] Largest files finder
    - [ ] Scan a path recursively
    - [ ] Rank files by size
    - [ ] Show top N results in TUI
  - [ ] S.M.A.R.T. data viewer
    - [ ] Read S.M.A.R.T. stats per drive
    - [ ] Show reallocated sectors, read errors, spin retries
    - [ ] Health status indicator (OK / Warning / Failing)

---

- [ ] **Hardware Info**
  - [ ] System info panel
    - [ ] CPU model and core/thread count
    - [ ] Architecture
    - [ ] Kernel version
    - [ ] OS and distro name
    - [ ] Hostname
    - [ ] Uptime and last reboot reason
    - [ ] Like `neofetch` but live and built in
  - [ ] GPU stats
    - [ ] VRAM usage
    - [ ] GPU utilization %
    - [ ] Clock speeds
    - [ ] NVIDIA support via NVML
    - [ ] AMD support via sysfs
  - [ ] PCI device viewer
    - [ ] List connected PCI devices
    - [ ] Show device class and vendor
  - [ ] USB device viewer
    - [ ] List connected USB devices
    - [ ] Show vendor, product, and speed

---

- [ ] **Logs**
  - [ ] System log viewer
    - [ ] Tail `/var/log/syslog` live in TUI
    - [ ] `journalctl` integration
    - [ ] Pause / resume scrolling
    - [ ] Filter by keyword
    - [ ] Filter by severity level
  - [ ] Kernel message viewer
    - [ ] Live `dmesg` output
    - [ ] Severity level highlighting
    - [ ] Filter by severity
  - [ ] Boot time analyser
    - [ ] Read systemd journal for boot sequence
    - [ ] Show time each service took to start
    - [ ] Highlight slowest services
  - [ ] Failed services viewer
    - [ ] List all failed systemd services
    - [ ] Show last error message per service
    - [ ] Restart a service from TUI (requires `sudo`)

---

- [ ] **Security**
  - [ ] Failed login attempts
    - [ ] Parse `/var/log/auth.log`
    - [ ] Show recent failed SSH / sudo attempts
    - [ ] Show source IP and timestamp
  - [ ] Listening services viewer
    - [ ] List all daemons with open listeners
    - [ ] Show port, protocol, and process name
    - [ ] Highlight unexpected or unusual listeners
  - [ ] Sudoers viewer
    - [ ] Parse `/etc/sudoers`
    - [ ] Show which users have sudo access
    - [ ] Show sudo rules per user
  - [ ] File integrity checker
    - [ ] Watch a list of key system files
    - [ ] Alert on unexpected modification
    - [ ] Show which file changed and when
  - [ ] World-writable file scanner
    - [ ] Scan a path for world-writable files
    - [ ] Show results in TUI table
    - [ ] Flag files owned by root
  - [ ] SUID/SGID binary finder
    - [ ] Scan for SUID/SGID binaries
    - [ ] Highlight unusual or unexpected ones