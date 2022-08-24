package main

import (
    "log"
    "net"
)

const dnsName = "baidu.com"

func main() {
    ips, err := net.LookupIP(dnsName)
    if err != nil {
        log.Fatalln(err)
    }

    log.Println(ips)
}
