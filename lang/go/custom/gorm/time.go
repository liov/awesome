package main

import (
	"database/sql"
	"github.com/hopeio/cherry/initialize"
	"github.com/hopeio/cherry/utils/log"
	"test/custom/gorm/confdao"
)

type Test struct {
	Id        int
	DeletedAt sql.NullTime
}

func main() {
	defer initialize.Start(confdao.Config, confdao.Dao)()
	var tests []*Test
	log.Info(confdao.Dao.DB.Table("test_json").Find(&tests))
}
