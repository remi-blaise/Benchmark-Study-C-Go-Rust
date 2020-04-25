package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
	"strconv"
	"fmt"
	"time"
)

// Data is a custom struct
type Data struct {
	A int
	B float64
	C string
}

func main() {
	var n = 100000
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		var dataObj = Data{A: 13273828327, B: 382283.537749, C: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet luctus leo sed imperdiet. Morbi ut dolor eu arcu pretium bibendum. Donec eleifend arcu sit amet sodales ultrices. Nam quis diam vel mi hendrerit egestas quis in velit. Aliquam non vulputate magna. Cras et magna bibendum, facilisis magna et, rhoncus."}

		var buf bytes.Buffer
		var enc = gob.NewEncoder(&buf)
		var dec = gob.NewDecoder(&buf)
		// Encode
		var err = enc.Encode(dataObj)
		if err != nil {
			log.Fatal("encode error:", err)
		}
		// Decode
		var dataObj2 Data
		err = dec.Decode(&dataObj2)
		if err != nil {
			log.Fatal("decode error:", err)
		}
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
