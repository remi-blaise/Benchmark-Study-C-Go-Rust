package main

import (
	"crypto/sha1"
	"io/ioutil"
	"fmt"
	"time"
)

func main() {
	var dat, err = ioutil.ReadFile("asset/lorem100mb")
	if err != nil {
		panic(err)
	}

	var start = time.Now()

	var h = sha1.New()
	h.Write(dat)
	h.Sum(nil)

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
