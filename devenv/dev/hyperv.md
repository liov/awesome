双系统双引导
需要两个盘
一个装windows
一个装mint,装的时候一定要手动分区,特别是efi分区,然后grub引导装这个分区里
windows磁盘管理里装mint那个盘脱机
windows打开hyperv服务,创建虚拟机,磁盘选稍后添加
创建成功后,SCSI控制器添加那个物理磁盘,安全启动关掉，检查点关掉
连接启动,只能进grub引导,启动后黑屏
无意间搞成功了
那就是进grub引导重装内核
然后重启
没测过还能不能从物理盘启动
好像对usb移动硬盘有影响,一直跳针
重启发现又不行了
设置里集成服务打开信号检测和来宾服务又可以了
重启又不行了
偶然的可行
确实是偶然可行
总结，hyperv是垃圾
试试vmware
vmware可以
但是建虚拟机的时候要选windows11,后面才能选uefi引导(linux选ubuntu64位好像也行)
网络选桥接，在编辑中手动选网卡
坑: 数字键盘不能用

sudo apt install open-vm-tools-desktop
sudo apt install python3-pip