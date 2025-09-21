package main

import (
	"encoding/json"
	"net/http"
	"os"
)

func main() {
	type NacosRes struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	}
	res, _ := http.Get("http://xxx:8848/nacos/v2/cs/config?dataId=room.toml&group=pg&namespaceId=pg")
	defer res.Body.Close()
	var nacosRes NacosRes
	json.NewDecoder(res.Body).Decode(&nacosRes)
	f, _ := os.OpenFile("room.toml", os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	f.WriteString(nacosRes.Data)

}
