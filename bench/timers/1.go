package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var b = 1_000_000_000 // 1 second
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()
	var end = time.Now()
	var counter = 0
	for end.Sub(start) < n {
		end = time.Now()
		counter++
	}
	fmt.Println(counter)
}
