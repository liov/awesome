package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"image"
	"image/jpeg"
	"math"
	"os"
	"time"
)

func calculateMSE(y int, img1, img2 []float64, minOverlap, maxOverlap int) int {

	minMean := math.MaxFloat64
	var overlap int
	data := make([]float64, y*maxOverlap)
	for i := minOverlap; i <= maxOverlap; i++ {
		a := mat.NewDense(i, y, img1[(maxOverlap-i)*y:])
		b := mat.NewDense(i, y, img2[:i*y])
		diff := mat.NewDense(i, y, data[:i*y])
		diff.Sub(a, b)
		var sum float64
		diff.Apply(func(i, j int, v float64) float64 {
			sum += v * v
			return v
		}, diff)
		mse := sum / float64(i*y)
		if mse < minMean {
			minMean = mse
			overlap = i
		}
	}

	return overlap
}

func main() {
	now := time.Now()
	file1, err := os.Open(`D:\work\scan-panel-0826\0--light1.jpg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file2, err := os.Open(`D:\work\scan-panel-0826\1--light1.jpg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	img1, err := jpeg.Decode(file1)
	img2, err := jpeg.Decode(file2)

	minOverlap := 1200
	maXOverlap := 1300
	data1, data2 := ImageData(img1, img2, maXOverlap)
	file1.Close()
	file2.Close()
	bounds1 := img1.Bounds()
	fmt.Println(calculateMSE(bounds1.Dy(), data1, data2, minOverlap, maXOverlap), time.Since(now))
}

func ImageData(img1, img2 image.Image, maXOverlap int) ([]float64, []float64) {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	data1 := make([]float64, 0, maXOverlap*bounds1.Dy())
	// 遍历原始图像的每个像素并转换为灰度值
	for x := bounds1.Max.X - maXOverlap; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			data1 = append(data1, gray)
		}
	}
	data2 := make([]float64, 0, maXOverlap*bounds2.Dy())
	for x := bounds2.Min.X; x <= maXOverlap; x++ {
		for y := bounds2.Min.Y; y < bounds2.Max.Y; y++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := 0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8)
			data2 = append(data2, gray)
		}
	}
	return data1, data2
}
