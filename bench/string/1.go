package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin sit amet euismod arcu. Duis ligula ex, feugiat mattis porta quis, maximus vitae ex. Praesent dapibus neque arcu, a consectetur sapien malesuada sit amet. Proin tempor ligula sed sem maximus, id."
	var delimiter = " "

	fmt.Println(strings.Split(s, delimiter))
}
