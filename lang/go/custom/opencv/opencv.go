package main

/*
#cgo CXXFLAGS: --std=c++11
#cgo pkg-config: opencv4

#include "wrapper.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	imagePath := C.CString(`D:\work\z\0.96\21.bmp`)
	defer C.free(unsafe.Pointer(imagePath))

	width := C.load_image_width(imagePath)
	if width == -1 {
		fmt.Println("Failed to load image")
	} else {
		fmt.Printf("Image width: %d\n", width)
	}
}
