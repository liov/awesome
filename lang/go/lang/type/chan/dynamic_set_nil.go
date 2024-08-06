package main

import "fmt"

type Engine struct {
	taskChanProducer chan int
	readyTaskHeap    []int
	waitTaskCount    int
}

func main() {
	e := &Engine{
		taskChanProducer: make(chan int, 1),
	}
	taskChan := e.taskChanProducer
	// go 无法动态的设置case,但是可以动态的把channel置为nil
	if len(e.readyTaskHeap) >= int(e.waitTaskCount) {
		taskChan = nil
	} else {
		taskChan = e.taskChanProducer
	}

	for {
		select {
		// 为nil 时，会阻塞
		case task := <-taskChan:
			fmt.Println(task)
		}
	}
}
