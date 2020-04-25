package main

import (
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var n = 30000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		time.Now()
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
