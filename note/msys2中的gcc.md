pacman -Syu

MSYS2 中的不同版本 GCC 工具链
# Mingw-w64 with MSVCRT (GNU C Library Runtime) `mingw-w64开头后无运行时会安装mingw目录下`

32-bit: mingw-w64-i686-gcc
64-bit: mingw-w64-x86_64-gcc
描述: 使用 MSVCRT 作为 C 运行时库的 MinGW-w64 工具链。适用于需要与传统 Windows 环境兼容的应用程序。
安装命令:

pacman -S mingw-w64-i686-gcc
pacman -S mingw-w64-x86_64-gcc

## Mingw-w64 with Clang 即使clang也是mingw的clang

32-bit: mingw-w64-i686-clang
64-bit: mingw-w64-x86_64-clang
描述: 使用 Clang 作为编译器的 MinGW-w64 工具链。提供了 GCC 的替代选择，并且可以利用 Clang 的现代编译器特性。
安装命令:

pacman -S mingw-w64-i686-clang
pacman -S mingw-w64-x86_64-clang


# Mingw-w64 with UCRT (Universal C Runtime) `mingw-w64-ucrt 开头会安装ucrt目录下`

64-bit: mingw-w64-ucrt-x86_64-gcc
描述: 使用 UCRT 作为 C 运行时库的 MinGW-w64 工具链。适用于需要现代 C 运行时库特性的应用程序。
安装命令:

pacman -S mingw-w64-ucrt-x86_64-gcc

# Mingw-w64 with clang `mingw-w64-clang开头会安装在clang目录下`
确认有这个包
pacman -S mingw-w64-clang-x86_64-gcc


# pacman -S gcc
在 MSYS2 环境中，运行 pacman -S gcc 会安装一个基于 MSYS2 的 GCC 编译器。这是一个 POSIX 兼容环境下的编译器，主要用于编译需要与 MSYS2 环境兼容的程序。

MSYS2 中的 gcc
当你安装 gcc 包时，你得到的是一个用于 MSYS2 自身环境的 GCC 编译器。这与 MinGW-w64 工具链（用于生成原生 Windows 二进制文件）的编译器不同。MSYS2 的 GCC 编译器通常用于编译那些依赖于 MSYS2 提供的 POSIX 环境的程序。

安装 gcc
运行以下命令来安装 MSYS2 环境中的 GCC：


pacman -S gcc
验证安装
安装完成后，你可以运行以下命令来检查 GCC 的版本：


gcc --version
你应该会看到类似如下的输出，表示 GCC 已成功安装：


gcc (Rev2, Built by MSYS2 project) x.x.x
Copyright (C) 201x Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
使用场景
使用 pacman -S gcc 安装的 GCC 编译器适用于以下场景：

编译需要 POSIX 兼容环境的程序：如果你正在编写或编译需要 POSIX 功能的程序（例如使用 Unix 系统调用和库），则需要使用 MSYS2 的 GCC 编译器。

开发和测试：在 MSYS2 环境中开发和测试需要 POSIX 兼容性的程序。

区别于 MinGW-w64 GCC
MSYS2 GCC：生成的二进制文件依赖于 MSYS2 的 POSIX 环境，需要 MSYS2 运行时才能正常运行。
MinGW-w64 GCC：生成的二进制文件是原生的 Windows 应用程序，可以在没有任何额外运行时支持的情况下直接在 Windows 上运行。


# clangarm64
在 MSYS2 环境中，clangarm64 指的是用于编译 ARM64 架构（即 AArch64 架构）应用程序的 Clang 编译器工具链。MSYS2 提供了多个不同架构的编译工具链，以支持跨平台开发和编译。

clangarm64 工具链
架构: ARM64 (AArch64)
编译器: Clang/LLVM
用途: 生成适用于 ARM64 架构的原生 Windows 二进制文件。
安装 clangarm64
你可以使用以下命令来安装 clangarm64 工具链：


pacman -S mingw-w64-clang-aarch64
验证安装
安装完成后，你可以运行以下命令来检查 Clang 编译器的版本：


aarch64-w64-mingw32-clang --version
你应该会看到类似如下的输出，表示 Clang 已成功安装：

clang version x.x.x (tags/RELEASE_XY/final)
Target: aarch64-w64-mingw32
Thread model: posix
InstalledDir: /usr/bin
使用场景
使用 clangarm64 工具链适用于以下场景：

跨平台开发: 如果你正在为 ARM64 架构的设备开发应用程序，这个工具链将非常有用。例如，开发适用于 ARM64 版本的 Windows 10 或 Windows 11 的应用程序。
性能优化: 为 ARM64 设备优化和编译程序，以充分利用该架构的性能优势。
嵌入式开发: 开发嵌入式系统或物联网设备上运行的应用程序，这些设备通常采用 ARM64 架构。
使用 clangarm64
安装 clangarm64 工具链后，你可以在 MSYS2 环境中使用该工具链进行编译。例如，编译一个简单的 C 程序：

创建一个 hello.c 文件：


#include <stdio.h>

int main() {
printf("Hello, ARM64!\n");
return 0;
}
使用 clangarm64 工具链进行编译：


aarch64-w64-mingw32-clang hello.c -o hello_arm64.exe
生成的 hello_arm64.exe 二进制文件可以在 ARM64 架构的 Windows 设备上运行。

总结
clangarm64 是一个用于编译 ARM64 架构应用程序的 Clang 编译器工具链，适用于跨平台开发、性能优化和嵌入式开发等场景。通过安装和使用这个工具链，你可以为 ARM64 设备开发和编译高效的原生 Windows 应用程序。

# 工具链记得区分平台
windows平台的用msvc的一律统一,比如cmake,msys2移植的没法在非msys2环境直接用