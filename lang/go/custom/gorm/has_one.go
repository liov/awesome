package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
	"test/custom/gorm/model"
)

func main() {

	db, _ := gorm.Open(tests.DummyDialector{}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	var (
		abs []*model.AB
	)
	tx := db.Joins("Node").Where(`id = ?`, 1)
	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&abs)
	})
	fmt.Println(sql)

}
