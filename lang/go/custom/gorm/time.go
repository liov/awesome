package main

import (
	"github.com/hopeio/cherry/initialize"
	"github.com/hopeio/cherry/utils/log"
	"test/custom/gorm/confdao"
)

func main() {
	defer initialize.Start(confdao.Config, confdao.Dao)()
	var tests []*Test
	log.Info(confdao.Dao.DB.Table("test_json").Find(&tests))
}
