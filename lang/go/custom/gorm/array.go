package main

import (
	"fmt"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"

	_ "github.com/hopeio/gox/database/sql/gorm/serializer"
	"gorm.io/gorm"
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
	defer confdao.Global.Cleanup()
	confdao.Dao.DB2155.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return tx.Create(&e)
	})
}
