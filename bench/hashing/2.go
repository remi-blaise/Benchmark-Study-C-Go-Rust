package main

import (
	"crypto/sha1"
	"io/ioutil"
)

func main() {
	var dat, err = ioutil.ReadFile("asset/lorem100mb")
	if err != nil {
		panic(err)
	}

	var h = sha1.New()
	h.Write(dat)
	h.Sum(nil)
}
