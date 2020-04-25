package main

import "fmt"
import "time"

func main() {
	var start = time.Now()

	fmt.Println("Hello, world!")

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
