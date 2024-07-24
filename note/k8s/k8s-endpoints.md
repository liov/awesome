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

以下是如何在 Docker 容器中使用 "NetworkMode": "container:<container_id>" 的示例：