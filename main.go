package main

import (
    "fmt"
    "net"
    "strconv"
    "os"
    "sync"
    "syscall"

    "github.com/fatih/color"
    "github.com/Anon-404/GoMapper/src"
)

// color blur

var blue = color.New(color.FgBlue)
var boldBlue = blue.Add(color.Bold)

// color red

var red = color.New(color.FgRed)
var boldRed = red.Add(color.Bold)

// color green

var Green = color.New(color.FgGreen)
var boldGreen = Green.Add(color.Bold)

func banner() {
    banner := `
   ______      __  ___                           
  / ____/___  /  |/  /___ _____  ____  ___  _____
 / / __/ __ \/ /|_/ / __ \/ __ \/ __ \/ _ \/ ___/
/ /_/ / /_/ / /  / / /_/ / /_/ / /_/ /  __/ /    
\____/\____/_/  /_/\__,_/ .___/ .___/\___/_/     
	               /_/   /_/                 
`
    cyan := color.New(color.FgCyan)

    boldCyan := cyan.Add(color.Bold)
    boldCyan.Println(banner)
}

func hlpMnu() {
    fmt.Println("Usage: GoMapper <option> <domain/ip>")
    fmt.Println("\nOptions:")
	fmt.Printf("\n%s, %s : To get the help page\n", color.GreenString("-h"),color.BlueString("--help"))
	fmt.Printf("%s, %s : To perform a port scan\n", color.GreenString("-n"),color.BlueString("--networkScan"))

	fmt.Printf("\n%s, %s : Print the version number\n", color.GreenString("-v"),color.BlueString("--version"))
    
}

func version(){
	boldGreen.Println("GoMapper version 1.0")
}

