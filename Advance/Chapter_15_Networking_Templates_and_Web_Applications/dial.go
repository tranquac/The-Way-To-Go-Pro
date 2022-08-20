package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000") // tcp ipv4
	checkConnection(conn, err)
	conn, err = net.Dial("udp", "127.0.0.1:8000") // udp
	checkConnection(conn, err)
	conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:8000") // tcp ipv6
	checkConnection(conn, err)
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!\n", err)
		return
	}
	fmt.Printf("Connection is made with %v\n", conn)
}
