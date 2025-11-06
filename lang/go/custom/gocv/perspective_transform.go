package main

import (
	"fmt"
	"image"

	imagex "github.com/hopeio/gox/media/image"
	"gocv.io/x/gocv"
)

func main() {
	centerX, centerY, length, width, angle := 8646, 17943, 5152, 1672, float64(180)
	points := imagex.RectRotateByCenter(centerX, centerY, length, width, angle)
	srcPoints := gocv.NewPointVectorFromPoints(points)
	dstPoints := gocv.NewPointVectorFromPoints([]image.Point{
		{X: 0, Y: 0},
		{X: length, Y: 0},
		{X: length, Y: width},
		{X: 0, Y: width},
	})
	// count perspective transform matrix
	transformMat := gocv.GetPerspectiveTransform(srcPoints, dstPoints)
	fmt.Println(transformMat.Type(), transformMat.Rows(), transformMat.Cols())
	src := gocv.NewMatWithSize(1, 1, gocv.MatTypeCV64FC2)
	src.SetDoubleAt(0, 0, 11028)
	src.SetDoubleAt(0, 1, 17375)
	dst := gocv.NewMatWithSize(1, 1, gocv.MatTypeCV64FC2)
	gocv.PerspectiveTransform(src, &dst, transformMat)
	src.Close()
	transformMat.Close()
	srcPoints.Close()
	dstPoints.Close()
	for i := 0; i < dst.Rows(); i++ {
		x := dst.GetDoubleAt(i, 0)
		y := dst.GetDoubleAt(i, 1)
		fmt.Printf("(%f, %f)\n", x, y)
	}
}
