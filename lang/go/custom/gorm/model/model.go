package model

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id        int
	A         int
	B         int
	C         int
	D         int
	E         int
	F         string
	H         string
	I         string
	J         string
	K         time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type ModelA struct {
	Id        int
	A         int
	B         int
	C         int
	D         int
	E         int
	F         string
	H         string
	I         string
	J         string
	K         string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

const ModelTable = "model"

type Test struct {
	Id        int
	DeletedAt sql.NullTime
}

type AB struct {
	Id  int `json:"id" gorm:"primaryKey"`
	BId int `json:"aId" `
	AId int `json:"bId"`
	B   B   `json:"b"`
	A   A   `json:"a"`
}

type A struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name" g`
}

type B struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	AId  int    `json:"aId"`
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	Account Account
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string

	Companies []Company
	Pet       Pet
}

type Company struct {
	ID        int
	AccountID int32
	Name      string
}

type Pet struct {
	gorm.Model

	AccountID *uint
	Name      string
}
