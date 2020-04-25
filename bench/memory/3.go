package main

import (
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var S = 1000000 // 1 MB
	var N = 1000
	var args = os.Args[1:]
	if len(args) >= 1 {
		S, _ = strconv.Atoi(args[0])
		if len(args) >= 2 {
			N, _ = strconv.Atoi(args[1])
		}
	}

	S = S / strconv.IntSize

	var start = time.Now()

	for i := 0; i < N; i++ {
		_ = make([]int, S)
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
