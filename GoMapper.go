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
    "github.com/likexian/whois"
    "github.com/likexian/whois-parser"
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

// color yellow

var Yellow = color.New(color.FgYellow)
var boldYellow = Yellow.Add(color.Bold)

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
    boldGreen.Println("William Steven(Anon404)\n")
}

func hlpMnu() {
    fmt.Println("Usage: GoMapper <option> <domain/ip>")
    fmt.Println()

    fmt.Println("Options:")
    fmt.Printf("  %s, %s [domain] : %s\n", color.GreenString("-a"), color.BlueString("--all"), "Perform all actions")
    fmt.Printf("  %s, %s [domain] : %s\n", color.GreenString("-n"), color.BlueString("--networkScan"), "Perform a network scan (IP, Port, OS Ditection)")
    fmt.Printf("  %s, %s [domain] : %s\n", color.GreenString("-d"), color.BlueString("--dnslookup"), "Perform a DNS lookup")
    fmt.Printf("  %s, %s [domain] : %s\n", color.GreenString("-w"), color.BlueString("--whoislookup"), "Perform a WHOIS lookup")

    fmt.Println()
    fmt.Println("Other options:")
    fmt.Printf("  %s, %s : %s\n", color.GreenString("-h"), color.BlueString("--help"), "Get the help page")
    fmt.Printf("  %s, %s : %s\n", color.GreenString("-v"), color.BlueString("--version"), "Print the version number")
}

