package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 50000000 // Number of operations
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	for i := 0; i < n; i++ {
		complex1 := complex(rand.Float32(), rand.Float32())
		complex2 := complex(rand.Float32(), rand.Float32())
		_ = complex1 * complex2
		_ = complex1 / complex2
	}
}
