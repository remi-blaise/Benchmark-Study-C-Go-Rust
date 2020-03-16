package main

import (
	"os"
	"strconv"
)

// Recursion function
func rec(n uint64) uint64 {
	if n <= 0 {
		return n
	}

	return rec(n - 1)
}

func main() {
	var n uint64 = 1000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.ParseUint(args[1], 10, 64)
	}

	n = rec(n)

	println(n)
}
