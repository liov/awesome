package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

type A struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Status string `json:"status"`
	B      B      `json:"B"`
}

type B struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	AID  int    `json:"a_id"`
	Name string `json:"name"`
}

func main() {

	db, _ := gorm.Open(tests.DummyDialector{}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	var (
		appNodes []*AppNode
	)
	tx := db.Model(&AppNode{}).Joins("Node").Where(`app_version_id = ?`, 10)
	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&appNodes)
	})
	fmt.Println(sql)

}
