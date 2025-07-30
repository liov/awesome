package errcode

import "github.com/hopeio/gox/errors/errcode"

const (
	DBError errcode.ErrCode = 21000
)

func init() {
	errcode.Register(DBError, "数据库错误")
}
