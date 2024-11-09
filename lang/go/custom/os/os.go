package main

import (
	"github.com/hopeio/utils/log"
	"os"
)

func main() {
	log.Println(os.Rename(`D:\chls.pem`, `E:\chls.pem`))
}
