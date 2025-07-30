package main

import (
	"github.com/hopeio/gox/log"
	debugi "github.com/hopeio/gox/runtime/debug"
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
)

// pacman -S mingw-w64-ucrt-x86_64-graphviz
// go tool pprof -http 127.0.0.1:8080 memprofile.out
func main() {
	debug.SetGCPercent(-1)
	runtime.MemProfileRate = 1
	f, err := os.Create("cpuprofile.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()
	//debug.SetGCPercent(5)
	do()
	mf, err := os.Create("memprofile.out")
	if err != nil {
		log.Fatal(err)
	}
	defer mf.Close()
	if err := pprof.WriteHeapProfile(mf); err != nil {
		log.Fatal(err)
	}
}

func do() {
	f, _ := os.Open(`D:\xxx.jpg`)
	debugi.PrintMemoryUsage(1)
	img, _ := jpeg.Decode(f)
	debugi.PrintMemoryUsage(2)
	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	debugi.PrintMemoryUsage(3)
	draw.Draw(result, bounds, img, image.Point{}, draw.Src)
	debugi.PrintMemoryUsage(4)
	outFile, _ := os.Create(`xxx-copy.jpg`)
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 99})
	debugi.PrintMemoryUsage(5)
}
