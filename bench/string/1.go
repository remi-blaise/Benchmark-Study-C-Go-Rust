package main

import (
	"fmt"
	"time"
	"io/ioutil"
	"strings"
)

func main() {
	var dat, err = ioutil.ReadFile("asset/lorem100mb")
	if err != nil {
		panic(err)
	}

	var start = time.Now()

	var result = strings.Split(string(dat), " ")
	fmt.Println(len(result))

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
