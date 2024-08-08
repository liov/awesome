-buildmode=pie 是一个 Go 语言编译器选项，用于生成位置无关的可执行文件（Position Independent Executable，简称 PIE）。PIE 是一种安全特性，它使得程序在运行时可以被加载到内存的任意位置，从而提高了安全性。

在 Go 语言中，可以使用 -buildmode=pie 选项来生成 PIE 可执行文件。例如：

go build -buildmode=pie -o myapp_pie myapp.go
这将编译 myapp.go 文件并生成一个名为 myapp_pie 的 PIE 可执行文件。

需要注意的是，生成 PIE 可执行文件可能会增加程序的体积，因为编译器需要添加额外的代码来实现位置无关性。此外，并非所有的操作系统和架构都支持 PIE。在使用 -buildmode=pie 选项之前，请确保你的目标平台支持 PIE。