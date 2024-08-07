package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/gorm/utils/tests"
	"sync"
	"test/custom/gorm/confdao"
)

func main() {

	user, _ := schema.Parse(&tests.User{}, &sync.Map{}, confdao.DB.NamingStrategy)
	stmt := gorm.Statement{DB: confdao.DB, Table: user.Table, Schema: user, Clauses: map[string]clause.Clause{}}
	clauses := []clause.Interface{clause.Select{}, clause.From{}, clause.Where{Exprs: []clause.Expression{clause.Eq{Column: clause.PrimaryColumn, Value: "1"}, clause.Gt{Column: "age", Value: 18}, clause.Or(clause.Neq{Column: "name", Value: "jinzhu"}, clause.Neq{Column: "id", Value: "1"})}}}

	for _, clause := range clauses {
		stmt.AddClause(clause)
	}

	stmt.Build("SELECT", "FROM", "WHERE")
	fmt.Println(stmt.SQL.String())
}
