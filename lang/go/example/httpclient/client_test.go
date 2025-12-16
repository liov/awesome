package client

import (
	"crypto/tls"
	"net/http"
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	httpx "github.com/hopeio/gox/net/http"
	"github.com/hopeio/gox/net/http/client"
	clientv2 "github.com/hopeio/gox/net/http/client/v2"
)

func TestUserList(t *testing.T) {
	var res httpx.RespData[UserListRes]
	client := client.New().AddHeader("Content-Type", "application/json").LogLevel(client.LogLevelInfo)
	err := client.Request("GET", "http://localhost:8080/api/user").Do(&Page{1, 2}, &res)
	if err != nil {
		t.Log(err)
	}
	spew.Dump(res)
}

func TestUserListV2(t *testing.T) {
	res, err := clientv2.NewRequest[httpx.RespData[UserListRes]]("GET", "http://localhost:8080/api/user").Do(&Page{1, 2})
	if err != nil {
		t.Log(err)
	}

	spew.Dump(res)
}

func TestU(t *testing.T) {
	data, _ := os.ReadFile("D:/work/debug_infer.json")
	var resp client.RawBytes
	err := client.New().HttpClient(&http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // 禁用证书验证
			},
		},
	}).Post("https://xxx", data, &resp)
	if err != nil {
		t.Log(err)
	}
	t.Log(string(resp))
}
