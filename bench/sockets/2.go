package main

import (
	"bufio"
	"fmt"
	"time"
	"net"
	"os"
	"strconv"
)

func main() {

	var n = 100000 // Number of requests
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var conn, _ = net.Dial("tcp", "127.0.0.1:8080")

	var start = time.Now()

	for i := 0; i < n; i++ {
		// Send message
		var message = "Hello from client\n"
		fmt.Fprintf(conn, message)

		// Receive reply
		var newmessage, _ = bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Reply: ", string(newmessage))
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
