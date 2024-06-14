# terminal

## git bash
命令行: D:/Program Files/Git/bin/bash.exe -i -l
## msys2
直接打开安装目录下的各个终端均无法直接运行windows中的程序,原因是PATH没有包含windows的PATH
ucrt 较新的依赖, mingw较旧的(默认有gcc)
名称: ucrt64
命令行: D:/sdk/msys64/msys2_shell.cmd -defterm -no-start -use-full-path -here -ucrt64
图标: D:/sdk/msys64/ucrt64.ico

镜像地址：https://mirrors.tuna.tsinghua.edu.cn/help/msys2/

进入 D:\msys64\etc\pacman.d目录，分别进行如下操作：

编辑 /etc/pacman.d/mirrorlist.mingw32 ，在文件开头添加：

编辑 /etc/pacman.d/mirrorlist.mingw32 ，在文件开头添加：

Server = https://mirrors.tuna.tsinghua.edu.cn/msys2/mingw/i686
编辑 /etc/pacman.d/mirrorlist.mingw64 ，在文件开头添加：

Server = https://mirrors.tuna.tsinghua.edu.cn/msys2/mingw/x86_64
编辑 /etc/pacman.d/mirrorlist.ucrt64 ，在文件开头添加：

Server = https://mirrors.tuna.tsinghua.edu.cn/msys2/mingw/ucrt64
编辑 /etc/pacman.d/mirrorlist.clang64 ，在文件开头添加：

Server = https://mirrors.tuna.tsinghua.edu.cn/msys2/mingw/clang64
编辑 /etc/pacman.d/mirrorlist.msys ，在文件开头添加：

Server = https://mirrors.tuna.tsinghua.edu.cn/msys2/msys/$arch
点击安装路径的mingw64.exe启动，然后执行 pacman -Sy 刷新软件包数据即可。

配置好pacman 镜像源，就可以安装常用软件了。

安装常用软件#
git#
pacman -S git
命令行里输入git version 可以查看版本：

Oh My Zsh#
官方网站: http://ohmyz.sh

GitHub: https://github.com/ohmyzsh/ohmyzsh

pacman -S zsh

sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"

# 或者
# sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
所以这时的zsh 基本已经配置完成,你需要一行命令就可以切换到 zsh 模式，终端下输入zsh切换，输入bash切回去。

ohmyzsh插件，路径：

~/.oh-my-zsh/plugins/
~/.oh-my-zsh/custom/plugins/
新增插件示例：

git clone https://github.com/zsh-users/zsh-syntax-highlighting $ZSH_CUSTOM/plugins/zsh-syntax-highlighting
git clone https://github.com/zsh-users/zsh-autosuggestions $ZSH_CUSTOM/plugins/zsh-autosuggestions
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ${ZSH_CUSTOM:-$HOME/.oh-my-zsh/custom}/themes/powerlevel10k

启用插件：修改~/.zshrc文件，示例：
ZSH_THEME="powerlevel10k/powerlevel10k"

plugins=(git zsh-autosuggestions zsh-syntax-highlighting)
默认是plugins=(git)

在source $ZSH/oh-my-zsh.sh这行之前添加PATH
PATH=/ucrt64/bin:/usr/local/bin:/usr/bin:/bin:/c/Windows/System32:/c/Windows:/c/Windows/System32/Wbem:/c/Windows/System32/WindowsPowerShell/v1.0/:$PATH:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl
source $ZSH/oh-my-zsh.sh

执行下 source ~/.zshrc激活插件。

配置右键打开终端#

新建mingw64.reg后缀文件：

Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\Directory\Background\shell\mingw64]
@="MinGW64 Here"
"icon"="C:\\msys64\\ucrt64.exe"

[HKEY_CLASSES_ROOT\Directory\Background\shell\mingw64\command]
@="C:\\msys64\\msys2_shell.cmd -ucrt64 -here"
双击导入即可。

配置 Windows Terminal 使用 msys2 bash
核心参数：

名称：msys2

命令行：C:\msys64\usr\bin\bash.exe(这种好多命令不能用,原因是只有windows的PATH)

icon: C:\msys64\msys2.ico

{
"guid":"{1c4de342-38b7-51cf-b940-2309a097f589}",
"hidden":false,
"name":"Bash",
"commandline":"C:\\msys64\\usr\\bin\\bash.exe",
"historySize":9001,
"closeOnExit":true,
"useAcrylic":true,
"acrylicOpacity":0.85,
"icon":"C:\\msys64\\msys2.ico",
"startingDirectory":null
}


配置 zsh 为 bash 默认终端#(不需要这步,直接用zsh可执行文件)
编辑 ~/.bashrc，加入下面的几行。

# Launch Zsh
if [ -t 1 ]; then
exec zsh
fi
配置idea#
配置终端使用msys2 bash#
点击 File -> Settings -> Tools -> Terminal ，配置shell path:

C:\msys64\usr\bin\bash.exe
或者

C:\msys64\usr\bin\zsh.exe
这样就可以执行bash命令了。例如：

$ cd /d/Download/

$ ls | wc -l
81

$ which mvn
/e/opt/apache-maven-3.6.3/bin/mvn

$ cd ~
附录1: pacman常用命令#
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

附录2：解决vscode不识别git问题#
使用msys2环境安装git，vscode识别不出来git，在源代码管理菜单中显示当前打开的文件夹没有git存储库，打开的文件夹实际存在.git文件夹。

解决办法：

编写 git-wrap.bat：

@echo off
setlocal

rem If you don't add path for msys2 into %PATH%, enable following line.
rem set PATH=c:\msys64\usr\bin;%PATH%

if "%1" equ "rev-parse" goto rev_parse
git %*
goto :eof
:rev_parse
for /f %%1 in ('git %*') do cygpath -w %%1
将 git-wrap.bat放到某个文件夹，例如位于：c:\msys64\git-wrap.bat。

在vscode设置 git.path ：
点击： File -> Preferences -> User Settings, 搜索 git.path 将 git-wrap.bat写到配置文件里：

"git.path": "C:\\msys64\\git-wrap.bat",
重启vscode

Have fun!

附录3：Windows Terminal右键#
编写wt.reg:

Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\Directory\Background\shell\wt]
@="Windows Terminal Here"

[HKEY_CLASSES_ROOT\Directory\Background\shell\wt\command]
@="C:\\Users\\你的用户名\\AppData\\Local\\Microsoft\\WindowsApps\\wt.exe"
注意替换"你的用户名"。

参考#
1、MSYS2
https://www.msys2.org/

2、Install Terminal + git-bash + zsh + oh-my-zsh on Windows 10 | MiaoTony's小窝
https://miaotony.xyz/2020/12/13/Server_Terminal_gitbash_zsh/

3、msys2软件包管理工具pacman常用命令_hustlei的专栏

https://blog.csdn.net/hustlei/article/details/86687621

4、玩转 Ubuntu 20 桌面版 - 飞鸿影 - 博客园
https://www.cnblogs.com/52fhy/p/9571463.html



pacman install -S zsh
