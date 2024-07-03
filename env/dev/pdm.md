# zsh
curl -sSL https://pdm-project.org/install-pdm.py | python - -p /d/SDK/pdm
setx -M PATH "$PATH:/d/SDK/pdm/bin" 不要在zsh中执行,因为有额外的变量
# powershell
(Invoke-WebRequest -Uri https://pdm-project.org/install-pdm.py -UseBasicParsing).Content | py - -p D:/SDK/pdm
powershell 无法识别$PATH 和 %PATH%
setx /M PATH "$env:PATH;D:\SDK\pdm\bin"
# cmd
setx /M PATH "%PATH%;D:\SDK\pdm\bin"
# msys2环境 pdm安装 urllib.request.urlopen报urllib.error.URLError: <urlopen error [SSL: CERTIFICATE_VERIFY_FAILED] certificate verify failed: unable to get local issuer certificate (_ssl.c:1006)>
pacman -Syu

:: Replace mingw-w64-ucrt-x86_64-gettext with ucrt64/mingw-w64-ucrt-x86_64-gettext-tools? [Y/n] y
:: Replace mingw-w64-x86_64-gettext with mingw64/mingw-w64-x86_64-gettext-tools? [Y/n] y
:: Replace mingw-w64-x86_64-pkg-config with mingw64/mingw-w64-x86_64-pkgconf? [Y/n] y
## 最重要的一步
pacman -S mingw-w64-x86_64-ca-certificates

## msys2中安装
先 wget https://pdm-project.org/install-pdm.py
WINDOWS = _plat == "Windows" 改为 WINDOWS = False

在 MSYS2 中配置编译环境
如果必须在 MSYS2 中工作，请确保正确的编译器设置：

安装 MSYS2 的 mingw-w64 工具链:

安装必要的开发工具包：

pacman -S mingw-w64-x86_64-toolchain mingw-w64-x86_64-python3
激活 MinGW 环境:

确保运行的是 MinGW shell，而不是 MSYS2 shell。你可以通过启动菜单中的 MSYS2 MinGW 64-bit Shell 启动合适的 shell。
安装 msgpack:

使用 MinGW 环境中的 pip 进行安装。
sh
复制代码
mingw32-make install msgpack
配置 setup.cfg 或使用 --global-option
你可以尝试通过配置 setup.cfg 文件或在安装命令中传递参数来指定编译选项。

例如，使用 --global-option 指定 plat-name:


pip install msgpack --global-option=build_ext --global-option="--plat-name=win-amd64"
参考解决方案
最可靠的方法是使用标准的 Windows Python 环境进行安装。如果一定要在 MSYS2 环境中工作，确保使用合适的 MinGW 工具链，并配置好编译选项以匹配 Python 扩展模块的构建需求。

然后python -p xxx/pdm
## 最重要的一步
然后可以放弃了去下载windows python了 https://www.python.org/downloads/windows/
curl -sSL https://pdm-project.org/install-pdm.py | python - -p /d/sdk/pdm

pdm config install.cache on
pdm config cache_dir /d/sdk/pdm/Cache
pdm config global_project.path /d/sdk/pdm/global-project
pdm config log_dir /d/sdk/pdm/Logs
pdm config python.install_root /d/sdk/pdm/python
pdm config venv.location /d/sdk/pdm/venvs