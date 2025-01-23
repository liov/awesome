package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"
	"test/custom/asm/asm"
	"time"
)

func main() {
	// 示例用法
	// img1 和 img2 是从文件加载的图像
	// minOverlap 和 maxOverlap 是重叠范围
	now := time.Now()
	file1, _ := os.Open(`D:\work\1.jpg`)
	img1, _ := jpeg.Decode(file1)
	file2, _ := os.Open(`D:\work\1.jpg`)
	img2, _ := jpeg.Decode(file2)
	minOverlap := 1200
	maXOverlap := 1300
	data1, data2 := ImageData(img1, img2, maXOverlap)
	file1.Close()
	file2.Close()
	bounds1 := img1.Bounds()
	overlap := calculateMSE(bounds1.Dy(), data1, data2, minOverlap, maXOverlap)
	fmt.Println("Best overlap:", overlap, time.Since(now))
}

func calculateMSE(y int, img1, img2 []int16, minOverlap, maxOverlap int) int {
	n := len(img1)
	minMean := math.MaxInt
	var overlap int
	for o := minOverlap; o <= maxOverlap; o++ {
		var sum int
		m := o * y
		z := m - 16
		subimg1 := img1[n-m:]
		subimg2 := img2[:m]
		var i = 0
		for ; i < z; i += 16 {
			sum += int(asm.Avx2SsdInt16(subimg1[i:i+16], subimg2[i:i+16]))
		}
		for i := range m - z {
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

func ImageData(img1, img2 image.Image, maXOverlap int) ([]int16, []int16) {
	bounds1, bounds2 := img1.Bounds(), img2.Bounds()
	data1 := make([]int16, 0, maXOverlap*bounds1.Dy())
	// 遍历原始图像的每个像素并转换为灰度值
	for x := bounds1.Max.X - maXOverlap; x < bounds1.Max.X; x++ {
		for y := bounds1.Min.Y; y < bounds1.Max.Y; y++ {
			r, g, b, _ := img1.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := int16(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data1 = append(data1, gray)
		}
	}
	data2 := make([]int16, 0, maXOverlap*bounds2.Dy())
	for x := bounds2.Min.X; x <= maXOverlap; x++ {
		for y := bounds2.Min.Y; y < bounds2.Max.Y; y++ {
			r, g, b, _ := img2.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := int16(0.299*float64(r>>8) + 0.587*float64(g>>8) + 0.114*float64(b>>8))
			data2 = append(data2, gray)
		}
	}
	return data1, data2
}
