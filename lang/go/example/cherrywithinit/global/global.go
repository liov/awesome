package global

import (
	"database/sql"
	"fmt"
	"github.com/hopeio/cherry"
	"github.com/hopeio/initialize/dao/gormdb/sqlite"
	"github.com/hopeio/utils/os/fs"
	timei "github.com/hopeio/utils/time"
	"runtime"
	"time"
)

var (
	Conf      = &config{}
	Dao  *dao = &dao{}
)

type config struct {
	//自定义的配置
	Customize serverConfig
	Server    cherry.Server
}

type serverConfig struct {
	Volume fs.Dir

	PassSalt string
	// 天数
	TokenMaxAge time.Duration
	TokenSecret string
	PageSize    int8

	LuosimaoSuperPW   string
	LuosimaoVerifyURL string
	LuosimaoAPIKey    string

	QrCodeSaveDir fs.Dir //二维码保存路径
	PrefixUrl     string
	FontSaveDir   fs.Dir //字体保存路径

}

func (c *config) BeforeInject() {
	c.Customize.TokenMaxAge = timei.Day
}

func (c *config) AfterInject() {
	if runtime.GOOS == "windows" {
	}

	c.Customize.TokenMaxAge = timei.StdDuration(c.Customize.TokenMaxAge, time.Hour)
}

// dao dao.
type dao struct {
	// GORMDB 数据库连接
	GORMDB sqlite.DB
	StdDB  *sql.DB
}

func (d *dao) BeforeInject() {
	d.GORMDB.Conf.Gorm.NowFunc = time.Now
}

func (d *dao) AfterInjectConfig() {
	fmt.Println("这里后执行")
}

func (d *dao) AfterInject() {
	db := d.GORMDB
	db.Callback().Create().Remove("gorm:save_before_associations")
	db.Callback().Create().Remove("gorm:save_after_associations")
	db.Callback().Update().Remove("gorm:save_before_associations")
	db.Callback().Update().Remove("gorm:save_after_associations")

	d.StdDB, _ = db.DB.DB()
}
