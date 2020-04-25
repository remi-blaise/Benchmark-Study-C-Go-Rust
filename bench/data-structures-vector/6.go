package main

import (
	"math/rand"
	"os"
	"sort"
	"strconv"
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 100000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var arr = make([]int, n)

	// Random array
	for i := 0; i < n; i++ {
		arr[i] = rand.Int()
	}

	var start = time.Now()

	sort.Ints(arr)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
