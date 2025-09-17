package main

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/consul/api"
)

func registerService() {
	// 创建 Consul 客户端配置
	config := api.DefaultConfig()
	config.Address = "127.0.0.1:8500" // Consul 地址

	// 创建客户端
	client, err := api.NewClient(config)
	if err != nil {
		log.Fatal("Consul client error: ", err)
	}

	// 服务注册信息
	registration := &api.AgentServiceRegistration{
		ID:      "my-service-1",       // 服务唯一ID
		Name:    "my-service",         // 服务名称
		Port:    8080,                 // 服务端口
		Address: "127.0.0.1",          // 服务地址
		Tags:    []string{"v1", "go"}, // 标签，可用于过滤

		// 健康检查配置
		Check: &api.AgentServiceCheck{
			HTTP:     "http://127.0.0.1:8080/health", // 健康检查地址
			Interval: "10s",                          // 检查间隔
			Timeout:  "5s",                           // 检查超时
		},
	}

	// 注册服务
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("Register service error: ", err)
	}

	fmt.Println("Service registered successfully!")
}

func main() {
	registerService()

	// 保持程序运行
	for {
		time.Sleep(1 * time.Second)
	}
}
