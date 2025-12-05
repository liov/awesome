package main

import (
	"image"
	"image/draw"
	"image/jpeg"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"

	"github.com/hopeio/gox/log"
	runtimex "github.com/hopeio/gox/runtime"
)

// pacman -S mingw-w64-ucrt-x86_64-graphviz
// go tool pprof -http=:8080 memprofile.out
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
	runtimex.PrintMemoryUsage(1)
	img, _ := jpeg.Decode(f)
	runtimex.PrintMemoryUsage(2)
	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	runtimex.PrintMemoryUsage(3)
	draw.Draw(result, bounds, img, image.Point{}, draw.Src)
	runtimex.PrintMemoryUsage(4)
	outFile, _ := os.Create(`xxx-copy.jpg`)
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 99})
	runtimex.PrintMemoryUsage(5)
}
