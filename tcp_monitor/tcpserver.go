package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}
		go doServerStuff(conn)
	}
}


func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}


//
//package main
//
//import (
//"flag"
//"fmt"
//"net"
//"syscall"
//)
//
//const maxRead = 25
//
//func main() {
//	if flag.NArg() != 2 {
//		panic("usage: host port")
//	}
//	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
//	listener := initServer(hostAndPort)
//	for {
//		conn, err := listener.Accept()
//		checkError(err, "Accept: ")
//		go connectionHandler(conn)
//	}
//}
//func initServer(hostAndPort string) net.Listener {
//	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
//	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
//	listener, err := net.ListenTCP("tcp", serverAddr)
//	checkError(err, "ListenTCP: ")
//	println("Listening to: ", listener.Addr().String())
//	return listener
//}
//func connectionHandler(conn net.Conn) {
//	connFrom := conn.RemoteAddr().String()
//	println("Connection from: ", connFrom)
//	sayHello(conn)
//	for {
//		var ibuf []byte = make([]byte, maxRead+1)
//		length, err := conn.Read(ibuf[0:maxRead])
//		ibuf[maxRead] = 0 // to prevent overflow
//		switch err {
//		case nil:
//			handleMsg(length, err, ibuf)
//		case syscall.EAGAIN: // try again
//			continue
//		default:
//			goto DISCONNECT
//		}
//	}
//DISCONNECT:
//	err := conn.Close()
//	println("Closed connection: ", connFrom)
//	checkError(err, "Close: ")
//}
//func sayHello(to net.Conn) {
//	obuf := []byte{'L', 'e', 't', '\'', 's', ' ', 'G', 'O', '!', '\n'}
//	wrote, err := to.Write(obuf)
//	checkError(err, "Write: wrote "+string(wrote)+" bytes.")
//}
//func handleMsg(length int, err error, msg []byte) {
//	if length > 0 {
//		print("<", length, ":")
//		for i := 0; ; i++ {
//			if msg[i] == 0 {
//				break
//			}
//			fmt.Printf("%c", msg[i])
//		}
//		print(">")
//	}
//}
//func checkError(error error, info string) {
//	if error != nil {
//		panic("ERROR: " + info + " " + error.Error()) // terminate program
//	}
//}