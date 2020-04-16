package main

import (
	"fmt"
	"time"
)

func main() {
	var start = time.Now()
	var end = time.Now()
	var counter = 0
	for end.Sub(start) < 1000000000 {
		end = time.Now()
		counter++
	}
	fmt.Println(counter)
}
