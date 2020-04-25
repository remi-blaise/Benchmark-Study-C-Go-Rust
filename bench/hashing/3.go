package main

import (
	"bufio"
	"crypto/sha512"
	"os"
	"fmt"
	"time"
)

func main() {
	var f, err = os.Open("asset/1000000passwords")
	if err != nil {
		panic(err)
	}
	var scanner = bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var start = time.Now()

	for scanner.Scan() {
		var h = sha512.New()
		h.Write([]byte(scanner.Text()))
		h.Sum(nil)
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
