package main

import (
	"os"
	"sort"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var n = 100000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var arr = make([]int, n)

	// Already sorted array
	for i := 0; i < n; i++ {
		arr[i] = i
	}

	var start = time.Now()

	sort.Ints(arr)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
