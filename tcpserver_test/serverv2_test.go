package testing

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)
const maxRead = 25


func testTcpServerV2() {
	flag.Parse()
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
		hostAndPort := fmt.Sprintf("%s %s", flag.Arg(0), flag.Arg(1))
		listener := initServer(hostAndPort)
		for {
			conn, err := listener.Accept()
			checkError(err, "Accept: ")
			go connectionHandler(conn)

		}
}


func initServer(hostAndPort string) *net.TCPListener {
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address: port failed: `"+hostAndPort+"`")
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "lisentcp: ")
	return listener
}


func connectionHandler(conn net.Conn) {
	connFrom := conn.RemoteAddr().String()
	println("Connection from: ", connFrom)
	sayHello(conn)
	for {
		ibuf := make([]byte, maxRead+1)
		length, err := conn.Read(ibuf[0:maxRead])
		ibuf[maxRead] = 0 // to prevent overflow
		switch err {
		case nil:
			HandleMsg(length, err, ibuf)
		case syscall.EAGAIN: // try again
			continue
		default:
			goto DISCONNECT
		}
	}
DISCONNECT:
	err := conn.Close()
	println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	obuf := []byte{'L', 'E', 'T', '\'', 'S', ' ', 'G', 'O', '!', '\n'}
	wrote, err := to.Write(obuf)
	fmt.Println(wrote)
	checkError(err, "Writexxxx")
}


func HandleMsg(length int, err error, msg []byte) {
	if length > 0 {
		print("<", length, ":")
		for i := 0;  ; i++ {
			if msg[i] == 0 {
				break
			}
			fmt.Printf("%c", msg[i])
			fmt.Println(err.Error())
		}
	}
}


func checkError(error error, info string){
	if error != nil {
		panic("ERROR: " + info + " " + error.Error())
	}
}
