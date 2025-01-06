choco install minikube
curl kubectl

minikube ssh
sudo passwd docker


curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
# docker
apt install docker.io
vim /etc/docker/daemon.json
{
"registry-mirrors": ["https://docker.mirrors.ustc.edu.cn"],
"insecure-registries":["${ip}"],
}
## upgrade
apt upgrade doker

## cri-dockerd
### 二进制（未验证）
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.14/cri-dockerd-0.3.14.amd64.tgz
tar zxvf cri-dockerd-0.3.14.amd64.tgz
sudo mv cri-dockerd/cri-dockerd /usr/bin/cri-dockerd

vim /etc/systemd/system/cri-docker.service
```ini
[Unit]
Description=CRI Interface for Docker Application Container Engine
Documentation=https://docs.mirantis.com
After=network-online.target firewalld.service docker.service
Wants=network-online.target
Requires=cri-docker.socket

[Service]
Type=notify
ExecStart=/usr/bin/cri-dockerd --network-plugin=cni --pod-infra-container-image=registry.aliyuncs.com/google_containers/pause:3.7
ExecReload=/bin/kill -s HUP $MAINPID
TimeoutSec=0
RestartSec=2
Restart=always

StartLimitBurst=3

StartLimitInterval=60s

LimitNOFILE=infinity
LimitNPROC=infinity
LimitCORE=infinity

TasksMax=infinity
Delegate=yes
KillMode=process

[Install]
WantedBy=multi-user.target
```
vim /usr/lib/systemd/system/cri-docker.socket
```ini
[Unit]
Description=CRI Docker Socket for the API
PartOf=cri-docker.service

[Socket]
ListenStream=%t/cri-dockerd.sock
SocketMode=0660
SocketUser=root
SocketGroup=docker

[Install]
WantedBy=sockets.target
```
systemctl daemon-reload ; systemctl enable cri-docker --now
systemctl is-active cri-docker
### deb (可行)
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.3.14/cri-dockerd_0.3.14.3-0.ubuntu-focal_amd64.deb
sudo dpkg -i cri-dockerd_0.3.14.3-0.ubuntu-focal_amd64.deb
sudo systemctl daemon-reload
sudo systemctl enable cri-docker.socket
sudo systemctl start cri-docker.socket cri-docker
cri-dockerd --version
ls -al /var/run/cri-dockerd.sock

## cni plugins
wget https://github.com/containernetworking/plugins/releases/download/v1.5.0/cni-plugins-linux-amd64-v1.5.0.tgz
tar -xf cni-plugins-linux-amd64-v1.5.0.tgz -C /opt/cni/bin
sudo systemctl restart docker
sudo mkdir -p /etc/cni/net.d
### shell
plugin.sh
```shell
CNI_PLUGIN_VERSION="<version_here>"
CNI_PLUGIN_TAR="cni-plugins-linux-amd64-$CNI_PLUGIN_VERSION.tgz" # change arch if not on amd64
CNI_PLUGIN_INSTALL_DIR="/opt/cni/bin"
curl -LO "https://github.com/containernetworking/plugins/releases/download/$CNI_PLUGIN_VERSION/$CNI_PLUGIN_TAR"
sudo mkdir -p "$CNI_PLUGIN_INSTALL_DIR"
sudo tar -xf "$CNI_PLUGIN_TAR" -C "$CNI_PLUGIN_INSTALL_DIR"
rm "$CNI_PLUGIN_TAR"
```

# minikube
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
# kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
curl -LO https://dl.k8s.io/release/v1.30.0/bin/linux/amd64/kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# crictl
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.30.0/crictl-v1.30.0-linux-amd64.tar.gz

tar zxvf crictl-v1.30.0-linux-amd64.tar.gz
sudo mv crictl /usr/local/bin

vim /etc/crictl.yaml
```yaml
runtime-endpoint: unix:///var/run/docker.sock
timeout: 2
debug: false
```

cat > /etc/crictl.yaml <<EOF
runtime-endpoint: unix:///var/run/docker.sock
timeout: 2
debug: false
EOF
# start
minikube start --driver=none --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --extra-config=kube-proxy.mode=ipvs --extra-config=apiserver.advertise-address=0.0.0.0 --apiserver-ips=0.0.0.0 --apiserver-port=6443  --apiserver-name=localhost --extra-config=apiserver.service-node-port-range=1-39999 --extra-config=apiserver.authorization-mode=Node,RBAC --bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0 --cni calico(可选的)

