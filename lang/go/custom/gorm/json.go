package main

import (
	"github.com/hopeio/gox/log"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	//confdao.Dao.DB3111.Migrator().CreateTable(&model.Test{})
	defer confdao.Global.Cleanup()
	//confdao.Dao.DB3111.Create(&model.Test{
	//		V: datatypes.JsonT[model.Tag]{
	//			V: model.Tag{
	//				Model: gorm.Model{
	//					ID: 1,
	//				},
	//			},
	//		},
	//	})
	var e2 []model.Test
	confdao.Dao.DB3111.Find(&e2)
	log.Info(e2)
}
