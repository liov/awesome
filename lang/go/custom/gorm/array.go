package main

import (
	"fmt"
	"github.com/hopeio/initialize"
	_ "github.com/hopeio/utils/dao/database/gorm/serializer"
	"gorm.io/gorm"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	e := model.ModelArray{
		Array:  []string{"1", "2"},
		Array2: []string{"2", "3"},
	}
	sql := confdao.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&e)
	})
	fmt.Println(sql)
	global := initialize.NewGlobal[confdao.Config, confdao.Dao]()
	defer global.Cleanup()
	global.Dao.DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&e)
	})
}
