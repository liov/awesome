package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command(`D:/Download/gerbv-2.10.0-win64/bin/gerbv.exe`, "-a", "-x", "png", "-B", "0", "-D", "508", "-o", "output.png", "D:/xxx")
	fmt.Println(cmd.String())
	// 将命令的标准输出和标准错误都设置为当前进程的标准输出和标准错误
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd.Run())
}
