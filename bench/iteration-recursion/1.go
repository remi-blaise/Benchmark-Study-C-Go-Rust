package main

import (
	"os"
	"strconv"
	"fmt"
	"time"
)

func main() {
	var n uint64 = 1000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.ParseUint(args[0], 10, 64)
	}

	var start = time.Now()

	for n > 0 {
		n--
	}

	println(n)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
