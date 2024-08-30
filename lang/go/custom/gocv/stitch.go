package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"os"
	"path/filepath"
	"sync"
)

var (
	currentRow int
	currentCol int
	outputDir  string
)

func calculateOverlap(img1, img2 gocv.Mat, minOverlap, maxOverlap int) int {
	fmt.Println("calculating overlap ...")
	bestOverlap := 0
	minMSE := float64(1e9)

	for overlap := minOverlap; overlap <= maxOverlap; overlap++ {
		overlapRegion1 := img1.Region(image.Rect(img1.Cols()-overlap, 0, img1.Cols(), img1.Rows()))
		overlapRegion2 := img2.Region(image.Rect(0, 0, overlap, img2.Rows()))

		mse := calculateMSE(overlapRegion1, overlapRegion2)

		if mse < minMSE {
			minMSE = mse
			bestOverlap = overlap
		}
	}

	fmt.Println("overlap: ", bestOverlap)
	return bestOverlap
}

func calculateMSE(img1, img2 gocv.Mat) float64 {
	diff := gocv.NewMat()
	gocv.Subtract(img1, img2, &diff)
	squaredDiff := gocv.NewMat()
	gocv.Multiply(diff, diff, &squaredDiff)
	sum := gocv.Sum(squaredDiff).Val[0]
	return sum / float64(img1.Rows()*img1.Cols())
}

func stitchRow(imagesRow []gocv.Mat, minOverlap, maxOverlap int) (gocv.Mat, []int) {
	fmt.Println("stitching row - ", currentRow)
	stitchedRow := imagesRow[0]
	horizontalOverlaps := []int{}

	for i := 1; i < len(imagesRow); i++ {
		img1 := stitchedRow
		img2 := imagesRow[i]

		overlap := calculateOverlap(img1, img2, minOverlap, maxOverlap)
		horizontalOverlaps = append(horizontalOverlaps, overlap)

		newWidth := img1.Cols() + img2.Cols() - overlap
		height := img1.Rows()

		newStitchedRow := gocv.NewMatWithSize(height, newWidth, gocv.MatTypeCV8UC3)
		gocv.CopyTo(img1.Region(image.Pt(0, 0), image.Pt(img1.Cols()-overlap, img1.Rows())), &newStitchedRow.Region(image.Pt(0, 0), image.Pt(img1.Cols()-overlap, img1.Rows())))
		gocv.CopyTo(img2, &newStitchedRow.Region(image.Pt(img1.Cols()-overlap, 0), image.Pt(newWidth, img1.Rows())))

		stitchedRow = newStitchedRow
	}

	outPath := filepath.Join(outputDir, fmt.Sprintf("%d.png", currentRow))
	saveImage(stitchedRow, outPath)
	currentRow++
	return stitchedRow, horizontalOverlaps
}

func stitchColumn(stitchedRows []gocv.Mat, minOverlap, maxOverlap int) (gocv.Mat, []int) {
	fmt.Println("stitching col - ", currentCol)
	stitchedImage := stitchedRows[0]
	verticalOverlaps := []int{}

	for i := 1; i < len(stitchedRows); i++ {
		img1 := stitchedImage
		img2 := stitchedRows[i]

		img1, img2 = resizeImagesToSameWidth(img1, img2)

		overlap := calculateOverlap(img1.T(), img2.T(), minOverlap, maxOverlap)
		verticalOverlaps = append(verticalOverlaps, overlap)

		newHeight := img1.Rows() + img2.Rows() - overlap
		width := img1.Cols()

		newStitchedImage := gocv.NewMatWithSize(newHeight, width, gocv.MatTypeCV8UC3)
		gocv.CopyTo(img1.Region(image.Pt(0, 0), image.Pt(img1.Cols(), img1.Rows()-overlap)), &newStitchedImage.Region(image.Pt(0, 0), image.Pt(img1.Cols(), img1.Rows()-overlap)))
		gocv.CopyTo(img2, &newStitchedImage.Region(image.Pt(0, img1.Rows()-overlap), image.Pt(img1.Cols(), newHeight)))

		stitchedImage = newStitchedImage
	}

	return stitchedImage, verticalOverlaps
}

func saveOverlapInfo(horizontalOverlaps, verticalOverlaps []int, filePath string) {
	overlapInfo := map[string][]int{
		"horizontal_overlaps": horizontalOverlaps,
		"vertical_overlaps":   verticalOverlaps,
	}

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	err = encoder.Encode(overlapInfo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func saveImage(image gocv.Mat, path string) {
	img := toPILImage(image)
	img.Save(path, "png")
}

func toPILImage(mat gocv.Mat) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, mat.Cols(), mat.Rows()), color.Palette{color.White})
	for y := 0; y < mat.Rows(); y++ {
		for x := 0; x < mat.Cols(); x++ {
			val := mat.GetUCharAt(y, x)
			img.SetColorIndex(x, y, int(val))
		}
	}
	return img
}

