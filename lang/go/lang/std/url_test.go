package main

import (
	"fmt"
	url2 "net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	url, _ := url2.Parse("c:/log/log.txt")
	fmt.Println(url.Path)
	url, _ = url2.Parse("file://c:/log/log.txt")
	fmt.Println(url.Path)
}

func TestUrl2(t *testing.T) {
	url, err := url2.Parse("https://a.b/index/?&a")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(url.Path)
}
