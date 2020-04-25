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

	var m = make(map[int]int)

	var n = 10000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		m[rand.Int()] = 0
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
