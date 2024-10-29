package main

import (
	"github.com/hopeio/utils/log"
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
	printMemoryUsage(1)
	img, _ := jpeg.Decode(f)
	printMemoryUsage(2)
	bounds := img.Bounds()
	result := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	printMemoryUsage(3)
	draw.Draw(result, bounds, img, image.Point{}, draw.Src)
	printMemoryUsage(4)
	outFile, _ := os.Create(`xxx-copy.jpg`)
	jpeg.Encode(outFile, img, &jpeg.Options{Quality: 99})
	printMemoryUsage(5)
}

func printMemoryUsage(flag any) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	log.Printf("%v Alloc = %v MiB,TotalAlloc = %v MiB, Sys = %v MiB,Mallocs = %v,HeapAlloc = %v,StackInuse = %v,StackSys = %v,NumGC = %v", flag, bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), bToMb(m.Mallocs), bToMb(m.HeapAlloc), bToMb(m.StackInuse), bToMb(m.StackSys), m.NumGC)
}

func printStack() {
	// 创建一个 1MB 的缓冲区来存储堆栈信息
	buf := make([]byte, 1<<20) // 1MB 缓冲区
	// 获取当前 Goroutine 的堆栈信息
	stackLen := runtime.Stack(buf, false)
	log.Printf("当前堆栈信息:\n%s", buf[:stackLen])
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
