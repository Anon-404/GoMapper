package main

import (
    "fmt"
    "net"
    "strconv"
    "os"
    "sync"

    "github.com/fatih/color"
)

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
    fmt.Printf("\n%s, %s : To get the help page\n", color.BlueString("--help"),color.GreenString("-h"
))
    fmt.Printf("%s, %s : To perform a port scan\n", color.BlueString("--networkScan"),color.GreenStri
ng("-n"))

}

func portScan(wg *sync.WaitGroup, domain string, port int) {
    defer wg.Done()
    link := domain + ":" + strconv.Itoa(port)
    _, err := net.Dial("tcp", link)
    if err == nil {
        fmt.Println(port,"\b/tcp")
    } else {
        // fmt.Println("Port closed:", port)
    }
}

func main() {
    banner()

    if len(os.Args) < 2 {
        fmt.Println("Usage: GoMapper <option> [domain/ip]")
        fmt.Println("Manual: GoMapper -h")
        return
    }

    option := os.Args[1]
    if option == "-h" || option == "--help" {
        hlpMnu()
        return
    }

    if len(os.Args) < 3 {
        fmt.Println("Error: Domain/IP missing")
        fmt.Println("Usage: GoMapper <option> <domain/ip>")
        fmt.Println("Manual: GoMapper -h")
        return
    }

    domain := os.Args[2]

    if option == "-n" || option == "--networkScan" {
        var wg sync.WaitGroup
        fmt.Print("[+] Scanning ")
        blue := color.New(color.FgBlue)
        boldBlue := blue.Add(color.Bold)
        boldBlue.Print(domain,"\n\n")
        for port := 1; port <= 65535; port++ {
            wg.Add(1)
            go portScan(&wg, domain, port)
        }
        wg.Wait()
    } else {

        red := color.New(color.FgRed)
        boldRed := red.Add(color.Bold)
        boldRed.Println("[-] Unknown option",option)
        return
    }
