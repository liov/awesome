package main

import (
	"fmt"
	"go/build"
	_ "golang.org/x/exp/shiny/driver"
	"log"
	"os"
	"path/filepath"
	"strings"
	_ "test/lang/cgo/disable/a"
)

func main() {

	//unicorn.Unicorn()
	//a.Hello()

	ctxt := new(build.Context)
	*ctxt = build.Default
	ctxt.CgoEnabled = false

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	aDir := filepath.Join(wd, "lang/cgo/disable/a")
	p, err := ctxt.ImportDir(aDir, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strings.Join(p.Imports, "\n"))
}