func portScan(wg *sync.WaitGroup, domain string, port int) {
    defer wg.Done()

    service := "unknown"

    portsrv.PortSrv(port)
    service = portsrv.Service

	/*
    if port == 80 {
        service = "http"
    } else if port == 53 {
        service = "domain"
    } else if port == 443 {
        service = "https"
    } else if port == 22 {
        service = "ssh"
    } else if port == 21 {
        service = "ftp"
    } else if port == 23 {
        service = "telnet"
    } else if port == 25 {
        service = "smtp"
    } else if port == 123 {
        service = "ntp"
    } else if port == 2086 {
		service = "cPanel(non-SSL)"
    } else if port == 2087 {
		service = "cPanel(SSL)"
    } else if port == 2052 {
		service = "cloudflare(http-alt)"
    } else if port == 2053 {
		service = "cloudflare(https-alt)"
    } else if port == 2096 {
		service = "cPanel-Webmail(SSL)"
    } else if port == 2095 {
		service = "cPanel-Webmail(non-SSL)"
    } else if port == 445 {
        service = "smb"
    } else if port == 67 {
        service = "dhcp"
    } else if port == 110 {
        service = "pop3"
    } else if port == 143 {
        service = "imap"
    } else if port == 465 {
        service = "smtps"
    } else if port == 990 {
        service = "ftps"
    } else if port == 389 {
        service = "ldap"
    } else if port == 636 {
        service = "ldaps"
    } else if port == 3306 {
        service = "mysql"
    } else if port == 5432 {
        service = "postgresql"
    } else if port == 3389 {
        service = "rdp"
    } else if port == 6379 {
        service = "redis"
    } else if port == 27017 {
        service = "mongodb"
    } else if port == 9200 {
        service = "elasticsearch"
    } else if port == 11211 {
        service = "memcached"
    } else if port == 161 {
        service = "snmp"
    } else if port == 179 {
        service = "bgp"
    } else if port == 88 {
        service = "kerberos"
    } else if port == 500 {
        service = "ike"
    } else if port == 5060 {
        service = "sip"
    } else if port == 1883 {
        service = "mqtt"
    } else if port == 5900 {
        service = "vnc"
    } else if port == 69 {
        service = "tftp"
    } else if port == 5222 {
        service = "xmpp"
    } else if port == 9418 {
        service = "git"
    } else if port == 873 {
        service = "rsync"
    } else if port == 2049 {
        service = "nfs"
    } else if port == 3128 {
        service = "squid"
    } else if port == 5901 {
        service = "vnc"
    } else if port == 3000 {
        service = "node"
    } else if port == 7000 {
        service = "afp"
    } else if port == 9000 {
        service = "php-fpm"
    } else if port == 3690 {
        service = "svn"
    } else if port == 4444 {
        service = "metasploit"
    } else if port == 5061 {
        service = "sips"
    } else if port == 7080 {
        service = "http-proxy"
    } else if port == 8443 {
        service = "https-alt"
    } else if port == 2022 {
        service = "bitbucket"
    } else if port == 25565 {
        service = "minecraft"
    } else if port == 6667 {
        service = "irc"
    } else if port == 10000 {
        service = "webmin"
    } else if port == 49152 {
        service = "dynamic"
    } else if port == 5000 {
        service = "upnp"
    } else if port == 6000 {
        service = "x11"
    } else if port == 8000 {
        service = "http-alt"
    } else if port == 8081 {
        service = "http-alt"
    } else if port == 8888 {
        service = "http-alt"
    } else if port == 7001 {
        service = "weblogic"
    } else if port == 4848 {
        service = "glassfish"
    } else if port == 5433 {
        service = "postgresql"
    } else if port == 8880 {
        service = "http-alt"
    } else if port == 6378 {
        service = "redis"
    } else if port == 9300 {
        service = "elasticsearch"
    } else if port == 9201 {
        service = "elasticsearch"
    } else if port == 2181 {
        service = "zookeeper"
    } else if port == 11214 {
        service = "memcached"
    } else if port == 8082 {
        service = "http-alt"
    } else if port == 27018 {
        service = "mongodb"
    } else if port == 50000 {
        service = "sap"
    } else if port == 2048 {
        service = "oracle"
    } else if port == 9500 {
        service = "apache"
    } else if port == 50030 {
        service = "hadoop"
    } else if port == 50070 {
        service = "hadoop"
    } else if port == 50090 {
        service = "hadoop"
    } else if port == 61616 {
        service = "activemq"
    } else if port == 8181 {
        service = "websphere"
    } else if port == 9090 {
        service = "websphere"
    } else if port == 7474 {
        service = "neo4j"
    } else if port == 6001 {
        service = "x11"
    } else if port == 28017 {
        service = "mongodb"
    } else if port == 9202 {
        service = "elasticsearch"
    } else if port == 9301 {
        service = "elasticsearch"
    } else if port == 5001 {
        service = "kafka"
    } else if port == 6002 {
        service = "x11"
    } else if port == 3700 {
        service = "cassandra"
    } else if port == 8444 {
        service = "solr"
    } else if port == 5902 {
        service = "vnc"
    } else if port == 8085 {
        service = "http-alt"
    } else if port == 8086 {
        service = "influxdb"
    } else if port == 8090 {
        service = "webadmin"
    } else if port == 8087 {
        service = "freepbx"
    } else if port == 9091 {
        service = "webmin"
    } else if port == 3127 {
        service = "backdoor"
    } else if port == 2222 {
        service = "ssh"
    } else if port == 37015 {
        service = "cassandra"
    } else if port == 37777 {
        service = "torrent"
    } else if port == 15555 {
        service = "backdoor"
    } else if port == 31000 {
        service = "backdoor"
    } else if port == 49153 {
        service = "dynamic"
    } else if port == 49154 {
        service = "dynamic"
    } else if port == 49155 {
        service = "dynamic"
    } else if port == 49156 {
        service = "dynamic"
    } else if port == 49157 {
        service = "dynamic"
    } else if port == 5431 {
        service = "db2"
    } else if port == 2181 {
        service = "zookeeper"
    } else if port == 16384 {
        service = "epmap"
    } else if port == 2266 {
        service = "ebgp"
    } else if port == 3846 {
        service = "bmc"
    } else if port == 3715 {
        service = "dawn"
    } else if port == 3701 {
        service = "smtp"
    } else if port == 5480 {
        service = "admin"
    } else if port == 6464 {
        service = "flash"
    } else if port == 8088 {
        service = "http-alt"
    } else if port == 9092 {
        service = "kafka"
    } else if port == 8445 {
        service = "apache"
    } else if port == 8008 {
        service = "http-alt"
    } else if port == 6100 {
        service = "x11"
    } else if port == 4445 {
        service = "backdoor"
    } else if port == 5353 {
        service = "mdns"
    } else if port == 6881 {
        service = "torrent"
    } else if port == 33060 {
        service = "mysqlx"
    } else if port == 20048 {
        service = "nfs3"
    } else if port == 135 {
        service = "msrpc"
    } else if port == 162 {
        service = "snmptrap"
    } else if port == 5431 {
        service = "db2"
    } else if port == 8001 {
        service = "http-alt"
    } else if port == 7547 {
        service = "cwmp"
    } else if port == 50001 {
        service = "sap"
    } else if port == 5701 {
        service = "hazelcast"
    } else if port == 1514 {
        service = "syslog"
    } else if port == 8002 {
        service = "http-alt"
    } else if port == 7575 {
        service = "couchbase"
    } else if port == 6882 {
        service = "torrent"
    } else if port == 1099 {
        service = "rmi"
    } else if port == 2017 {
        service = "robo"
    } else if port == 8099 {
        service = "bigfix"
    } else if port == 9929 {
    	service = "rsync"
    } else if port ==  31337 {
    	service = "back-orifice"
    } else {
        service = "unknown"
    }
    */
    
    link := domain + ":" + strconv.Itoa(port)
    _, err := net.Dial("tcp", link)
    ports := strconv.Itoa(port) + "/tcp"
    if err == nil {
        fmt.Printf("%-12s %-9s %s\n", ports, "open", service)
    } else {
        // fmt.Println("Port closed:", port)
    }
}

