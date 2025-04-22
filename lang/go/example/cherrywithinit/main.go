package main

import (
	"test/example/cherrywithinit/api"
	"test/example/cherrywithinit/global"

	"github.com/hopeio/cherry"

	"github.com/hopeio/initialize"
	"github.com/hopeio/initialize/conf_center/nacos"
)

//go:generate protogen.exe go -e -w -v -p proto -o proto
func main() {
	defer initialize.Start(global.Conf, global.Dao, nacos.ConfigCenter)()
	global.Conf.Server.WithOptions(cherry.WithGrpcHandler(api.GrpcRegister), cherry.WithGinHandler(api.GinRegister)).Run()
}
