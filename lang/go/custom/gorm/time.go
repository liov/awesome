package main

import (
	"github.com/hopeio/initialize"
	"github.com/hopeio/utils/log"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	defer initialize.Start(confdao.Config, confdao.Dao)()
	var tests []*model.Model
	log.Info(confdao.Dao.DB.Table("test").Find(&tests))
}
