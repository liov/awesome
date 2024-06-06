package main

func main() {
	c1()
}

func c1() {
	var ch chan int
	// 这种操作会阻塞
	<-ch
}
func c2() {
	var ch chan int
	// 这种操作会阻塞
	ch <- 1
}
