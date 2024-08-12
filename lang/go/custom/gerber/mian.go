package main

import (
	"encoding/json"
	"github.com/hopeio/utils/encoding/gerber"
	"github.com/hopeio/utils/log"
	"os"

	"test/custom/gerber/svg"
)

func main() {
	log.Println("开始")
	p := svg.NewProcessor()
	processor := gerber.NewParser(p)
	file, err := os.OpenFile(`D:\work\Gerber\Gerber_TopLayer.GTL`, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = processor.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}
