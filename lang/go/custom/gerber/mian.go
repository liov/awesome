package main

import (
	"github.com/hopeio/gox/encoding/gerber"
	"github.com/hopeio/gox/encoding/gerber/svg"
	"github.com/hopeio/gox/log"
	"os"
)

func main() {
	log.Println("开始")
	p := svg.NewProcessor()
	p.PanZoom = false
	processor := gerber.NewParser(p)
	file, err := os.OpenFile(`D:\work\Gerber\Gerber_TopLayer.GTL`, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	err = processor.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	svg, _ := os.Create(`./output.svg`)
	p.Write(svg)
	svg.Close()
}
