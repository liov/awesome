package main

import (
	debugi "github.com/hopeio/gox/runtime/debug"
	"image"
	"image/jpeg"
	"os"
	"strconv"
)

func main() {
	var img image.RGBA
	ptr(&img)
	debugi.PrintMemoryUsage(len(img.Pix))
	ptr(&img)
	debugi.PrintMemoryUsage(len(img.Pix))
	ptr(&img)
	debugi.PrintMemoryUsage(len(img.Pix))
	for i := range 3 {
		outFile, _ := os.Create("D:/" + strconv.Itoa(i) + ".jpg")
		jpeg.Encode(outFile, &img, &jpeg.Options{Quality: 100})
		outFile.Close()
		debugi.PrintMemoryUsage(1)
	}
}

func ptr(img *image.RGBA) {
	if len(img.Pix) == 0 {
		*img = *image.NewRGBA(image.Rect(0, 0, 12990, 35688))
	}
}
