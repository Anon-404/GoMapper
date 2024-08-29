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
	fmt.Printf("%s, %s : To perform a port scan\n", color.GreenString("-d"),color.BlueString("--dnslookup"))

	fmt.Printf("\n%s, %s : Print the version number\n", color.GreenString("-v"),color.BlueString("--version"))
    
}

func version(){
	boldGreen.Println("GoMapper version 2.3")
}

func portScan(wg *sync.WaitGroup, domain string, port int) {
    defer wg.Done()

    service := "unknown"

    portsrv.PortSrv(port)
    service = portsrv.Service
    
    link := domain + ":" + strconv.Itoa(port)
    _, err := net.Dial("tcp", link)
    ports := strconv.Itoa(port) + "/tcp"
    if err == nil {
        fmt.Printf("%-12s %-9s %s\n", ports, "open", service)

    } else {
        // fmt.Println("Port closed:", port)
    }
}

func dnsLookup(domain string) {
	
	conn, err := net.LookupIP(domain)
    if err != nil {
    	fmt.Println(err)
        boldRed.Print("[-] Error : ",err)
    } else {
    
    	for _, ip := range conn {
        	ipv4 := ip.To4()
        	if ipv4 != nil {
            	fmt.Print("[+] IPv4 : ")
            	boldBlue.Print(ipv4,"\n")
        	}
 	   }
    }
    
    for _, ip := range conn {
        if ip.To4() == nil {
            fmt.Print("[+] IPv6 : ")
            boldBlue.Print(ip,"\n")
        }
    }

    cname, err := net.LookupCNAME(domain)
    if err != nil {
        fmt.Println(err)
        boldRed.Print("[-] Error :",err)
    } else {
    
    fmt.Print("[+] CNAME : ")
    boldBlue.Print(cname,"\n")

	NS, err := net.LookupNS("daudkandipiloths.edu.bd")
    	if err != nil {
        	fmt.Println("error :",err)
    	}
    	for _, ns := range NS {
        	fmt.Print("[+] Name Server(NS) : ")
			boldBlue.Print(ns.Host,"\n")
    	}
	}
	
	MX, err := net.LookupMX("google.com")
    	if err != nil {
        	fmt.Println("[-] Mail Exchange host(MX) : ", err)
    	} else {
    		for _, mx := range MX {
        		fmt.Print("[+] Mail Exchange host(MX) : ")
        		boldBlue.Print(mx.Host,"\n")
        		fmt.Print("[+] Mail Exchange preference(MX) : ")
        		boldBlue.Print(mx.Pref,"\n")
    		}
    	}

	TXT, err := net.LookupTXT(domain)
    if err != nil {
        boldRed.Println("[-] TXT Lookup Error:", err)
    } else {

    	fmt.Println("[+] TXT Records : ")
    	for i, txt := range TXT {
        	blue.Printf("   %d. %s\n", i+1, txt)

    	}
	}

	target, port, err := net.LookupSRV("sip", "tcp", domain)
    if err != nil {
    	boldRed.Println("[-] SRV Lookup Error:", err)
    } else {

    	fmt.Println("[+] SRV Record Found : ")
    	boldBlue.Printf("    Target: %s\n", target)
    	boldBlue.Printf("    Port: %d\n", port)
    	//boldBlue.Printf("    Priority: %d\n", priority)
    	//boldBlue.Printf("    Weight: %d\n", weight)

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
    	boldRed.Println("[-] Error :",err)
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
    } else if option == "-d" || option == "--dnslookup" {
    	dnsLookup(domain)	
    } else {

    	red := color.New(color.FgRed)
    	boldRed := red.Add(color.Bold)
    	boldRed.Println("[-] Unknown option",option)
        return
    }
}

