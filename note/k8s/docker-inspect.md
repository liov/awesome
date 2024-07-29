使用 docker inspect 命令查看容器的详细信息：

docker inspect <container_id>

# ip
在输出的 JSON 信息中，找到 NetworkSettings 属性下的 Networks 属性。在该属性下，您可以找到与容器关联的每个网络的详细信息。在每个网络对象中，查找 IPAddress 属性，即为该容器在该网络中的 IP 地址

# network
在Docker中，--network参数用于设置容器的网络模式。以下是一些常见的网络模式及其设置方法：

1. Bridge模式（默认模式）
   这是Docker的默认网络模式。在这种模式下，每个容器都有自己的网络命名空间，并通过Docker的桥接网络进行通信。

docker run --network bridge my_image
2. Host模式
   在这种模式下，容器共享主机的网络命名空间，因此容器可以直接访问主机的网络接口。

docker run --network host my_image
3. None模式
   在这种模式下，容器没有任何网络配置，无法访问外部网络。

docker run --network none my_image
4. Overlay模式
   这种模式主要用于多主机环境，通过Docker Swarm或Kubernetes等工具进行容器编排。

docker run --network overlay my_image
5. Macvlan模式
   这种模式允许容器直接连接到物理网络，每个容器都有自己的MAC地址。

docker run --network macvlan --mac-address=02:42:ac:11:00:02 my_image
6. IPvlan模式
   与Macvlan类似，但IPvlan模式更轻量级，适用于高性能网络应用。

docker run --network ipvlan --ip 192.168.1.10 my_image
7. 自定义网络
   你可以创建自定义网络，并将容器连接到这些网络。

## 创建自定义网络
docker network create my_network

## 运行容器并连接到自定义网络
docker run --network my_network my_image
示例
假设你有一个名为my_image的镜像，并且你想将其运行在自定义网络my_network中，可以使用以下命令：


## 创建自定义网络
docker network create my_network
## 运行容器并连接到自定义网络
docker run --network my_network my_image