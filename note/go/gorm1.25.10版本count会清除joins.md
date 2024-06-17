从2024年1月29开始
https://github.com/go-gorm/gorm/issues/6715
https://github.com/go-gorm/gorm/pull/6771
https://github.com/go-gorm/gorm/issues/7025
```go
type AB struct {
	Id  int `json:"id" gorm:"primaryKey"`
	BId int `json:"aId" gorm:"index"`
	AId int `json:"bId"`
	B   B   `json:"b"`
	A   A   `json:"a"`
}

type A struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

type B struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"uniqueIndex"`
}

func main() {
	db, _ := gorm.Open(tests.DummyDialector{}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	var (
		count    int64
		appNodes = []*AB{}
	)
	db:= db.Model(&AB{}).Joins("A").Joins("B").Where(`id = ?`, 1)

	db.Count(&count)
	db.Limit(1).Offset(0).Order("id DESC").Find(&appNodes)
}
```
与 `db := db.Model(&AB{}).Preload("A").Preload("B").Where(`id = ?`, 1)`结果不一致