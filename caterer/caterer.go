package main

import (
    "fmt"
    "golang/ip"
    "golang/node"
)


func main() {
    fmt.Println("Caterer")

    n := node.New("foo.bar.com")
    n.Eth0 = ip.New("10.0.0.1")
    n.Eth1 = ip.New("10.1.0.1")

    fmt.Printf("node = %s\n", n)
}

