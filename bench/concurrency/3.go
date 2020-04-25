package main

import (
	"os"
	"strconv"
	"syscall"
	"fmt"
	"time"
)

const (
	CHILD_ARG = "child"
)

func main() {
	var n = 1500 // Number of processes
	var args = os.Args[1:]
	if len(args) >= 1 {
		if args[0] == CHILD_ARG {
			// I'm a child
			return
		}
		n, _ = strconv.Atoi(args[0])
	}

	var start = time.Now()

	for i := 0; i < n; i++ {
		syscall.ForkExec(os.Args[0], []string{os.Args[0], CHILD_ARG}, &syscall.ProcAttr{Files: []uintptr{0, 1, 2}})
	}

	fmt.Println(time.Now().Sub(start).Nanoseconds())
}
