在 Kubernetes (k8s) 中，Service 和 Endpoint 是两个核心概念，它们共同工作以允许外部访问集群内的应用程序

1. **Service**：Service 是一种抽象，它定义了访问和暴露集群内 Pod 的一种方式。Service 使得您可以不必关心具体的 Pod IP 地址，就能通过一个稳定的 IP 地址和端口访问服务。Service 通过标签（Label）来选择后端的一组 Pod，并路由流量到这些 Pod。

2. **Endpoints**：Endpoints 是 Service 的后端实现。它是一个包含 IP 地址和端口的对象，代表了能够提供服务的一组 Pod。当 Service 被创建或更新时，Endpoints 会自动根据 Service 的选择器（Selector）来发现并添加相应的 Pod。同样，当 Pod 被创建、更新或删除时，Endpoints 也会自动更新。

Service 和 Endpoint 之间的关系可以简化为：Service 是一个前端接口，它定义了如何访问一组具有相同功能的 Pod；而 Endpoint 是这个前端接口的后端实现，它包含了实际提供服务的一组 Pod 的 IP 地址和端口。

要查看特定 Service 的 Endpoint，您可以使用 `kubectl` 命令行工具：

```bash
kubectl get endpoints <service-name> -n <namespace>
```

这将显示特定命名空间中指定 Service 的 Endpoint 信息，包括 IP 地址和端口。

Kubernetes 通常使用自己的网络插件（如 Calico、Flannel 或 Cilium）来管理集群内的网络通信。这意味着 Kubernetes 中的容器可能不会直接连接到 Docker 的默认 bridge 网络

在 Docker 中，"NetworkMode": "container:<container_id>" 表示一个容器将共享另一个容器的网络命名空间

这种网络模式允许一个容器直接访问另一个容器的网络环境，就像它们在同一个网络命名空间中一样。这对于需要紧密协作的容器（例如，一个容器需要访问另一个容器的内部服务）非常有用。



```yaml
apiVersion: v1
kind: Pod
metadata:
  name: no-network-container
spec:
  containers:
  - name: no-network
    image: my-image
    networkMode: "none"
```

# ClusterIP和Endpoint的区别
在Kubernetes（k8s）中，ClusterIP和Endpoint是两个核心概念，它们在服务发现和负载均衡中扮演着不同的角色。以下是它们之间的主要区别：

ClusterIP
定义：ClusterIP是Kubernetes为Service分配的一个虚拟IP地址，它只能在集群内部访问。这个IP地址用于在集群内部路由流量到与服务匹配的Pod 。
作用：ClusterIP的主要作用是提供内部负载均衡，使得集群内的服务可以通过一个固定的IP地址访问，即使这些Pod的数量和位置在不断变化 。
使用场景：适用于只需要在Kubernetes集群内部可访问的服务，如数据库连接、队列服务等。它也可以作为其他类型的Service（如NodePort、LoadBalancer或Ingress）的后端服务 。
Endpoint
定义：Endpoint是Kubernetes中的一个资源对象，用于记录一个Service对应的所有Pod的访问地址。它包含了Pod的IP地址和端口信息，是Service实现实际服务的端点的集合 。
作用：Endpoint使得Service能够找到后端的Pod，即使这些Pod是动态变化的。它配合Service使用，可以实现负载均衡，确保Service总是指向正确的Pod 。
使用场景：Endpoint通过自动更新来反映Service的selector或Pod的标签变化，确保Service的高可用性和可扩展性。它对Service的使用者是透明的，只需要使用Service的名称和端口号即可访问到后端Pod 。
总的来说，ClusterIP是Kubernetes为Service提供的一个虚拟IP地址，用于内部负载均衡和通信，而Endpoint是记录Service对应的所有Pod的访问地址的资源对象，它确保了Service能够动态地找到和负载均衡到正确的Pod。两者共同工作，实现了Kubernetes中服务的高可用性和可扩展性。

# kube-ipvs0网卡
kubernetes将kube-proxy的设置为ipvs模式后会在每个创建一个ipvs0的网卡，并会在每个节点的ipvs0网卡上配置所有service的ip，等于每个节点都配置了很多相同的ip，为什么不会出现ip冲突？集群内部访问service ip如何访问？集群外部访问service ip又是如何访问？

ipvs0网卡是一块 dummy 类型的虚拟网卡，可以发现ipvs0网卡有一个NOARP的标志，表示的是禁用arp，

# 实测，ipvs模式，calico插件，无论ClusterIP和Endpoint都可以直接访问