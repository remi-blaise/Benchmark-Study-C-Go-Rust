package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/google/btree"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var n = 1000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var tr = btree.New(2)

	for i := 0; i < n; i++ {
		tr.ReplaceOrInsert(btree.Int(rand.Int()))
	}
}
