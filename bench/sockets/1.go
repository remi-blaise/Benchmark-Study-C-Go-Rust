package main

import (
	"bufio"
	"fmt"
	"io"
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
		var _, err = bufio.NewReader(conn).ReadString('\n')
		if err == io.EOF {
			return
		}

		// Reply
		var newmessage = "Hello from server\n"
		fmt.Fprintf(conn, newmessage)
	}
}
