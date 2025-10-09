package errcode

import (
	"github.com/hopeio/gox/errors"
)

const (
	DBError errors.ErrCode = 21000
)

func init() {
	errors.Register(DBError, "数据库错误")
}
