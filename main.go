package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	CONNECT := "localhost:1337"
	connection, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(connection, text+"\n")

		message, _ := bufio.NewReader(connection).ReadString('\n')
		fmt.Print("->: " + message)
		if message == "exit" {
			fmt.Println("TCP client exiting...")
			connection.Close()
			return
		}
	}
}
