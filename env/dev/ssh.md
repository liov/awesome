SSH 的工作原理
默认的 SSH 客户端允许您通过安全通道连接到 Linux 工作站，并且默认使用 TCP 端口 22。SSH 对数据包进行编码，使任何人都无法观察您的活动。与不加密数据包的 telnet 不同，telnet 使其他人有可能读取您使用 SSH 所做的操作。在本文中，我们将向您展示如何在 Linux Mint 上启用安全 shell (ssh)。

如何安装Openssh服务器
您可以通过打开终端并在其中键入以下命令来安装 OpenSSH 服务器。

$ sudo apt install openssh-server
SSH 应自动设置为在系统启动时启动，并且应在安装后运行。然而，我们不会只是假设，而是会仔细检查。



要检查 SSH 是否已启用并在系统启动时自动启动，请运行下面给出的命令。

$ sudo systemctl is-enabled ssh
如果返回“启用”，则 SSH 应在计算机启动时立即启动。如果它被禁用或者状态为非活动状态，如下图所示：

然后使用下面提到的命令来启用它：

$ sudo systemctl enable ssh
现在，您可以通过键入以下内容来启动 SSH 服务：

$ sudo systemctl start ssh
并检查状态：

$ sudo systemctl status ssh
我们还可以使用“systemctl status”来接收所有信息的快速摘要；在上图中，我们可以看到服务已启动并正在运行以及其他有用的详细信息。


如何在防火墙中允许 ssh 连接
有时防火墙会限制您使用客户端和服务器之间的连接。因此，要允许您需要输入。

$ sudo ufw allow ssh
这将在您的防火墙上添加规则以允许 ssh 连接。如果当前防火墙已禁用，稍后您可以通过键入来启用防火墙。

$ sudo ufw enable
您需要刷新新设置才能实施它们，您可以通过键入来完成此操作。

$ sudo ufw reload
在上面的命令中，UFW是一个“简单的防火墙”，用于管理Linux防火墙：



您还可以通过键入来检查 ufw 防火墙的状态和完整详细信息。

$ sudo ufw status verbose