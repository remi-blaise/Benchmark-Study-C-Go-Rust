package main

import (
	"os"
	"strconv"
)

func main() {
	var S = 1_000_000_000 // 1 GB
	var args = os.Args[1:]
	if len(args) >= 1 {
		S, _ = strconv.Atoi(args[0])
	}

	S = S / strconv.IntSize

	var p = make([]int, S)
	for i := range p {
		p[i] = 42
	}
}
