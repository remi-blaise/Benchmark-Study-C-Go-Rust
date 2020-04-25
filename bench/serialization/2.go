package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
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

	for i := 0; i < n; i++ {
		var dataObj = Data{A: 13273828327, B: 382283.537749, C: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed laoreet luctus leo sed imperdiet. Morbi ut dolor eu arcu pretium bibendum. Donec eleifend arcu sit amet sodales ultrices. Nam quis diam vel mi hendrerit egestas quis in velit. Aliquam non vulputate magna. Cras et magna bibendum, facilisis magna et, rhoncus."}

		// Encode
		var data, err = json.Marshal(dataObj)
		if err != nil {
			log.Fatal("encode error:", err)
		}
		// Decode
		var dataObj2 Data
		err = json.Unmarshal(data, &dataObj2)
		if err != nil {
			log.Fatal("decode error:", err)
		}
	}
}
