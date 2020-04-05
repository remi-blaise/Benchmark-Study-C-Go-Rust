package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen
	var ln, _ = net.Listen("tcp", ":8080")

	// Accept connection
	var conn, _ = ln.Accept()

	// Run loop forever
	for {
		// Will listen for message to process ending in newline (\n)
		var message, _ = bufio.NewReader(conn).ReadString('\n')
		if message != "" {
			fmt.Print("Message received:", string(message))
		}

		// Reply
		var newmessage = "Hello from server\n"
		fmt.Fprintf(conn, newmessage)
	}
}
