package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
)

func main() {
	fmt.Println("开始")

	SearchCircle(`D:\11.bmp`, image.Rect(0, 0, 5120, 5120))

}

func SearchCircle(file string, rect image.Rectangle) {
	rimg := gocv.IMRead(file, gocv.IMReadColor)
	defer rimg.Close()
	gimg := gocv.NewMat()
	defer gimg.Close()
	gocv.CvtColor(rimg, &gimg, gocv.ColorBGRToGray)

	// 定义高斯核的大小和标准差
	blurred := gocv.NewMat()
	defer blurred.Close()
	img := gimg.Region(rect)
	defer img.Close()
	gocv.GaussianBlur(img, &blurred, image.Pt(9, 9), 2.0, 2.0, gocv.BorderDefault)
	edges := gocv.NewMat()
	defer edges.Close()
	gocv.Canny(blurred, &edges, 50, 150)
	circles := gocv.NewMat()
	defer circles.Close()
	gocv.HoughCirclesWithParams(edges, &circles, gocv.HoughGradient, 1.2, 200, 200, 80, 20, 100)
	if !circles.Empty() {
		for i := 0; i < circles.Cols(); i++ {
			v := circles.GetVecfAt(0, i)
			x := int(v[0])
			y := int(v[1])
			r := int(v[2])
			// 检查圆是否完整，即圆的边缘不会超出图像边界
			if (x-r) > 0 && (x+r) < gimg.Cols() && (y-r) > 0 && (y+r) < gimg.Rows() {
				// 获取圆形边缘区域
				edgeMask := gocv.NewMatWithSize(gimg.Rows(), gimg.Cols(), gimg.Type())
				gocv.Circle(&edgeMask, image.Pt(x, y), r, color.RGBA{R: 255, G: 255, B: 255, A: 1}, 2)

				// 计算边缘区域的平均亮度
				meanEdge := gimg.MeanWithMask(edgeMask).Val1

				// 计算圆外区域的平均亮度
				outerMask := gocv.NewMatWithSize(gimg.Rows(), gimg.Cols(), gimg.Type())
				gocv.Circle(&outerMask, image.Pt(x, y), r+10, color.RGBA{R: 255, G: 255, B: 255, A: 1}, -1)
				gocv.Subtract(outerMask, edgeMask, &outerMask)
				meanOutside := gimg.MeanWithMask(outerMask).Val1

				// 保留边缘和外部对比度高的圆
				if math.Abs(meanEdge-meanOutside) > 0 { // 根据需要调整此阈值
					// 绘制外圆
					gocv.Circle(&rimg, image.Pt(x, y), r, color.RGBA{R: 255, G: 0, B: 0, A: 0}, 2)
					// 绘制圆心
					gocv.Circle(&rimg, image.Pt(x, y), 2, color.RGBA{R: 0, G: 0, B: 255, A: 0}, 3)
				}
			}
		}
		// 创建一个窗口
		window := gocv.NewWindow("Image")

		// 在窗口中显示图片
		window.IMShow(rimg)

		// 等待按键按下，然后关闭窗口
		window.WaitKey(0)

		// 关闭窗口
		window.Close()
	}
}
