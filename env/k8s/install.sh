# docker
../docker/install.md
# k8s
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube

minikube start --driver=none  --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers --extra-config=apiserver.service-node-port-range=1-39999  --extra-config=kube-proxy.mode=ipvs --extra-config=kubelet.authentication-token-webhook=true --extra-config=kubelet.authorization-mode=Webhook --extra-config=scheduler.bind-address=0.0.0.0 --extra-config=controller-manager.bind-address=0.0.0.0 --bootstrapper=kubeadm

cp -r /var/lib/minikube/certs/etcd /root/certs/ && chmod 666 /root/certs/etcd/server.key
minikube addons enable dashboard
minikube addons enable logviewer
minikube addons enable efk
minikube addons enable helm-tiller

# prometheus
kubectl create namespace monitoring
helm install kube-prometheus prometheus-community/kube-prometheus-stack -f helm.yaml -n monitoring
# apisix
kubectl create namespace ingress-apisix
# acme
kubectl apply -f tls.yaml
# tools
kubectl create namespace tools

