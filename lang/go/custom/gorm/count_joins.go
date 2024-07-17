package main

import (
	"github.com/hopeio/utils/log"
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
		count int64
		abs   = []*model.AB{}
	)
	tx := db.Model(&model.AB{}).Preload("A").Joins("B").Where(`id = ?`, 11)

	db.Count(&count)

	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&abs)
	})
	log.Info(sql)
}
