package main

import (
	"os"
	"strconv"
)

var x = 0

func subroutine() {
	x++
}

func main() {
	var n = 10000 // Number of goroutines
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	for i := 0; i < n; i++ {
		go subroutine()
	}
}
