package main

import (
	"os"
	"strconv"
	"fmt"
	"time"
)

// Recursion function
func fib(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	var n uint64 = 40
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.ParseUint(args[0], 10, 64)
	}

	var start = time.Now()

	n = fib(n)

	println(n)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
