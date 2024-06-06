# win10教育版
slmgr /ipk NW6C2-QMPVW-D7KKK-3GKT6-VCFB2

slmgr /skms kms.03k.org

slmgr /ato

## win11专业版
slmgr -ipk W269N-WFGWX-YVC9B-4J6C9-T83GX

slmgr -skms kms.0t.net.cn

slmgr -ato
## win激活，出现“非核心版本的计算机”的处理方法
1 打开“注册表编辑器”；（Windows + R然后输入 Regedit）

2 修改SkipRearm 的值为1；（在HKEY_LOCAL_MACHINE–》SOFTWARE–》Microsoft–》Windows NT–》CurrentVersion–》SoftwareProtectionPlatform里面，将SkipRearm的值修改为1）重启电脑

3 以管理员身份启动cmd，输入SLMGR   -REARM，根据提示，再次重启电脑！

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