func version(){
	boldYellow.Println("GoMapper version 3.2")
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

	NS, err := net.LookupNS(domain)
    	if err != nil {
        	fmt.Println("error :",err)
    	}
    	for _, ns := range NS {
        	fmt.Print("[+] Name Server(NS) : ")
			boldBlue.Print(ns.Host,"\n")
    	}
	}
	
	MX, err := net.LookupMX(domain)
    	if err != nil {
        	fmt.Println("[-] Mail Exchange host(MX) : \n", err)
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

func whoisLkpDom(domain string) {

    conn, err := whois.Whois(domain)
    if err != nil {
        boldRed.Println("[-] Error:", err)
        return
    }

    parsedResponse, err := whoisparser.Parse(conn)
    if err != nil {
        boldRed.Println("[-] Error parsing WHOIS data:", err)
        boldRed.Println("Raw WHOIS Response:", conn)
        return
    }

    boldGreen.Println("############### [Domain Information] ###############\n")

    // Domain Information
    fmt.Print("[+] Domain Name : ")
    boldBlue.Print(parsedResponse.Domain.Domain,"\n")

	fmt.Print("[+] Status : ")
    boldBlue.Print(parsedResponse.Domain.Status,"\n")

	fmt.Print("[+] DNSSEC : ")
    boldBlue.Print(parsedResponse.Domain.DNSSec,"\n")

	fmt.Print("[+] Domain Name : ")
    boldBlue.Print(parsedResponse.Domain.Domain,"\n")

    fmt.Print("[+] Whois Server : ")
    boldBlue.Print(parsedResponse.Domain.WhoisServer,"\n")

    fmt.Print("[+] Creation Date : ")
    boldBlue.Print(parsedResponse.Domain.CreatedDate,"\n")

    fmt.Print("[+] Expiration Date : ")
    boldBlue.Print(parsedResponse.Domain.ExpirationDate,"\n")

    fmt.Print("[+] Update Date : ")
    boldBlue.Print(parsedResponse.Domain.UpdatedDate,"\n")

    fmt.Print("[+] Name Server[s] : ")
    boldBlue.Print(parsedResponse.Domain.NameServers,"\n")
}

func whoisLkpReg(domain string) {

    conn, err := whois.Whois(domain)
    if err != nil {
        boldRed.Println("[-] Error:", err)
        return
    }

    parsedResponse, err := whoisparser.Parse(conn)
    if err != nil {
        boldRed.Println("[-] Error parsing WHOIS data:", err)
        boldRed.Println("Raw WHOIS Response:", conn)
        return
    }

    boldGreen.Println("\n############### [Registrar Information] ###############\n")

    // Registrar Information
    fmt.Print("[+] Registrar Name : ")
    boldBlue.Print(parsedResponse.Registrar.Name, "\n")

    fmt.Print("[+] Registrar ID : ")
    boldBlue.Print(parsedResponse.Registrar.ID, "\n")

    fmt.Print("[+] Registrar Phone : ")
    boldBlue.Print(parsedResponse.Registrar.Phone, "\n")

    fmt.Print("[+] Registrar Email : ")
    boldBlue.Print(parsedResponse.Registrar.Email, "\n")

    boldGreen.Println("\n############### [Registrant Information] ###############\n")

    // Registrant Information
    fmt.Print("[+] Registrant Name : ")
    boldBlue.Print(parsedResponse.Registrant.Name, "\n")

    fmt.Print("[+] Registrant Organization : ")
    boldBlue.Print(parsedResponse.Registrant.Organization, "\n")

    fmt.Print("[+] Registrant Street : ")
    boldBlue.Print(parsedResponse.Registrant.Street, "\n")

    fmt.Print("[+] Registrant City : ")
    boldBlue.Print(parsedResponse.Registrant.City, "\n")

    fmt.Print("[+] Registrant Postal Code : ")
    boldBlue.Print(parsedResponse.Registrant.PostalCode, "\n")

    fmt.Print("[+] Registrant Country : ")
    boldBlue.Print(parsedResponse.Registrant.Country, "\n")

    fmt.Print("[+] Registrant Phone : ")
    boldBlue.Print(parsedResponse.Registrant.Phone, "\n")

    fmt.Print("[+] Registrant Email : ")
    boldBlue.Print(parsedResponse.Registrant.Email, "\n")
}
    
func whoisLkpTech(domain string) {

    conn, err := whois.Whois(domain)
    if err != nil {
        boldRed.Println("[-] Error:", err)
        return
    }

    parsedResponse, err := whoisparser.Parse(conn)
    if err != nil {
        boldRed.Println("[-] Error parsing WHOIS data:", err)
        boldRed.Println("Raw WHOIS Response:", conn)
        return
    }

    boldGreen.Println("\n############### [Technical Information] ###############\n")

    // Technical Information
    fmt.Print("[+] Tech Name : ")
    boldBlue.Print(parsedResponse.Technical.Name, "\n")

    fmt.Print("[+] Tech Organization : ")
    boldBlue.Print(parsedResponse.Technical.Organization, "\n")

    fmt.Print("[+] Tech Street : ")
    boldBlue.Print(parsedResponse.Technical.Street, "\n")

    fmt.Print("[+] Tech City : ")
    boldBlue.Print(parsedResponse.Technical.City, "\n")

    fmt.Print("[+] Tech Postal Code : ")
    boldBlue.Print(parsedResponse.Technical.PostalCode, "\n")

    fmt.Print("[+] Tech Country : ")
    boldBlue.Print(parsedResponse.Technical.Country, "\n")

    fmt.Print("[+] Tech Phone : ")
    boldBlue.Print(parsedResponse.Technical.Phone, "\n")

    fmt.Print("[+] Tech Email : ")
    boldBlue.Print(parsedResponse.Technical.Email, "\n")
}

func whoisLkpBuild(domain string) {

    conn, err := whois.Whois(domain)
    if err != nil {
        boldRed.Println("[-] Error:", err)
        return
    }

    parsedResponse, err := whoisparser.Parse(conn)
    if err != nil {
        boldRed.Println("[-] Error parsing WHOIS data:", err)
        boldRed.Println("Raw WHOIS Response:", conn)
        return
    }

    boldGreen.Println("\n############### [Billing Information] ###############\n")

    if parsedResponse.Billing != nil {
    
        fmt.Print("[+] Billing Name : ")
        boldBlue.Print(parsedResponse.Billing.Name, "\n")

        fmt.Print("[+] Billing Organization : ")
        boldBlue.Print(parsedResponse.Billing.Organization, "\n")

        fmt.Print("[+] Billing Street : ")
        boldBlue.Print(parsedResponse.Billing.Street, "\n")

        fmt.Print("[+] Billing City : ")
        boldBlue.Print(parsedResponse.Billing.City, "\n")

        fmt.Print("[+] Billing Postal Code : ")
        boldBlue.Print(parsedResponse.Billing.PostalCode, "\n")

        fmt.Print("[+] Billing Country : ")
        boldBlue.Print(parsedResponse.Billing.Country, "\n")

        fmt.Print("[+] Billing Phone : ")
        boldBlue.Print(parsedResponse.Billing.Phone, "\n")

        fmt.Print("[+] Billing Email : ")
        boldBlue.Print(parsedResponse.Billing.Email, "\n")
    } else {
        boldYellow.Println("[-] Billing information not available in the WHOIS response.")
    }
}

func detectOS(domain string) {
    
    conn, err := net.Dial("tcp", domain+":80")
    if err != nil {
        fmt.Println("Connection error:", err)
        return
    }
    defer conn.Close()

    rawConn, err := conn.(*net.TCPConn).SyscallConn()
    if err != nil {
        fmt.Println("Error accessing raw connection:", err)
        return
    }

    var ttl int
    err = rawConn.Control(func(fd uintptr) {
        
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
    boldBlue.Print(domain, "\n")

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
        boldBlue.Print(addr, "\n")
    }

    fmt.Print("[+] IPv6 Address: ")
    for _, addr := range ipv6s {
        boldBlue.Print(addr, "\n")
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

} else if option == "-w" || option == "--whoislookup" {
    whoisLkpDom(domain)
    whoisLkpReg(domain)
    whoisLkpTech(domain)
    whoisLkpBuild(domain)

} else if option == "-a" || option == "--all" {
    // Perform all actions

    var wg sync.WaitGroup

    // Network Scan
    fmt.Print("[+] Scanning: ")
    boldBlue.Print(domain, "\n")

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
        boldBlue.Print(addr, "\n")
    }

    fmt.Print("[+] IPv6 Address: ")
    for _, addr := range ipv6s {
        boldBlue.Print(addr, "\n")
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

    // DNS Lookup
    dnsLookup(domain)

    // WHOIS Lookup
    whoisLkpDom(domain)
    whoisLkpReg(domain)
    whoisLkpTech(domain)
    whoisLkpBuild(domain)

} else {
    red := color.New(color.FgRed)
    boldRed := red.Add(color.Bold)
    boldRed.Println("[-] Unknown option", option)
    return
}
}
