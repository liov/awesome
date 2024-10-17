package main

import (
	"fmt"
	"image"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// 打开源图片文件
	src, err := imaging.Open("D:\\xxx")
	if err != nil {
		log.Fatalf("打开图片失败：%v", err)
	}

	// 定义裁剪区域的坐标和大小 (x, y, width, height)
	rect := image.Rect(6070, 17107, 11222, 18779)

	// 裁剪图片
	croppedImg := imaging.Crop(src, rect)

	// 保存裁剪后的图片
	err = imaging.Save(croppedImg, "cropped.jpg")
	if err != nil {
		log.Fatalf("保存图片失败：%v", err)
	}

	fmt.Println("图片裁剪完成并保存为 cropped.jpg")
}
