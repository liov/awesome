package main

import (
	"gorm.io/gorm"
	"log"
	"test/custom/gorm/model"
)

func main() {

}

func Scan(db *gorm.DB) {
	var id int
	err := db.Table(model.ModelTable).Select("MAX(id)").Scan(&id).Error
	if err != nil {
		log.Println(err)
	}
	log.Println(id)
}

func RawScan(db *gorm.DB) {
	var exists bool
	err := db.Raw(`SELECT EXISTS(SELECT * FROM tsp_info WHERE id = 59)`).Scan(&exists).Error
	if err != nil {
		log.Println(err)
	}
	log.Println(exists)
}
