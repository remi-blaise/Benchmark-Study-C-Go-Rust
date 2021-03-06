package main

import (
	"math/cmplx"
	"math/rand"
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 50000000 // Number of operations
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		comp := complex128(complex(rand.Float32(), rand.Float32()))
		_ = cmplx.Cos(comp)
		_ = cmplx.Sin(comp)
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
