package main

import (
    "fmt"
    "net"
    "strconv"
    "sync"
    "time"
    "os"
)

func portScan(wg *sync.WaitGroup, i int) {
    defer wg.Done()
    time.Sleep(3 * time.Millisecond)
    link := os.Args[1] // Changed from os.Args[2] to os.Args[1]
    address := link + ":" + strconv.Itoa(i) // Added ":" between
 the IP and port
    conn, err := net.Dial("tcp", address)
    if err == nil {
        fmt.Println("[+] Open port:", i)
        conn.Close()
    } else {
        // fmt.Println("[-] Closed port:", i)
    }
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run portscan.go <target-ip>")
        return
    }

    var wg sync.WaitGroup

    for i := 1; i <= 1024; i++ {
        wg.Add(1)
        go portScan(&wg, i)
    }

    wg.Wait()
}