etcd默认监听127.0.0.1和本机内部ip

sudo minikube start --driver=none  --image-mirror-country='cn' --extra-config=apiserver.service-node-port-range=1-39999  --extra-config=kube-proxy.mode=ipvs --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0 --bootstrapper=kubeadm --apiserver-ips=0.0.0.0,:: --apiserver-port=6443 --extra-config=apiserver.service-cluster-ip-range=10.96.0.0/12,fd00:10::/108 --extra-config=controller-manager.cluster-cidr=10.96.0.0/12,fd00:10::/108 --service-cluster-ip-range=10.96.0.0/12,fd00:10::/108 --extra-config=kube-proxy.cluster-cidr=10.96.0.0/12,fd00:10::/108  --extra-config=controller-manager.node-cidr-mask-size-ipv4=12 --extra-config=controller-manager.node-cidr-mask-size-ipv6=64

sudo minikube start --driver=none  --image-mirror-country='cn' --extra-config=apiserver.service-node-port-range=1-39999  --extra-config=kube-proxy.mode=ipvs --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0 --bootstrapper=kubeadm --apiserver-ips=0.0.0.0 --apiserver-port=6443

❗  kubectl 和 minikube 配置将存储在 /root 中
❗  如需以您自己的用户身份使用 kubectl 或 minikube 命令，您可能需要重新定位该命令。例如，如需覆盖您的自定义设置，请运行：

    ▪ sudo mv /root/.kube /root/.minikube $HOME
    ▪ sudo chown -R $USER $HOME/.kube $HOME/.minikube

💡  此操作还可通过设置环境变量 CHANGE_MINIKUBE_NONE_USER=true 自动完成
真扯淡,还要把$HOME/.kube/config 中路径中 /root改为..


minikube addons enable dashboard
minikube addons enable logviewer
minikube addons enable efk
minikube addons enable helm-tiller

# 特定版本
curl -LO https://dl.k8s.io/release/v1.23.0/bin/linux/amd64/kubectl
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
-- 挂载data目录不成功，可能是权限问题
docker中部署Kubernetes
minikube start --driver=docker --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --mount --mount-string=$HOME:/host --cpus=4 --memory='8192M'
root 直接部署
minikube start --driver=none --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --extra-config=kube-proxy.mode=ipvs --extra-config=apiserver.advertise-address=0.0.0.0 --apiserver-ips=0.0.0.0 --apiserver-port=6443  --apiserver-name=localhost
-- port
--extra-config=apiserver.service-node-port-range=1-39999 
-- prometheus-operator
--extra-config=apiserver.authorization-mode=Node,RBAC #默认配置是这个 --extra-config=apiserver.authorization-mode=RBAC 官方文档是这个，怀疑后来重启logs报无权限跟这个有关
--kube-prometheus
--bootstrapper=kubeadm --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0

# 对外开放（试了没用啊）
--apiserver-ips=0.0.0.0 --apiserver-port=6443
## 对于 docker 和 podman 驱动程序，使用--listen-address标志：
--listen-address=0.0.0.0
外网通过代理访问docker中的服务
--url只打印url不自动打开浏览器

## 通过代理暴露集群内ip
kubectl proxy --port=8001 --address='0.0.0.0' --accept-hosts='^.*' &
curl http://[k8s-proxy-ip]:8001/api/v1/namespaces/[namespace-name]/services/[service-name]:80/proxy
curl http://[k8s-proxy-ip]:8001/api/v1/namespaces/[namespace-name]/pods/[pod-name]:8080/proxy
& 号将命令放到后台运行
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/http:kubernetes-dashboard:/proxy/#/overview?namespace=default

## 端口转发到本地
kubectl port-forward --address 0.0.0.0 service/${svcname} 8080:${svcport} -n ${namespace} --kubeconfig=stage/config

# 挂载目录
9P Mounts
9P mounts are flexible and work across all hypervisors, but suffers from performance and reliability issues when used with large folders (>600 files). See Driver Mounts as an alternative.

To mount a directory from the host into the guest using the mount subcommand:

minikube mount <source directory>:<target directory>
For example, this would mount your home directory to appear as /host within the minikube VM:

minikube mount $HOME:/host

# etcd 
## 直接使用minikube的etcd
etcdctl --endpoints=https://172.17.0.3:2379 --cacert=/var/lib/minikube/certs/etcd/ca.crt --cert=/var/lib/minikube/certs/etcd/server.crt --key=/var/lib/minikube/certs/etcd/server.key member list

