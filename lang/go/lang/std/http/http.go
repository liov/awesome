package main

import (
	"log"
	"mime"
	"net/http"
)

func main() {
	log.Println(mime.ParseMediaType("application/octet-stream"))
	log.Println(mime.ParseMediaType("form-data; name=\"file\"; filename=\"example.txt\""))
	log.Println(http.ParseCookie("form-data; name=\"file\"; filename=\"example.txt\""))
}
