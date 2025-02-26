package confdao

import (
	"github.com/hopeio/initialize"
	postgresi "github.com/hopeio/initialize/dao/gormdb/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
)

type config struct {
	initialize.EmbeddedPresets
}

type dao struct {
	initialize.EmbeddedPresets
	DB postgresi.DB
}

var Dao = &dao{}
var Config = &config{}

// var Global = initialize.NewGlobal[*config, *dao]()
var DB, _ = gorm.Open(tests.DummyDialector{}, &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
})
