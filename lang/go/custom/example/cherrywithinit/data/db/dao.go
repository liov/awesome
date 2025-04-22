package db

import (
	"test/example/cherrywithinit/model"

	"github.com/hopeio/context/httpctx"
	"github.com/hopeio/utils/log"
	"gorm.io/gorm"
)

type userDao struct {
	*httpctx.Context
	db *gorm.DB
}

func GetDao(ctx *httpctx.Context, db *gorm.DB) *userDao {
	if ctx == nil {
		log.Fatal("ctx can't nil")
	}
	return &userDao{ctx, db}
}

func (d *userDao) GetJsonArrayT(id int) (*model.TestJson, error) {
	var t model.TestJson
	err := d.db.First(&t, id).Error
	if err != nil {
		return nil, err
	}
	return &t, nil
}
