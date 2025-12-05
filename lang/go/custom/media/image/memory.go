package main

import (
	"image"
	"image/jpeg"
	"os"
	"strconv"

	runtimex "github.com/hopeio/gox/runtime"
)

func main() {
	var img image.RGBA
	ptr(&img)
	runtimex.PrintMemoryUsage(len(img.Pix))
	ptr(&img)
	runtimex.PrintMemoryUsage(len(img.Pix))
	ptr(&img)
	runtimex.PrintMemoryUsage(len(img.Pix))
	for i := range 3 {
		outFile, _ := os.Create("D:/" + strconv.Itoa(i) + ".jpg")
		jpeg.Encode(outFile, &img, &jpeg.Options{Quality: 100})
		outFile.Close()
		runtimex.PrintMemoryUsage(1)
	}
}

func ptr(img *image.RGBA) {
	if len(img.Pix) == 0 {
		*img = *image.NewRGBA(image.Rect(0, 0, 12990, 35688))
	}
}
