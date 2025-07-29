package model

import "github.com/hopeio/utils/datax/database/datatypes"

type TestJson struct {
	ID        uint                      `json:"id" gorm:"primaryKey"`
	JsonArray datatypes.NullJson[[]Foo] `json:"json_array" gorm:"jsonb"`
}

type Foo struct {
	A int
	B string
}
