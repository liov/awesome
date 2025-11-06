package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"
)

// binarizeAndTransparent 将图片二值化并将白色部分透明化
func binarizeAndTransparent(img image.Image, threshold uint8) *image.RGBA {
	bounds := img.Bounds()
	// 创建一个新的RGBA图像
	newImg := image.NewRGBA(bounds)

	// 遍历每个像素
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// 获取原始像素颜色
			oldColor := img.At(x, y)
			r, g, b, _ := oldColor.RGBA()

			// 转换为8位值
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)

			// 计算灰度值（使用标准的亮度公式）
			gray := uint8(0.299*float64(r8) + 0.587*float64(g8) + 0.114*float64(b8))

			// 应用阈值进行二值化
			var newColor color.RGBA
			if gray > threshold {
				// 白色部分设为透明
				newColor = color.RGBA{0, 0, 0, 0}
			} else {
				// 黑色部分保持黑色，完全不透明
				newColor = color.RGBA{0, 0, 0, 255}
			}

			// 设置新像素颜色
			newImg.Set(x, y, newColor)
		}
	}

	return newImg
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供输入的JPG文件路径")
		fmt.Println("用法: go run binarize_transparent.go <input.jpg> [output.png] [threshold]")
		fmt.Println("示例: go run binarize_transparent.go input.jpg output.png 128")
		return
	}

	inputPath := os.Args[1]
	outputPath := "output.png"
	threshold := uint8(128)

	if len(os.Args) >= 3 {
		outputPath = os.Args[2]
	}

	if len(os.Args) >= 4 {
		var t int
		fmt.Sscanf(os.Args[3], "%d", &t)
		if t >= 0 && t <= 255 {
			threshold = uint8(t)
		}
	}

	// 打开输入文件
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("打开文件失败: %v\n", err)
		return
	}
	defer file.Close()

	// 解码JPG图像
	img, err := jpeg.Decode(file)
	if err != nil {
		fmt.Printf("解码JPG图像失败: %v\n", err)
		return
	}

	// 处理图像
	fmt.Println("正在处理图像...")
	resultImg := binarizeAndTransparent(img, threshold)

	// 创建输出文件
	outFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("创建输出文件失败: %v\n", err)
		return
	}
	defer outFile.Close()

	// 编码为PNG并保存
	err = png.Encode(outFile, resultImg)
	if err != nil {
		fmt.Printf("编码PNG图像失败: %v\n", err)
		return
	}

	fmt.Printf("处理完成！已保存到: %s\n", outputPath)
	fmt.Printf("使用的阈值: %d\n", threshold)
}