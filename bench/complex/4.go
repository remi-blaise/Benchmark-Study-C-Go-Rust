package main

import (
	"math/cmplx"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 15000000 // Number of operations
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	for i := 0; i < n; i++ {
		comp := complex128(complex(rand.Float32(), rand.Float32()))
		_ = cmplx.Phase(comp)
	}
}
