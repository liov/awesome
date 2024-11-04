package main

import (
	"github.com/hopeio/utils/log"
	"github.com/hopeio/utils/os/shm"
	"github.com/hopeio/utils/runtime/debug"
)

func main() {
	mem, err := shm.New("test", 1024*1024*1024)
	if err != nil {
		panic(err)
	}
	var data []byte
	mem.ReadMemory(0, 1024*1024*1024, &data)
	log.Println(len(data))
	debug.PrintMemoryUsage(1)

}
