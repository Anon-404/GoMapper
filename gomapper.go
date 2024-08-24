package main

import (
    "fmt"
    "net/http"
    "strings"
    "io/ioutil"
)

func postrqst(){

    link := "https://github.com"
/*
    jsnDta := strings.NewReader(`
        {
            "Name":"William",
            "Age":18,
            "Sec":"Penetration testing"
        }
    `)
*/
    resp, err := http.Get(link)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    data, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))

}

func main(){

    link := "https://github.com"
  
    json := strings.NewReader(`
        {
            "Name":"William",
            "Age":18
        }
    `)
    rqst, err := http.Post(link,"application/json",json)
  
    if err != nil {
        panic(err)
    }

    defer rqst.Body.Close()

    resp, err := ioutil.ReadAll(rqst.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(resp))
}
