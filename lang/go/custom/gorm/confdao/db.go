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
	DB2155 postgresi.DB
}

var Global = initialize.NewGlobal[*config, *dao]()

var Dao = Global.Dao
var Config = Global.Config

// var Global = initialize.NewGlobal[*config, *dao]()
var DB, _ = gorm.Open(tests.DummyDialector{}, &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	},
})
