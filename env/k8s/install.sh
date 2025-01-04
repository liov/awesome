# docker
../docker/install.md
# k8s
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
## conntrack
sudo apt install conntrack
## crictl
# 下载最新版本的 crictl (以 v1.32.0 为例)
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.32.0/crictl-v1.32.0-linux-amd64.tar.gz
# 解压文件
tar zxvf crictl-v1.32.0-linux-amd64.tar.gz -C /usr/local/bin
# 清理下载文件
rm crictl-v1.32.0-linux-amd64.tar.gz
## cri
# 添加官方 GPG 密钥
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

# 设置稳定版仓库
echo "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

# 更新包索引并安装 Docker Engine
sudo apt-get update
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# 下载并安装 cri-dockerd
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.2.7/cri-dockerd_0.2.7_amd64.deb
sudo dpkg -i cri-dockerd_0.2.7_amd64.deb

sudo sysctl fs.protected_regular=0

sudo mkdir -p /opt/cni/bin
curl -L https://github.com/containernetworking/plugins/releases/download/v1.6.1/cni-plugins-linux-amd64-v1.6.1.tgz | sudo tar zxvf -C /opt/cni/bin
export MINIKUBE_CNI_CONF_DIR=/etc/cni/net.d
export PATH=$PATH:/opt/cni/bin
minikube start --network-plugin=cni

minikube start --driver=none  --image-mirror-country='cn' --extra-config=apiserver.service-node-port-range=1-39999  --extra-config=kube-proxy.mode=ipvs --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0 --bootstrapper=kubeadm --apiserver-ips=0.0.0.0 --apiserver-port=6443

❗  kubectl 和 minikube 配置将存储在 /root 中
❗  如需以您自己的用户身份使用 kubectl 或 minikube 命令，您可能需要重新定位该命令。例如，如需覆盖您的自定义设置，请运行：

    ▪ sudo mv /root/.kube /root/.minikube $HOME
    ▪ sudo chown -R $USER $HOME/.kube $HOME/.minikube

💡  此操作还可通过设置环境变量 CHANGE_MINIKUBE_NONE_USER=true 自动完成
真扯淡,还要把$HOME/.kube/config 中路径中 /root改为..
# helm

$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh


minikube addons enable dashboard
minikube addons enable logviewer
minikube addons enable efk
minikube addons enable helm-tiller

# prometheus
kubectl create namespace monitoring
helm install kube-prometheus prometheus-community/kube-prometheus-stack -f helm.yaml -n monitoring
# apisix
cp -r /var/lib/minikube/certs/etcd /root/certs/ && chmod 666 /root/certs/etcd/server.key
kubectl create namespace ingress-apisix
# acme
kubectl apply -f tls.yaml
# tools
kubectl create namespace tools

