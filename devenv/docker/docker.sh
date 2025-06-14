# docker rmi:删除所有镜像
docker rmi `docker images -q`
docker rmi $(docker images -q)

挂载目录 映射端口
docker run -d -v /data/timepill:/data/timepill -p port:port $image_id

要获取所有容器名称及其IP地址只需一个命令。
docker inspect -f '{{.Name}} - {{.NetworkSettings.IPAddress }}' $(docker ps -aq)

# docker system prune:清理
docker system prune
## 日志
{
"log-driver": "json-file",
"log-opts": {"max-size": "10m", "max-file": "3"}
}

ls -lh $(find /var/lib/docker/containers/ -name *-json.log)

docker run --rm -v /var/lib/docker:/var/lib/docker alpine sh -c "echo '' > $(docker inspect --format='{{.LogPath}}' CONTAINER_NAME)"


truncate -s 0 /var/lib/docker/containers/*/*-json.log
sudo truncate -s 0 `docker inspect --format='{{.LogPath}}' <container>`

# 要删除所有标签为 <none> 的 Docker 镜像，可以使用以下命令：

docker rmi $(docker images -q -f "dangling=true")
这条命令首先使用 docker images -q -f "dangling=true" 查找所有没有标签（即 <none>）的镜像的 ID，并且 -q 参数使得命令只输出镜像ID，随后 docker rmi 命令利用这些 ID 删除这些镜像。

请注意，在执行这个操作之前，确保没有正在运行的容器依赖于这些镜像，否则 docker rmi 命令会失败。如果需要同时停止并删除依赖于这些镜像的容器，可以先执行以下命令：

Bash
docker stop $(docker ps -a -q -f "ancestor=<none>")
docker rm $(docker ps -a -q -f "ancestor=<none>")
但请注意，上面的 ancestor=<none> 过滤条件可能不会直接按预期工作，因为 <none> 镜像通常不直接作为容器的 ancestor 列出。更安全的做法是先检查容器状态和关联镜像，确保不会误删重要容器。通常，直接删除 <none> 镜像是安全的，因为这些镜像不再被任何容器使用。

# 删除所有名字中带 “provider”
docker rmi $(docker images | grep "provider" | awk '{print $3}')
# 查看容器ip
docker inspect d7f29df68dd4 | grep IPAddress
# 删除所有容器
docker rm -f `docker ps -a -q`
# 自动重启
docker run --restart=always
docker update --restart=always <CONTAINER ID>
# root用户
docker run --user="root"
# 覆盖entrypoint
docker run --entrypoint /bin/bash
# 执行后删除
docker run --rm
# dind 特权模式
docker run --user="root" --privileged.
# 细粒度权限控制
--cap-add=NET_ADMIN
# 打包镜像到文件
docker save -o <file> <image>
docker load -i <file>
# 下表列出了Linux功能选项，这些选项是默认允许的，可以删除
SETPCAP |修改进程的权限

MKNOD |使用mknod(2)创建特殊文件

AUDIT_WRITE |将记录写入内核审计日志

CHOWN |任意更改文件UIDs和GIDs(见chown(2))

NET_RAW |使用 RAW 和 PACKET 套接字

DAC_OVERRIDE |绕过文件的读、写和执行权限检查

FOWNER |绕过对进程的文件系统UID与文件的UID进行权限匹配的检查操作

FSETID |当文件被修改时，不要清除set-user-ID和set-group-ID权限位

KILL |绕过发送信号的权限检查

SETGID |自定义处理进程GID和补充GID列表

SETUID |自定义处理进程UID

NET_BIND_SERVICE |将套接字绑定到互联网域名专用端口(端口号小于1024)。

SYS_CHROOT |使用chroot(2)，更改根目录

SETFCAP |设置文件功能

## 下表显示了默认情况下未授予的功能，可以手动添加这些功能
SYS_MODULE |加载和卸载内核模块

SYS_RAWIO |执行I / O端口操作(iopl(2)和ioperm(2))

SYS_PACCT |使用acct(2)，打开或关闭进程计数

SYS_ADMIN |执行一系列系统管理操作

SYS_NICE |提高进程的nice值(nice(2)，setpriority(2))并更改任意进程的nice值

SYS_RESOURCE |覆盖资源限制

SYS_TIME |设置系统时钟(settimeofday(2)，stime(2)，adjtimex(2)); 设置实时(硬件)时钟

SYS_TTY_CONFIG |使用vhangup(2); 在虚拟终端上使用各种特权的ioctl(2)操作

AUDIT_CONTROL |启用和禁用内核审核； 更改审核过滤器规则； 检索审核状态和过滤规则

MAC_ADMIN |允许MAC配置或状态更改。 为Smack LSM而实现的功能

MAC_OVERRIDE |覆盖强制访问控制(MAC)。 为Smack Linux安全模块(LSM)实现

NET_ADMIN |执行各种与网络相关的操作

SYSLOG |执行syslog(2)的权限操作。

DAC_READ_SEARCH |绕过文件读取权限检查以及目录读取和执行权限检查

LINUX_IMMUTABLE |设置FS_APPEND_FL和FS_IMMUTABLE_FL i-node 标志

NET_BROADCAST |使套接字可以实现广播，并监听广播包

IPC_LOCK |锁定内存(mlock(2)，mlockall(2)，mmap(2)，shmctl(2))

IPC_OWNER |绕过对System V IPC对象操作的权限检查

SYS_PTRACE |使用ptrace(2)跟踪任意进程

SYS_BOOT |使用reboot(2)和kexec_load(2)，重新引导并加载新内核以供程序执行

LEASE |在任意文件上建立Lease租约(请参阅fcntl(2))

WAKE_ALARM |触发唤醒系统的操作

BLOCK_SUSPEND |开启可以阻止系统挂起的功能

如果想与系统的网络堆栈进行交互，应该使用`--cap-add=NET_ADMIN`来修改网络接口


docker run --rm -v /mnt/d/SDK/gopath:/go -v $PWD:/work -w /work/server/go golang go run /work/server/go/tools/install.go

# 使用宿主机网络
docker run --net=host

docker buildx build --platform linux/amd64,linux/arm64,linux/arm/v7 -t gcr.io/my-project/my-image:latest .


# 日志清理
#!/bin/sh

echo "======== start clean docker containers logs ========"

logs=$(find /var/lib/docker/containers/ -name *-json.log)

for log in $logs
        do
                echo "clean logs : $log"
                cat /dev/null > $log
        done

echo "======== end clean docker containers logs ========"

## 删除日志文件
# find /var/lib/docker/containers/ -type f -name "*-json.log" -delete # 删除
# find /var/lib/docker/containers/ -type f -name "*-json.log" -exec rm {} \; # 删除
# k8s需要一起删除/var/log/pods下的文件 find /var/log/pods/ -type f -name "*.log" -delete;
不要删除文件，如果删除了，sudo systemctl restart docker

## 清空日志文件
# find /var/lib/docker/containers/ -type f -name "*-json.log" -exec truncate -s 0 {} \; 清空

# 日志清理配置
/etc/docker/daemon.json
{
  "log-driver": "json-file",
  "log-opts": {
    "max-size": "10m",  // 设置每个日志文件的最大大小为10MB
    "max-file": "3"     // 设置保留的日志文件数量
  }
}
