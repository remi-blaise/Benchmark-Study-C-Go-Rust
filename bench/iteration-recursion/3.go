package main

import (
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var n = 40
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	var x, y = 0, 1
	for i := 0; i < n; i++ {
		z := x + y
		x = y
		y = z
	}

	println(x)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
