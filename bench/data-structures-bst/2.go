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

	var S = 1000000
	var n = 1000000
	var args = os.Args[1:]
	if len(args) >= 1 {
		S, _ = strconv.Atoi(args[0])
		if len(args) >= 2 {
			n, _ = strconv.Atoi(args[1])
		}
	}

	var tr = btree.New(2)

	for i := 0; i < S; i++ {
		tr.ReplaceOrInsert(btree.Int(rand.Int()))
	}

	for i := 0; i < n; i++ {
		tr.Delete(btree.Int(rand.Int()))
	}
}
