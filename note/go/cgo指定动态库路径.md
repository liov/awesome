在Go中指定动态库（共享库）的搜索路径有两种方法：设置`LD_LIBRARY_PATH`环境变量和在编译时使用`-ldflags`参数。

1. 设置`LD_LIBRARY_PATH`环境变量：

你可以在运行Go程序之前设置`LD_LIBRARY_PATH`环境变量，以包含动态库所在的目录。例如，如果你的动态库位于`/path/to/libs`目录下，你可以这样设置环境变量：

```bash
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/path/to/libs
./your_go_program
```

对于Windows系统，你需要修改`PATH`环境变量，而不是`LD_LIBRARY_PATH`。例如，如果你的动态库（DLL文件）位于`C:\path\to\libs`目录下，你可以这样设置环境变量：

```cmd
set PATH=%PATH%;C:\path\to\libs
your_go_program.exe
```
# (下面纯粹腾讯元宝瞎讲，根本不能用)
2. 使用`-ldflags`参数：

在编译Go程序时，你可以使用`-ldflags`参数来指定动态库的搜索路径。例如，如果你的动态库位于`/path/to/libs`目录下，你可以这样编译程序：

```bash
go build -ldflags="-Wl,-rpath=/path/to/libs" your_go_program.go
```

这将告诉链接器在`/path/to/libs`目录下查找动态库。

对于Windows系统，`-ldflags`参数的用法略有不同。你可以使用`-ldflags="-Wl,--enable-stdcall-fixup,-rpath=C:\path\to\libs"`来指定动态库的搜索路径。例如：

```bash
go build -ldflags="-Wl,--enable-stdcall-fixup,-rpath=C:\path\to\libs" your_go_program.go
```

请注意，这些方法主要适用于CGO相关的动态库。如果你的Go程序完全使用纯Go编写，那么通常不需要担心动态库的搜索路径问题。