func detectOS(domain string) {
    // Dial a TCP connection to the target on port 80 (HTTP)
    conn, err := net.Dial("tcp", domain+":80")
    if err != nil {
        fmt.Println("Connection error:", err)
        return
    }
    defer conn.Close()

    // Get the file descriptor of the connection
    rawConn, err := conn.(*net.TCPConn).SyscallConn()
    if err != nil {
        fmt.Println("Error accessing raw connection:", err)
        return
    }

    var ttl int
    err = rawConn.Control(func(fd uintptr) {
        // Get TTL value using syscall
        ttl, err = syscall.GetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.
IP_TTL)
    })
    if err != nil {
        fmt.Println("Error getting TTL:", err)
        return
    }

    // fmt.Println("TTL value:", ttl)
    if ttl == 64 {
        fmt.Print("\n[+] OS type:")
        boldBlue.Print(" Linux/Unix (guessing) \n")
    } else if ttl == 128 {
        fmt.Print("\n[+] OS type:")
        boldBlue.Print(" Windows (guessing) \n")
    } else if ttl == 255 {
        fmt.Print("\n[+] OS type:")
        boldBlue.Print(" Cisco-IOS (guessing) \n")
    } else {
        fmt.Print("\n[+] OS type:")
        boldRed.Print(" Windows Unknown ! \n")
    }
}

func main() {

/*
	// color blur 
	
	blue := color.New(color.FgBlue)
	boldBlue := blue.Add(color.Bold)

	// color red

	red := color.New(color.FgRed)
	boldRed := red.Add(color.Bold)
*/
	
    banner()

    if len(os.Args) < 2 {
        boldBlue.Println("Usage: GoMapper <option> [domain/ip]")
        boldBlue.Println("Manual: GoMapper -h")
        boldBlue.Println("Version: GoMapper -v")
        return
    }

    option := os.Args[1]
    if option == "-h" || option == "--help" {
        hlpMnu()
        return
    } else if option == "-v" || option == "--version" {
    	version()
    	return
    }

    if len(os.Args) < 3 {
        boldRed.Println("Error: Domain/IP missing")
        boldBlue.Println("Usage: GoMapper <option> <domain/ip>")
        boldBlue.Println("Manual: GoMapper -h")
        boldBlue.Println("Version: GoMapper -v")
        return
    }

    domain := os.Args[2]

    ip, err := net.LookupIP(domain)
    if err != nil {
    	boldRed.Println("[-] poor network or host not found")
    	// fmt.Println("[-] check internet connection")
    	return
    }

    if option == "-n" || option == "--networkScan" {
        var wg sync.WaitGroup
        fmt.Print("[+] Scanning: ")
        boldBlue.Print(domain,"\n")

	var ipv4s, ipv6s []string

	for _, addr := range ip {
		if addr.To4() != nil {
			ipv4s = append(ipv4s, addr.String())
		} else if addr.To16() != nil {
			ipv6s = append(ipv6s, addr.String())
		}
	}

	fmt.Print("[+] IPv4 Address: ")
	for _, addr := range ipv4s {
		boldBlue.Print(addr,"\n")
	}

	fmt.Print("[+] IPv6 Address: ")
	for _, addr := range ipv6s {
		boldBlue.Print(addr,"\n")
	}
	if ipv6s == nil {
		fmt.Println("\n ")
	} else {
		fmt.Println(" ")
	}
        
        for port := 1; port <= 65535; port++ {
            wg.Add(1)
            go portScan(&wg, domain, port)
        }
        wg.Wait()
        detectOS(domain)
    } else {

    	red := color.New(color.FgRed)
    	boldRed := red.Add(color.Bold)
    	boldRed.Println("[-] Unknown option",option)
        return
    }
}

