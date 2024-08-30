
# GoMapper 🚀
**Enumeration and Network Scanner Tool**  
*Written in Go*

<div align="center">
  <img src="https://raw.githubusercontent.com/Anon-404/My-assets/main/GoMapper/GoMapper.jpg" alt="GoMapper Logo" width="200"/>
</div>

## Features ✨
- 🔍 **Network Scanning and OS Detection**
- ⚡ **Faster Scanning:** Scans 65,535 ports quicker than Nmap
- 🌐 **DNS Lookup**
- 🕵️‍♂️ **Whois Lookup**
- 🎯 **User-Friendly Interface**
- 📖 **Comprehensive User Manual**

## Installation 🛠️

### Method 1: Using `go install`

#### Step 1: Install Go
- **Arch-based Linux:**
  ```bash
  sudo pacman -S go
  ```
- **Debian-based Linux:**
  ```bash
  sudo apt install go -y
  ```
- **Fedora:**
  ```bash
  sudo dnf install go
  ```
- **Termux:**
  ```bash
  pkg install golang -y
  ```

#### Step 2: Install GoMapper
- **For Linux:**
  ```bash
  go install -v github.com/Anon-404/GoMapper@latest
  sudo cp $HOME/go/bin/GoMapper /usr/bin
  ```
- **For Termux:**
  ```bash
  go install -v github.com/Anon-404/GoMapper@latest
  cp $HOME/go/bin/GoMapper ../usr/bin
  ```

### Method 2: Cloning the Repository

#### Step 1: Clone and Build
- **For Linux:**
  ```bash
  git clone https://github.com/Anon-404/GoMapper
  cd GoMapper
  go build -o GoMapper
  sudo cp GoMapper /usr/bin
  ```
- **For Termux:**
  ```bash
  git clone https://github.com/Anon-404/GoMapper
  cd GoMapper
  go build -o GoMapper
  cp GoMapper ../../usr/bin
  ```

## Usage 🧑‍💻

```bash
GoMapper <option> <domain/ip>
```

### Main Options:
- **`-a`, `--all [domain]`**  
  🔗 **Perform all actions**  
  Executes network scan, DNS lookup, and WHOIS lookup in a single command.

- **`-n`, `--networkScan [domain]`**  
  🌐 **Network Scan**  
  Performs a thorough network scan, including IP, port scanning, and OS detection.

- **`-d`, `--dnslookup [domain]`**  
  🛠 **DNS Lookup**  
  Retrieves detailed DNS records for the specified domain.

- **`-w`, `--whoislookup [domain]`**  
  🔍 **Whois Lookup**  
  Gathers WHOIS registration data for the specified domain.

### Additional Options:
- **`-h`, `--help`**  
  📝 **Help**  
  Displays this help page with descriptions of all available commands.

- **`-v`, `--version`**  
  🆚 **Version**  
  Prints the current version number of GoMapper.


## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.
