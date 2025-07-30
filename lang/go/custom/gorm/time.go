package main

import (
	"github.com/hopeio/initialize"
	"github.com/hopeio/gox/log"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	defer initialize.Start(confdao.Config, confdao.Dao)()
	var tests []*model.Model
	log.Info(confdao.Dao.DB3111.Table("test").Find(&tests))
}
