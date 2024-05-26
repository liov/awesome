package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
)

func main() {
	// -a "b"    -c 'c ' -d 1 --f aaa -e `aaa` -d ~!#!@
	fmt.Println(os.Args)
	// *.exe -a b -c c  -d 1 --f aaa -e `aaa` -d ~!#!@]
	commandLine := pflag.NewFlagSet(os.Args[0], pflag.ContinueOnError)
	var a string
	commandLine.StringVarP(&a, "aaa", "a", "1", "a")
	commandLine.StringVarP(&a, "bbb", "b", "2", "b")
	commandLine.Parse([]string{"go", " run", "--aaa", "b", "exit"})
	fmt.Println(a, commandLine.Args())

	commandLine.StringVarP(&a, "ccc", "c", "c", "c")
	commandLine.Parse([]string{"go", " run", "--ccc", "c", "exit"})
	fmt.Println(a)
}
