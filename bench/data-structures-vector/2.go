package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var S = 1000000
	var n = 1000000
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

	for i := 0; i < n; i++ {
		fmt.Println(arr[rand.Int()%S])
	}
}
