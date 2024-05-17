package main

import (
	"gorm.io/gorm"
	"log"
	"test/custom/gorm/model"
)

func Select(db *gorm.DB) {
	var models []*model.Model
	err := db.Table(model.ModelTable).Find(&models).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, model := range models {
		log.Println(model)
	}
}

func SelectStrTime(db *gorm.DB) {
	var models []*model.ModelA
	err := db.Table(model.ModelTable).Find(&models).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, model := range models {
		log.Println(model.K)
	}
}
