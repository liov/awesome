package main

import (
	"fmt"
	imagei "github.com/hopeio/utils/media/image"
	"image/jpeg"
	"os"
)

func main() {
	// 示例用法
	// img1 和 img2 是从文件加载的图像
	// minOverlap 和 maxOverlap 是重叠范围
	f1, _ := os.Open(`D:\work\scan-panel-0826\0--light1.jpg`)
	img1, _ := jpeg.Decode(f1)
	f2, _ := os.Open(`D:\work\scan-panel-0826\1--light1.jpg`)
	img2, _ := jpeg.Decode(f2)
	overlap := imagei.CalculateOverlap(img1, img2, 1200, 1300)
	fmt.Println("Best overlap:", overlap)
}
