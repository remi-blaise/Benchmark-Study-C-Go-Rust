package main

import (
	"bufio"
	"crypto/sha512"
	"os"
)

func main() {
	var f, err = os.Open("asset/1000000passwords")
	if err != nil {
		panic(err)
	}
	var scanner = bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		var h = sha512.New()
		h.Write(scanner.Text())
		h.Sum(nil)
	}
}
