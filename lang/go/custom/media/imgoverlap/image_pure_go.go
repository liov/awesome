package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"
	"time"
)

// 恍然大悟，图片其实就是一维数组
func calculateMSE(y int, img1, img2 []uint8, minOverlap, maxOverlap int) int {
	n := len(img1)
	minMean := math.MaxInt
	var overlap int
	for o := minOverlap; o <= maxOverlap; o++ {
		var sum int
		m := o * y
		subimg1 := img1[n-m:]
		subimg2 := img2[:m]
		for i := range m {
			v := int(subimg1[i]) - int(subimg2[i])
			sum += v * v
		}
		mse := sum / m
		if mse < minMean {
			minMean = mse
			overlap = o
		}
	}

	return overlap
}

func main() {
	now := time.Now()
	file1, err := os.Open(`D:\work\1.jpg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	file2, err := os.Open(`D:\work\1.jpg`)
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

func ImageData(img1, img2 image.Image, maXOverlap int) ([]uint8, []uint8) {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	data1 := make([]uint8, 0, maXOverlap*bounds1.Dy())
	// 遍历原始图像的每个像素并转换为灰度值
	for x := bounds1.Max.X - maXOverlap; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data1 = append(data1, gray)
		}
	}
	data2 := make([]uint8, 0, maXOverlap*bounds2.Dy())
	for x := bounds2.Min.X; x <= maXOverlap; x++ {
		for y := bounds2.Min.Y; y < bounds2.Max.Y; y++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := uint8(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data2 = append(data2, gray)
		}
	}
	return data1, data2
}
