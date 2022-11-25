package main

import (
	"fmt"
	"net"
)

func telnet(host string, port int) {
	// connect to server
	address := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("address=<%s>\n", address)
	_, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("telnet failed")
		return
	}
	fmt.Println("telnet success")
}
