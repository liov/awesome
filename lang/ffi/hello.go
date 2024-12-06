package main

import "C"
import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

var (
	newplus          = syscall.MustLoadDLL("libnewplus.dll")
	currentTimestamp = newplus.MustFindProc("current_timestamp")
	plusone          = newplus.MustFindProc("plusone")
)

func run(count int) {
	start, _, _ := currentTimestamp.Call()

	var x int = 0
	for x < count {
		ux, _, _ := plusone.Call(uintptr(x))
		x = int(ux)
	}
	end, _, _ := currentTimestamp.Call()
	fmt.Println(end - start)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("First arg (0 - 2000000000) is required.")
		return
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil || count <= 0 || count > 2000000000 {
		fmt.Println("Must be a positive number not exceeding 2 billion.")
		return
	}

	// load
	start, _, _ := currentTimestamp.Call()
	plusone.Call(start)

	// start
	run(count)
}
