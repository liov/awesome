package main

import (
	"fmt"
	url2 "net/url"
)

func main() {
	url, _ := url2.Parse("c:/log/log.txt")
	fmt.Println(url.Path)
	url, _ = url2.Parse("file://c:/log/log.txt")
	fmt.Println(url.Path)
}
