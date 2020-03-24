package main

import (
	"os"
	"strconv"
)

func main() {
	var S = 1_000_000 // 1 MB
	var N = 1000
	var args = os.Args[1:]
	if len(args) >= 1 {
		S, _ = strconv.Atoi(args[0])
		if len(args) >= 2 {
			N, _ = strconv.Atoi(args[1])
		}
	}

	S = S / strconv.IntSize

	for i := 0; i < N; i++ {
		var p = make([]int, S)
	}
}
