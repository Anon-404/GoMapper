
# GoMapper 🚀
**Enumeration and Network Scanner Tool**  
*Written in Go*

![GoMapper Logo](https://raw.githubusercontent.com/Anon-404/My-assets/main/GoMapper/GoMapper.jpg)

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
  pkg install go -y
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
For detailed usage instructions, please refer to the [User Manual](link_to_manual).

## Contributions 🤝
Contributions are welcome! Feel free to open issues or submit pull requests.

## License 📜
This project is licensed under the MIT License.

```

This version consolidates all steps into a single markdown file. You can replace `link_to_manual` with an actual link if needed.
