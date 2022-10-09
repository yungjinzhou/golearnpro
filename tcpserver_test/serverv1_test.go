package testing

import (
	"fmt"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	fmt.Println("Starting tht server....")
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error Reading", err.Error())
			return
		}
		fmt.Printf("Received data : %v", string(buf[:len]))
	}
}


