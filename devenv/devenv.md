# msys2
msys2很好
但对一些语言不是那么友好，比如python, nodejs等
这些软件的官方版本是msvc编译的,因此大量需要与C交互的库都是msvc编译的,因此使用gnu的版本就会时不时受到限制,例如python使用numpy就要安装gnu版本的
还好，rust提供了msvc和gnu两个版本，写扩展都可以使用rust
# go
https://golang.google.cn/dl/go1.22.5.windows-amd64.zip
## linux
# [go](https://golang.google.cn/doc/)
wget https://dl.google.com/go/go1.xx.x.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.xx.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
export GOPRIVATE=go.hoper.xyz
export GOPROXY=https://goproxy.io,direct

# [protoc](https://github.com/protocolbuffers/protobuf/releases)


# rust
curl https://sh.rustup.rs -sSf | sh

# ssh

powershell ssh
```
Add-WindowsCapability -Online -Name OpenSSH-Client
  ssh root@IP  -p PORT -i .\.ssh\id_rsa
ssh -R [local port]:[remote host]:[remote port] [SSH hostname]
ssh  -fNg -L <本地端口>:<服务器数据库地址>  <用户名>@<服务器地址>
想让SSH连接一直连接，可以加上 -NTf 参数。
exit
```


# gradle maven
```groovy
maven { url 'https://maven.aliyun.com/repository/public' }//central和jcenter的聚合
maven { url 'https://maven.aliyun.com/repository/central' }
maven { url 'https://maven.aliyun.com/repository/jcenter' }
maven { url 'https://maven.aliyun.com/repository/google'}
maven { url 'https://maven.aliyun.com/repository/gradle-plugin' }
maven { url "https://jitpack.io" }
```

# etcd
etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/kubernetes/pki/etcd/ca.crt --cert=/etc/kubernetes/pki/etcd/peer.crt --key=/etc/kubernetes/pki/etcd/peer.key member list

peer.xxx节点用 server.xxx外部用

# [helm](https://helm.sh/)
wget https://get.helm.sh/helm-v3.6.0-linux-amd64.tar.gz
tar -zxvf helm-v3.0.0-linux-amd64.tar.gz
mv linux-amd64/helm /usr/local/bin/helm
---------------------------------------
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

helm repo add gitlab https://charts.gitlab.io/
helm repo add aliyun https://kubernetes.oss-cn-hangzhou.aliyuncs.com/charts
helm repo add incubator https://kubernetes-charts-incubator.storage.googleapis.com/

# linux
vim /etc/profile
export PATH=$PATH:/usr/local/go/bin:/usr/local/node/bin
export HISTCONTROL=ignoredups
source /etc/profile

# netcat
https://eternallybored.org/misc/netcat/


# vsftp 

```bash
sudo apt install vsftpd

sudo passwd jyb

mkdir /home/jyb/ftp

chmod 777 -R /home/jyb/ftp

sudo vim /etc/vsftpd.conf

connect_from_port_21=YES

local_root=/home/jyb/ftp

allow_writeable_chroot=YES

将#chroot_local_user=YES前的注释去掉

pam_service_name=ftp原配置中为vsftpd，ubuntu用户需要更改成ftp

sudo service vsftpd start

sudo service vsftpd restart

```