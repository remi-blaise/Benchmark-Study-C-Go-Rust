package main

import (
	"os"
	"strconv"
)

func main() {
	var n = 100
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var x, y = 0, 1
	for i := 0; i < n; i++ {
		z := x + y
		x = y
		y = z
	}

	println(x)
}
