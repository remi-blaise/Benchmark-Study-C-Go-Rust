package main

import (
	"os"
	"strconv"
)

func main() {
	var M = 1000 // Rows
	var N = 1000 // Columns
	var args = os.Args[1:]
	if len(args) >= 1 {
		M, _ = strconv.Atoi(args[0])
		if len(args) == 2 {
			N, _ = strconv.Atoi(args[1])
		}
	}

	var mat = make([][]int, M)
	for i := range mat {
		mat[i] = make([]int, N)
	}
}
