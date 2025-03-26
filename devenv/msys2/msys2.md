pacman常用命令#
pacman命令较多，作为新手，将个人最常用的命令总结如下：

pacman -Sy: 从服务器下载新的软件包数据库（实际上就是下载远程仓库最新软件列表到本地）。

pacman -Syu: 升级系统及所有已经安装的软件。

pacman -S 软件名: 安装软件。也可以同时安装多个包，只需以空格分隔包名即可。

pacman -Rs 软件名: 删除软件，同时删除本机上只有该软件依赖的软件。

pacman -Ru 软件名: 删除软件，同时删除不再被任何软件所需要的依赖。

pacman -Ssq 关键字: 在仓库中搜索含关键字的软件包，并用简洁方式显示。

pacman -Qs 关键字: 搜索已安装的软件包。

pacman -Qi 软件名: 查看某个软件包信息，显示软件简介,构架,依赖,大小等详细信息。

pacman -Sg: 列出软件仓库上所有的软件包组。

pacman -Sg 软件包组: 查看某软件包组所包含的所有软件包。

pacman -Sc：清理未安装的包文件，包文件位于 /var/cache/pacman/pkg/ 目录。

pacman -Scc：清理所有的缓存文件

不要用msys2的软件包，node(使用nvm)，python(使用pdm)

# 装机
pacman -S fish zip zsh
pacman -S mingw-w64-ucrt-x86_64-ffmpeg
pacman -S mingw-w64-ucrt-x86_64-opencv
pacman -S mingw-w64-ucrt-x86_64-toolchain


# pdm(放弃这个，请使用msvc python)
## msys2环境 pdm安装 urllib.request.urlopen报urllib.error.URLError: <urlopen error [SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed: unable to get local issuer certificate (_ssl.c:1006)>
pacman -S mingw-w64-x86_64-ca-certificates

## msys2中安装
先 wget https://pdm-project.org/install-pdm.py
WINDOWS = _plat == "Windows" 改为 WINDOWS = False

在 MSYS2 中配置编译环境
如果必须在 MSYS2 中工作，请确保正确的编译器设置：

安装 MSYS2 的 mingw-w64 工具链:

安装必要的开发工具包：

pacman -S mingw-w64-ucrt-x86_64-toolchain mingw-w64-ucrt-x86_64-python3
激活 MinGW 环境:

确保运行的是 MinGW shell，而不是 MSYS2 shell。你可以通过启动菜单中的 MSYS2 MinGW 64-bit Shell 启动合适的 shell。
安装 msgpack:

使用 MinGW 环境中的 pip 进行安装。
mingw32-make install msgpack
配置 setup.cfg 或使用 --global-option
你可以尝试通过配置 setup.cfg 文件或在安装命令中传递参数来指定编译选项。

例如，使用 --global-option 指定 plat-name:


pip install msgpack --global-option=build_ext --global-option="--plat-name=win-amd64"
参考解决方案
最可靠的方法是使用标准的 Windows Python 环境进行安装。如果一定要在 MSYS2 环境中工作，确保使用合适的 MinGW 工具链，并配置好编译选项以匹配 Python 扩展模块的构建需求。

然后python -p xxx/pdm

windows中的用户名下.ssh目录下文件复制一份到msys2 /home/username/.ssh中