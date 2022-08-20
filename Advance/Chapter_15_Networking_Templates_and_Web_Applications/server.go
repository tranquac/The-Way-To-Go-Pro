package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Starting the server...")
	// create a listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening ", err.Error())
	}

	// listen for  and accept connections from clients
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting ", err.Error())
			return // terminate the program
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn){
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading ", err.Error())
			return // terminate the program
		}
		fmt.Printf("Recieved data: %v", string(buf[:len]))
	}
}
