package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/hopeio/initialize"
	"github.com/hopeio/utils/log"
	"test/custom/gorm/confdao"
	"test/custom/gorm/model"
)

func main() {
	defer initialize.Start(confdao.Config, confdao.Dao)()
	DB := confdao.Dao.DB2111.DB
	user := model.User{
		Account: model.Account{
			Number: "123456",
			Companies: []model.Company{
				{Name: "Corp1"}, {Name: "Corp2"},
			},
			Pet: model.Pet{
				Name: "Pet1",
			},
		},
	}
	DB.Migrator().DropTable(&model.User{}, &model.Account{}, &model.Pet{}, &model.Company{})
	DB.AutoMigrate(&model.User{}, &model.Account{}, &model.Pet{}, &model.Company{})
	DB.Create(&user)
	fmt.Println("-------------------------------------------------------")
	var count int64
	var result model.User
	DB = DB.Model(&model.User{}).
		Joins("Account").
		Joins("Account.Pet").
		Preload("Account.Companies")

	if err := DB.Count(&count).Error; err != nil {
		log.Errorf("Failed, got error: %v", err)
	}
	DB.Clauses()

	if err := DB.First(&result, user.ID).Error; err != nil {
		log.Errorf("Failed, got error: %v", err)
	}

	if len(result.Account.Companies) != 2 {
		log.Errorf("Failed, got %v", len(result.Account.Companies))
	}

	if result.Account.Pet.Name != "Pet1" {
		log.Errorf("Failed, got '%v'", result.Account.Pet.Name)
	}
	log.Info(count)
	spew.Dump(result)
}
