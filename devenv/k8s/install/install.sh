# docker
../docker/install.md

## conntrack
sudo apt install conntrack
## crictl
# 下载最新版本的 crictl (以 v1.32.0 为例)
wget https://github.com/kubernetes-sigs/cri-tools/releases/download/v1.32.0/crictl-v1.32.0-linux-amd64.tar.gz
# 解压文件
tar zxvf crictl-v1.32.0-linux-amd64.tar.gz -C /usr/local/bin

# 下载并安装 cri-dockerd
wget https://github.com/Mirantis/cri-dockerd/releases/download/v0.2.7/cri-dockerd_0.2.7_amd64.deb
sudo dpkg -i cri-dockerd_0.2.7_amd64.deb

sudo sysctl fs.protected_regular=0

sudo mkdir -p /opt/cni/bin
curl -L https://github.com/containernetworking/plugins/releases/download/v1.6.1/cni-plugins-linux-amd64-v1.6.1.tgz | sudo tar zxvf -C /opt/cni/bin
export MINIKUBE_CNI_CONF_DIR=/etc/cni/net.d
export PATH=$PATH:/opt/cni/bin
#kubeadm
curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.32/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.32/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
sudo containerd config default > /etc/containerd/config.toml
下面的镜像似乎没必要
[plugins."io.containerd.grpc.v1.cri"]
  sandbox_image = "registry.cn-hangzhou.aliyuncs.com/google_containers/pause:3.10"
[plugins."io.containerd.grpc.v1.cri".registry]
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors]
    [plugins."io.containerd.grpc.v1.cri".registry.mirrors."docker.io"]
      endpoint = ["https://docker.hoper.xyz"]
    [plugins."io.containerd.grpc.v1.cri".registry.mirrors."k8s.gcr.io"]
      endpoint = ["https://registry.cn-hangzhou.aliyuncs.com/google_containers"]
```/etc/containerd/config.toml
[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc]
  [plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]
    SystemdCgroup = true
```
sudo systemctl restart containerd
sudo usermod -aG docker $USER
sudo usermod -aG containerd $USER
newgrp docker
containerd比较麻烦,还是cri-dockerd吧
sudo rm -r /etc/systemd/system/kubelet.service.d # minikube抢占
sudo systemctl daemon-reload
sudo systemctl restrart kubelet
sudo kubeadm init --config /var/code/hopeio/hoper/awesome/env/k8s/kubeadm-init-config.yaml
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
# 移除所有与 control-plane 相关的污点
kubectl taint nodes pc node-role.kubernetes.io/control-plane-  #pc<node_name>
# 将 NoSchedule 改为 PreferNoSchedule
kubectl taint nodes pc node-role.kubernetes.io/control-plane=:PreferNoSchedule #pc<node_name>
# 添加节点（单节点不用）
kubeadm join 192.168.31.212:6443 --token yesx5d.4setvzq366owe409 \
	--discovery-token-ca-cert-hash sha256:b99ab0ded9dacd03cc6b12996238eb306d3f7f4397d81fa64bd63ccf3c27f784
sudo kubeadm join --config /var/code/hopeio/hoper/awesome/env/k8s/kubeadm-join-config.yaml

# 验证双协议栈
kubectl get nodes pc -o go-template --template='{{range .spec.podCIDRs}}{{printf "%s\n" .}}{{end}}' #pc<node_name>
##验证节点寻址
kubectl get nodes pc -o go-template --template='{{range .status.addresses}}{{printf "%s: %s\n" .type .address}}{{end}}' #pc<node_name>
##验证 Pod 寻址
kubectl get pods -l app=postgres -n tools -o go-template --template='{{range .status.podIPs}}{{printf "%s\n" .ip}}{{end}}' #pc<node_name>
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.29.1/manifests/tigera-operator.yaml
删除/etc/cni/net.d/1-k8s.conflist
sudo cp /var/code/hopeio/hoper/awesome/env/k8s/1-k8s.conflist /etc/cni/net.d/1-k8s.conflist
编辑/etc/crictl.yaml
# 删除节点
kubectl drain <节点名称> --delete-emptydir-data --force --ignore-daemonsets
sudo kubeadm reset --cri-socket unix:///var/run/cri-dockerd.sock
sudo iptables -F && sudo iptables -t nat -F && sudo iptables -t mangle -F && sudo iptables -X
sudo ipvsadm -C
sudo ip link set dev bridge down
sudo ip link delete bridge type bridge
sudo ip link set dev cni0 down
sudo ip link delete cni0 type bridge
# calico
CURL -O -L https://raw.githubusercontent.com/projectcalico/calico/v3.29.1/manifests/tigera-operator.yaml
curl -O -L https://raw.githubusercontent.com/projectcalico/calico/v3.29.1/manifests/custom-resources.yaml

