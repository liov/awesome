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
# git
pacman -S git(不建议这样安装,自行搜索windows版本安装添加PATH)
export PATH=$PATH:"/d/Program Files/Git/cmd"
命令行里输入git version 可以查看版本：
# zsh
pacman -S zsh
export PATH=$PATH:"/d/Program Files/Git/cmd"

# Oh My Zsh
参考 dev/zsh.md/oh my zsh部分


在source ~/.zshrc这行之前添加PATH
~~PATH=/mingw64/bin:/ucrt64/bin:/usr/bin:/usr/local/bin:$PATH:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl~~
/mingw64/bin 建议直接加入windows PATH win7
/ucrt64/bin 较新的运行时,可与mingw64二选一,未测试cgo是否可用,建议直接加入windows PATH win10+ 建议用这个
/usr/local/bin:/bin 这两个目录实际不存在
/usr/bin 建议直接加入windows PATH
PATH=$PATH:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl


执行下 source ~/.zshrc激活插件。

# 配置右键打开终端(废弃，请使用Windows Terminal)

新建mingw64.reg后缀文件：

Windows Registry Editor Version 5.00

[HKEY_CLASSES_ROOT\Directory\Background\shell\mingw64]
@="MinGW64 Here"
"icon"="C:\\msys64\\ucrt64.exe"

[HKEY_CLASSES_ROOT\Directory\Background\shell\mingw64\command]
@="C:\\msys64\\msys2_shell.cmd -ucrt64 -here"
双击导入即可。

# 配置 Windows Terminal 使用 msys2 bash
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

