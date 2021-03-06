package main

import (
	"os"
	"strconv"
	"sync"
	"fmt"
	"time"
)

var x = 0
var mux sync.Mutex

func subroutine() {
	mux.Lock()
	x++
	mux.Unlock()
}

func main() {
	var n = 2000000 // Number of goroutines
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		go subroutine()
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
