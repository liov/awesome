package main

import (
	"fmt"
	"github.com/ctessum/geom"
	ggeom2 "github.com/go-spatial/geom"
	ggeom "github.com/twpayne/go-geom"
)

func main() {
	point := geom.Point{}
	point2 := ggeom.Point{}
	point3 := ggeom2.Point{}
	fmt.Println(point, point2, point3)
}
