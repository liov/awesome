package main

import (
	"encoding/json"
	"fmt"
	"test/custom/gorm/confdao"
)

func main() {
	m := make(map[string]interface{})
	err := confdao.Dao.DB3111.Raw(`SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public' 
	`).Scan(m).Error
	if err != nil {
		panic(err)
	}
	data, _ := json.Marshal(m)
	fmt.Println(string(data))
	m = make(map[string]interface{})
	err = confdao.Dao.DB3111.Raw(`SELECT json_agg(row_to_json(t)) FROM (SELECT * FROM tbl_33 LIMIT 10) t`).Scan(m).Error
	if err != nil {
		panic(err)
	}
	data, _ = json.Marshal(m)
	fmt.Println(string(data))
}
