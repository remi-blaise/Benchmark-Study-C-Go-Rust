package main

import (
	"os"
	"strconv"
)

// Recursion function
func fib(n uint64) uint64 {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func main() {
	var n uint64 = 100
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.ParseUint(args[0], 10, 64)
	}

	n = fib(n)

	println(n)
}
