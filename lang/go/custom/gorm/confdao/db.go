package confdao

import (
	"github.com/hopeio/cherry/initialize"
	"github.com/hopeio/cherry/initialize/conf_dao/gormdb/postgres"
)

type config struct {
	initialize.EmbeddedPresets
}

type dao struct {
	initialize.EmbeddedPresets
	DB postgres.DB
}

var Dao = &dao{}
var Config = &config{}
