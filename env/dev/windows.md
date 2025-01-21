## win11专业版（报错）
slmgr -ipk W269N-WFGWX-YVC9B-4J6C9-T83GX

slmgr -skms kms.0t.net.cn

slmgr -ato

## massgrave
https://massgrave.dev/#download--how-to-use-it
网上搜个密钥
如 Windows11 激活密钥：NBQWQ-W9PTV-B4YWP-4K773-T6PKG
断网,系统更改密钥输入密钥,安装重启
```powershell
irm https://get.activated.win | iex
```
# 彻底关闭windows defender

win+R gpedit.msc
在【本地组策略编辑器】-【计算机配置】中依次打开【管理模板】、【Windows组件】、【Microsoft Defender防病毒】
在【Microsoft Defender防病毒】中双击【关闭Microsoft Defender防病毒】。选择【已启用】
然后双击【允许反恶意软件服务始终保持运行状态】。选择【已禁用】
双击【实时保护】。【关闭实时保护】 选择【已启用】
双击【扫描所有下载文件和附件】。【已禁用】
打开Microsoft Defender 全关
鼠标右击任务栏，点击【任务管理器】，【进程】和【启动】，相关的全部【关闭】和【禁用】。
# windows修改盘符
盘符修改是指更改电脑中分区或设备的驱动器号1。修改盘符的一般步骤是234：
按下组合键“Win+R”，输入“diskmgmt.msc”后按“回车”，打开磁盘管理器。
找到要修改盘符的分区或设备，右键点击，选择“更改驱动器号和路径”。
在弹出的窗口中，点击“更改”，从下拉列表中选择一个未被占用的盘符，点击“确定”。

# windows 端口占用
netstat -aon|findstr "8080"
taskkill /f /pid 12732

# windows 强制关应用
taskkill /im workwinlm.exe -f -t
taskkill /im system.dll -f -t

# 修改环境变量
## powershell
[Environment]::SetEnvironmentVariable("PATH", $env:PATH + ";C:\path\to\your\directory", "Machine")
如果你想只在当前用户下修改PATH，可以将"Machine"改为"User"。

## cmd
setx PATH "%PATH%;C:\path\to\your\directory"
这里"%PATH%是系统+用户的,设置的是用户的,setx /M PATH  系统的
其他终端慎用啊,直接清空了我的PATH

## zsh 
setx PATH "$PATH;/c/path/to/your/directory"

# 7zip
https://www.7-zip.org/download.html