package main

import (
	"fmt"
	"gorm.io/gorm"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {

	var (
		abs []*model.AB
	)
	tx := confdao.DB.Joins("Node").Where(`id = ?`, 1)
	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&abs)
	})
	fmt.Println(sql)

}