func resizeImagesToSameWidth(img1, img2 gocv.Mat) (gocv.Mat, gocv.Mat) {
	height1, width1 := img1.Rows(), img1.Cols()
	height2, width2 := img2.Rows(), img2.Cols()

	if width1 < width2 {
		img1 = addBorder(img1, 0, 0, 0, width2-width1, color.RGBA{0, 0, 0, 255})
	} else if width1 > width2 {
		img2 = addBorder(img2, 0, 0, 0, width1-width2, color.RGBA{0, 0, 0, 255})
	}

	return img1, img2
}

func addBorder(img gocv.Mat, top, bottom, left, right int, color color.RGBA) gocv.Mat {
	return gocv.CopyMakeBorder(img, top, bottom, left, right, gocv.BorderConstant, color)
}

func main() {
	outputDir = "out/"
	os.MkdirAll(outputDir, os.ModePerm)

	imageFolder := "tmppic"
	imageFilenamesGrid := [][]string{
		{"0.bmp", "1.bmp", "2.bmp", "3.bmp", "4.bmp", "5.bmp", "6.bmp"},
	}

	imagesGrid := loadImagesGrid(imageFolder, imageFilenamesGrid)

	minOverlap := 2500
	maxOverlap := 2600

	stitchedImage, horizontalOverlaps, verticalOverlaps := stitchImagesGrid(imagesGrid, minOverlap, maxOverlap)

	saveImage(stitchedImage, filepath.Join(outputDir, "stitched_result.png"))
	saveOverlapInfo(horizontalOverlaps, verticalOverlaps, filepath.Join(outputDir, "overlap_info.json"))
}

func loadImagesGrid(imageFolder string, imageFilenamesGrid [][]string) [][]gocv.Mat {
	images := [][]gocv.Mat{}

	for _, rowFilenames := range imageFilenamesGrid {
		rowImages := []gocv.Mat{}
		for _, filename := range rowFilenames {
			img := gocv.IMRead(filepath.Join(imageFolder, filename), gocv.IMReadColor)
			if img.Empty() {
				fmt.Println("Error reading image:", filename)
				continue
			}
			rowImages = append(rowImages, img)
		}
		images = append(images, rowImages)
	}

	return images
}

func stitchImagesGrid(imagesGrid [][]gocv.Mat, minOverlap, maxOverlap int) (gocv.Mat, []int, []int) {
	var wg sync.WaitGroup
	stitchedRows := [][]gocv.Mat{}
	allHorizontalOverlaps := [][]int{}

	for _, row := range imagesGrid {
		wg.Add(1)
		go func(row []gocv.Mat) {
			defer wg.Done()
			stitchedRow, horizontalOverlaps := stitchRow(row, minOverlap, maxOverlap)
			stitchedRows = append(stitchedRows, []gocv.Mat{stitchedRow})
			allHorizontalOverlaps = append(allHorizontalOverlaps, horizontalOverlaps)
		}(row)
	}

	wg.Wait()

	var finalStitchedImage gocv.Mat
	var verticalOverlaps []int

	if len(stitchedRows) > 0 {
		finalStitchedImage, verticalOverlaps = stitchColumn(stitchedRows[0], minOverlap, maxOverlap)
		for i := 1; i < len(stitchedRows); i++ {
			rowImage, rowOverlaps := stitchColumn([]gocv.Mat{stitchedRows[i][0]}, minOverlap, maxOverlap)
			finalStitchedImage = stitchImagesHorizontally(finalStitchedImage, rowImage)
			verticalOverlaps = append(verticalOverlaps, rowOverlaps...)
		}
	}

	return finalStitchedImage, flatten(allHorizontalOverlaps), verticalOverlaps
}

func stitchImagesHorizontally(img1, img2 gocv.Mat) gocv.Mat {
	overlap := calculateOverlap(img1, img2, 0, 0)
	newWidth := img1.Cols() + img2.Cols() - overlap
	height := img1.Rows()

	newStitchedImage := gocv.NewMatWithSize(height, newWidth, gocv.MatTypeCV8UC3)
	gocv.CopyTo(img1.Region(image.Pt(0, 0), image.Pt(img1.Cols(), img1.Rows())), &newStitchedImage.Region(image.Pt(0, 0), image.Pt(img1.Cols(), img1.Rows())))
	gocv.CopyTo(img2, &newStitchedImage.Region(image.Pt(img1.Cols()-overlap, 0), image.Pt(newWidth, img1.Rows())))

	return newStitchedImage
}

func flatten(sliceOfSlices [][]int) []int {
	flatSlice := []int{}
	for _, slice := range sliceOfSlices {
		flatSlice = append(flatSlice, slice...)
	}
	return flatSlice
}
