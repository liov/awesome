package main

import (
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	confdao.DB.Delete(&model.AB{}, 10)
}
