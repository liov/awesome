package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("开始")
	for _, z := range []string{"0.96", "1.72", "2.8", "3.52", "4.64"} {
		Sharpness(fmt.Sprintf(`D:\work\z\%s\21.bmp`, z))
	}
}

func Sharpness(file string) {

	img := gocv.IMRead(file, gocv.IMReadGrayScale)
	defer img.Close()

	laplacian := gocv.NewMat()
	defer laplacian.Close()
	// 计算拉普拉斯算子的方差
	gocv.Laplacian(img, &laplacian, gocv.MatTypeCV64F, 1, 1, 0, gocv.BorderDefault)
	// 计算方差
	mean, stddev := gocv.NewMat(), gocv.NewMat()
	defer mean.Close()
	defer stddev.Close()
	gocv.MeanStdDev(laplacian, &mean, &stddev)
	fmt.Print(stddev.GetDoubleAt(0, 0), " ")

	// 计算 Sobel 梯度
	sobelX := gocv.NewMat()
	defer sobelX.Close()
	sobelY := gocv.NewMat()
	defer sobelY.Close()

	gocv.Sobel(img, &sobelX, gocv.MatTypeCV64F, 1, 0, 3, 1, 0, gocv.BorderDefault)
	gocv.Sobel(img, &sobelY, gocv.MatTypeCV64F, 0, 1, 3, 1, 0, gocv.BorderDefault)

	// 计算梯度幅值
	magnitude := gocv.NewMat()
	defer magnitude.Close()
	gocv.Magnitude(sobelX, sobelY, &magnitude)

	// 计算方差
	mean = gocv.NewMat()
	defer mean.Close()
	stddev = gocv.NewMat()
	defer stddev.Close()
	gocv.MeanStdDev(magnitude, &mean, &stddev)

	fmt.Println(stddev.GetDoubleAt(0, 0))
}