kubectl create -f ~/Documents/tigera-operator.yaml
kubectl apply -f /var/code/hopeio/hoper/awesome/env/k8s/cni-calico.yaml
没装成功 kube-controller跑不起来
2025-01-06 01:55:16.074 [INFO][1] kube-controllers/main.go 92: Loaded configuration from environment config=&config.Config{LogLevel:"info", WorkloadEndpointWorkers:1, ProfileWorkers:1, PolicyWorkers:1, NodeWorkers:1, Kubeconfig:"", DatastoreType:"kubernetes"}
2025-01-06 01:55:16.074 [WARNING][1] kube-controllers/winutils.go 150: Neither --kubeconfig nor --master was specified.  Using the inClusterConfig.  This might not work.
2025-01-06 01:55:16.074 [INFO][1] kube-controllers/main.go 116: Ensuring Calico datastore is initialized
2025-01-06 01:55:46.075 [ERROR][1] kube-controllers/client.go 320: Error getting cluster information config ClusterInformation="default" error=Get "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/clusterinformations/default": dial tcp 10.96.0.1:443: i/o timeout
2025-01-06 01:55:46.075 [INFO][1] kube-controllers/client.go 248: Unable to initialize ClusterInformation error=Get "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/clusterinformations/default": dial tcp 10.96.0.1:443: i/o timeout
2025-01-06 01:56:16.087 [INFO][1] kube-controllers/client.go 254: Unable to initialize default Tier error=Post "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/tiers": dial tcp 10.96.0.1:443: i/o timeout
2025-01-06 01:56:16.088 [INFO][1] kube-controllers/client.go 260: Unable to initialize adminnetworkpolicy Tier error=client rate limiter Wait returned an error: context deadline exceeded
2025-01-06 01:56:16.088 [INFO][1] kube-controllers/main.go 123: Failed to initialize datastore error=Get "https://10.96.0.1:443/apis/crd.projectcalico.org/v1/clusterinformations/default": dial tcp 10.96.0.1:443: i/o timeout
2025-01-06 01:56:16.088 [FATAL][1] kube-controllers/main.go 136: Failed to initialize Calico datastore
反复重装k8s,最后跑起来了，最后全部跑起来要手动删一下coredns的pod
# 神坑！巨坑！无敌坑
先说结论：cri-docker没有支持ipv6双栈
不知道因为改了cni的配置后没重启还是cri-docker的service的ExecStart 加了--ipv6-dual-stack
sudo sytemctl status cri-docker
编辑/etc/systemd/system/cri-docker.service.d/10-cni.conf ExecStart参数加上--ipv6-dual-stack
sudo systemctl daemon-reload
sudo systemctl restart cri-docker
然后pod就有ipv6了
# helm

$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh
# ubuntu
```sh
curl https://baltocdn.com/helm/signing.asc | gpg --dearmor | sudo tee /usr/share/keyrings/helm.gpg > /dev/null
sudo apt-get install apt-transport-https --yes
echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/helm.gpg] https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

# prometheus
kubectl create namespace monitoring
helm install kube-prometheus prometheus-community/kube-prometheus-stack -f helm.yaml -n monitoring
# apisix
cp -r /etc/kubernetes/pki/etcd /root/certs/ && chmod 666 /root/certs/etcd/server.key
kubectl create namespace ingress-apisix
# acme
kubectl apply -f tls.yaml
# tools
kubectl create namespace tools

# 证书过期
sudo kubeadm certs renew all
sudo kubeadm init phase kubeconfig admin # 元宝问的没有这行
sudo cp -f /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# 查看 kubelet 当前使用的证书（需要 root 权限）
sudo openssl x509 -in /var/lib/kubelet/pki/kubelet.crt -text -noout
# 1. 删除旧的 kubelet 证书（自动触发重新生成）
sudo rm -f /var/lib/kubelet/pki/kubelet.crt
sudo rm -f /var/lib/kubelet/pki/kubelet.key

# 查找并重启相关容器
systemctl restart kubelet
docker restart $(docker ps | grep kube-apiserver | awk '{print $1}')
docker restart $(docker ps | grep kube-controller-manager | awk '{print $1}')
docker restart $(docker ps | grep kube-scheduler | awk '{print $1}')
## 自动续期
sudo vim /var/lib/kubelet/config.yaml
rotateCertificates: true
serverTLSBootstrap: true
