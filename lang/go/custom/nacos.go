package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/hopeio/gox/log"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

func main() {
	grpc()

}

func http1() {
	type NacosRes struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
	res, _ := http.Get("http://xxx:8848/nacos/v2/cs/config?dataId=xx&group=xx&namespaceId=xx")
	defer res.Body.Close()
	var nacosRes NacosRes
	json.NewDecoder(res.Body).Decode(&nacosRes)
	f, _ := os.OpenFile("xx", os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	f.WriteString(nacosRes.Data)
}

func grpc() {
	// 1. 初始化 Nacos 客户端配置
	clientConfig := constant.ClientConfig{
		NamespaceId:         "xx", // 命名空间ID，默认是 public
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		Username:            "nacos", // 如果开启鉴权
		Password:            "nacos",
	}

	// 2. 配置 Nacos 服务器地址
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "8.219.80.188", // Nacos 服务器IP
			Port:        8848,           // Nacos 服务器端口
			ContextPath: "/nacos",       // 默认路径
		},
	}

	// 3. 创建配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		log.Fatal("创建 Nacos 客户端失败:", err)
	}

	// 4. 获取配置
	dataId := "xx" // 配置ID
	group := "pg"  // 配置分组（默认分组）
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})
	if err != nil {
		log.Fatal("获取配置失败:", err)
	}

	fmt.Println("获取到的配置内容:", content)
}
