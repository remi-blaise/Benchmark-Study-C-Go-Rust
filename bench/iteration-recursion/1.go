package main

import (
	"os"
	"strconv"
)

func main() {
	var n uint64 = 1000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.ParseUint(args[1], 10, 64)
	}

	for n > 0 {
		n--
	}

	println(n)
}
