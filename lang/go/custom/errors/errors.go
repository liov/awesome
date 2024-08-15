package main

import (
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/xerrors"
)

func main() {
	err := errors.New("test")
	fmt.Println(err)
	fmt.Printf("%+v", err)
	err1 := xerrors.New("test")
	fmt.Println(err1)
	fmt.Printf("%+v", err1)
	err = errors.Wrap(err, "test1")
	fmt.Printf("%+v", err)
	err = errors.Errorf("test2")
	fmt.Println(err)
	err1 = xerrors.Errorf("test")
	fmt.Println(err1)
}
