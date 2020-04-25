package main

import (
	"math/rand"
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var M = 1000 // Rows
	var N = 1000 // Columns
	var args = os.Args[1:]
	if len(args) >= 1 {
		M, _ = strconv.Atoi(args[0])
		if len(args) == 2 {
			N, _ = strconv.Atoi(args[1])
		}
	}

	var a = make([][]int, M) // Matrix A: M*N
	for i := range a {
		a[i] = make([]int, N)
		for j := range a[i] {
			a[i][j] = rand.Int()
		}
	}

	var b = make([][]int, N) // Matrix B: N*M
	for i := range b {
		b[i] = make([]int, M)
		for j := range b[i] {
			b[i][j] = rand.Int()
		}
	}

	var start = time.Now()

	// Multiply matrix
	var prod = make([][]int, M) // Product: M*M
	for i := 0; i < M; i++ {
		prod[i] = make([]int, M)
		for k := 0; k < N; k++ {
			for j := 0; j < M; j++ {
				prod[i][j] += a[i][k] * b[k][j]
			}
		}
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
