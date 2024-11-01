package main

import (
	"fmt"
	"image"
)

func main() {
	var img *image.RGBA
	var img2 image.RGBA
	ptr(img)
	ptr(&img2)
	fmt.Println(img, img2)
	ptr2(img)
	ptr3(&img2)
	fmt.Println(img, img2)
	ptr3(img)
	fmt.Println(img, img2)
}

func ptr(img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, 10, 10))
}

func ptr2(img *image.RGBA) {
	img = image.NewRGBA(image.Rect(0, 0, 10, 10))
	*img = *img
}

func ptr3(img *image.RGBA) {
	*img = *image.NewRGBA(image.Rect(0, 0, 10, 10))
}
