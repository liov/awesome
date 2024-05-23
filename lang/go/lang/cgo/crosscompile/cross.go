package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strings"
	"test/lang/cgo/crosscompile/mobile"

	_ "golang.org/x/exp/shiny/driver"
	_ "test/lang/cgo/crosscompile/a"
)

func main() {

	//unicorn.Unicorn()
	//a.Hello()
	mobile.Mobile()
	ctxt := new(build.Context)
	*ctxt = build.Default
	ctxt.CgoEnabled = false

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	aDir := filepath.Join(wd, "a")
	p, err := ctxt.ImportDir(aDir, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(p.Imports, "\n"))
}
