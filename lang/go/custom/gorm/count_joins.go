package main

import (
	"github.com/hopeio/utils/log"
	"gorm.io/gorm"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	var (
		count int64
		abs   = []*model.AB{}
	)
	tx := confdao.DB.Model(&model.AB{}).Preload("A").Joins("B").Where(`id = ?`, 11)

	confdao.DB.Count(&count)

	sql := tx.ToSQL(func(db *gorm.DB) *gorm.DB {
		return db.Limit(1).
			Offset(0).
			Order("id DESC").
			Find(&abs)
	})
	log.Info(sql)
}
