package testing

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"testing"
)

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Println("Error Dialing", err.Error())
		return
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name")
	clientname, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientname, "\n")
	for {
		fmt.Println("what to send to the server type Q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")
		if trimmedInput == "Q" {
			return
		}
		_, err := conn.Write([]byte(trimmedClient + "says: " + trimmedInput))
		if err != nil {
			fmt.Println("Error sending", err.Error())
			return
		}
	}
}

