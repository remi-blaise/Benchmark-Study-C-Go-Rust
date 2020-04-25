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

	var S = 100000000
	var n = 100
	var args = os.Args[1:]
	if len(args) >= 1 {
		S, _ = strconv.Atoi(args[0])
		if len(args) >= 2 {
			n, _ = strconv.Atoi(args[1])
		}
	}

	var arr = make([]int, S)

	for i := 0; i < S; i++ {
		arr[i] = 42
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		var j = rand.Int() % S
		arr = append(arr[:j], arr[j+1:]...)
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
