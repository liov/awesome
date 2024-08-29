package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"image/jpeg"
	"math"
	"os"
)

func calculateMSE(img []float64, minOverlap int) int {
	img1 := img[:3319*4074]
	img2 := img[3000*4074:]
	minMean := math.MaxFloat64
	var overlap int
	for i := minOverlap; i < 3319; i++ {
		a := mat.NewDense(i, 4074, img1[(3319-i)*4074:])
		b := mat.NewDense(i, 4074, img2[:i*4074])
		// 计算矩阵 A 和 B 之间的差值矩阵
		diff := mat.NewDense(i, 4074, nil)
		diff.Sub(a, b)
		var sum float64
		diff.Apply(func(i, j int, v float64) float64 {
			sum += v * v
			return v
		}, diff)
		if sum < minMean {
			minMean = sum
			overlap = i
		} else {
			break
		}
		fmt.Println(sum)
	}
	return overlap
}

func main() {
	file, err := os.Open(`D:\collect\7ad4d9794f1a39a77224fb21db81e16c.jpeg`)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	bounds := img.Bounds()
	data := make([]float64, 0, bounds.Dx()*bounds.Dy())
	// 遍历原始图像的每个像素并转换为灰度值
	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// 使用加权平均公式计算灰度值
			gray := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			data = append(data, gray)
		}
	}
	fmt.Println(calculateMSE(data, 300))
}
