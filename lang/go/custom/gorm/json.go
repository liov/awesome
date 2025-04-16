package main

import (
	"fmt"
	"github.com/hopeio/initialize"
	"github.com/hopeio/utils/dao/database/datatypes"
	"gorm.io/gorm"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	e := model.ModelJson{
		Json: datatypes.JsonT[model.Tag]{
			Data: &model.Tag{
				Model: gorm.Model{
					ID: 1,
				},
			},
		},
	}
	sql := confdao.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&e)
	})
	fmt.Println(sql)
	defer initialize.Start(confdao.Config, confdao.Dao)()
	confdao.Dao.DB2111.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&e)
	})
}
