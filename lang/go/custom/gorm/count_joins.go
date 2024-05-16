package main

import (
	"github.com/hopeio/cherry/utils/log"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

type AB struct {
	Id  int `json:"id" gorm:"primaryKey"`
	BId int `json:"aId" gorm:"index"`
	AId int `json:"bId"`
	B   B   `json:"b"`
	A   A   `json:"a"`
}

type A struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

type B struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

func main() {
	db, _ := gorm.Open(tests.DummyDialector{}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	var (
		count    int64
		appNodes = []*AB{}
	)
	tx := db.Model(&AB{}).Preload("A").Joins("B").Where(`id = ?`, 11)

	db.Count(&count)

	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&appNodes)
	})
	log.Info(sql)
}
