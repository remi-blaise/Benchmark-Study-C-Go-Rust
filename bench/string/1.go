package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var dat, err = ioutil.ReadFile("asset/lorem100mb")
	if err != nil {
		panic(err)
	}

	var result = strings.Split(string(dat), " ")
	fmt.Println(len(result))
}
