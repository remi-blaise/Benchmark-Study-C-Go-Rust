package main

import (
	"hash/fnv"
	"os"
	"strconv"
)

func main() {
	var n = 10000 // Number of goroutines
	var args = os.Args[1:]
	if len(args) >= 1 {
		n, _ = strconv.Atoi(args[0])
	}

	var str = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nam id ante non enim sodales interdum. In imperdiet varius dui tincidunt pulvinar. Aenean enim lorem, posuere et nibh eu, scelerisque facilisis turpis. Cras efficitur aliquam pellentesque. Donec eleifend velit vel magna pellentesque, ut malesuada mauris tempor. Pellentesque quis augue dapibus, pulvinar eros et, posuere eros. Praesent suscipit lacinia lorem, quis feugiat elit pulvinar sed. Nullam purus est, porta in ante mattis, tempus pellentesque leo. Cras quam nisi, pellentesque et est sed, vestibulum faucibus felis. Class aptent taciti sociosqu ad litora torquent per conubia nostra, per inceptos himenaeos."

	for i := 0; i < n; i++ {
		var h = fnv.New64()
		h.Write([]byte(str))
		h.Sum64()
	}
}
