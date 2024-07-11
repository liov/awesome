package main

import (
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -LD:/sdk/gerbv -lgerbv-1
#cgo CFLAGS: -ID:/code/gerbv/src
#cgo pkg-config: glib-2.0 gtk+-2.0
// 安装msys2环境,安装相关库
// 删除 msys64\mingw64\lib\pkgconfig目录下gkd-win32-2.0pc Libs 行的-Wl,

#include <gerbv.h>

void set_project(gerbv_project_t* project) {

	project->file[1]->transform.translateY = 0.02;
	project->file[1]->transform.translateX = 0.02;


	GdkColor greenishColor = {0, 10000, 65000, 10000};
	project->file[0]->color = greenishColor;
}
*/
import "C"

type GerberProject C.gerbv_project_t

func CreateProject() *GerberProject {
	project := C.gerbv_create_project()
	return (*GerberProject)(unsafe.Pointer(project))
}

// OpenLayer opens a Gerber file and returns a GerberFile object.
func OpenLayer(project *GerberProject, filename string) {
	cFilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cFilename))

	C.gerbv_open_layer_from_filename((*C.gerbv_project_t)(unsafe.Pointer(project)), cFilename)
}

func main() {
	project := CreateProject()
	// 打开 Gerber 文件
	OpenLayer(project, "D:\\work\\Gerber\\m1")
	OpenLayer(project, "D:\\work\\Gerber\\m2")

	fileCount := project.max_files
	if fileCount < 2 {
		fmt.Println("There was an error parsing the files.")
		return
	}
	cproject := (*C.gerbv_project_t)(unsafe.Pointer(project))
	C.set_project(cproject)
	// 渲染 Gerber 文件为 PNG
	output := C.CString("output.png")
	defer C.free(unsafe.Pointer(output))
	C.gerbv_export_png_file_from_project_autoscaled(cproject, 640, 480, output)
	C.gerbv_destroy_project(cproject)
	fmt.Println("Gerber file has been successfully rendered to output.png")
}
