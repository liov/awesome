package main

import (
	"gocv.io/x/gocv"
	"image"
	"log"
	"os"
	"strconv"
)

func main() {
	MergeTiff([][]int{{0, 1, 2}, {5, 4, 3}, {6, 7, 8}, {11, 10, 9}, {12, 13, 14}, {17, 16, 15}, {18, 19, 20}}, func(i int) ([]byte, error) {
		return os.ReadFile(`D:\work\` + strconv.Itoa(i) + ".tiff")
	},
		image.Rect(0,
			0, 5120, 5120),
		[]int{1223, 1224}, []int{37, 40, 36, 38, 38, 40},
		`D:\result.tiff`)
}
func MergeTiff(imgs [][]int, getImage func(int) ([]byte, error), bounds image.Rectangle,
	horizontalOverlaps,
	verticalOverlaps []int, dst string) error {
	var resultWidth, resultHeight int
	for i := range imgs[0] {
		resultWidth += bounds.Dx()
		if i < len(horizontalOverlaps) {
			resultWidth -= horizontalOverlaps[i]
		}
	}
	for i := range imgs {
		resultHeight += bounds.Dy()
		if i < len(verticalOverlaps) {
			resultHeight -= verticalOverlaps[i]
		}
	}
	data, err := getImage(0)
	if err != nil {
		return err
	}
	img0, err := gocv.IMDecode(data, gocv.IMReadGrayScale|gocv.IMReadAnyDepth)
	if err != nil {
		return err
	}
	log.Print(resultWidth, resultHeight)
	result := gocv.NewMatWithSize(resultHeight, resultWidth, img0.Type())

	var rbounds = bounds
	// 将 img1 复制到结果图片中
	for i, rimg := range imgs {
		for j, imgIdx := range rimg {
			if imgIdx != 0 {
				data, err = getImage(imgIdx)
				if err != nil {
					return err
				}
			}
			img, err := gocv.IMDecode(data, gocv.IMReadGrayScale|gocv.IMReadAnyDepth)
			if err != nil {
				return err
			}
			log.Print(imgIdx, rbounds)
			rect := result.Region(rbounds)
			img.CopyTo(&rect)
			if j < len(horizontalOverlaps) {
				rbounds.Min.X += bounds.Dx() - horizontalOverlaps[j]
				rbounds.Max.X = bounds.Dx() + rbounds.Min.X
			}
		}
		if i < len(verticalOverlaps) {
			rbounds.Min.Y += bounds.Dy() - verticalOverlaps[i]
			rbounds.Max.Y = bounds.Dy() + rbounds.Min.Y
			rbounds.Min.X = 0
			rbounds.Max.X = bounds.Dx()
		}
	}
	gocv.IMWrite(dst, result)
	return nil
}
