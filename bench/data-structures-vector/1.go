package main

import (
	"os"
	"strconv"
)

func main() {
	var n = 100000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var arr = make([]int, 0, n)

	for i := 0; i < n; i++ {
		arr = append(arr, 42)
	}
}
