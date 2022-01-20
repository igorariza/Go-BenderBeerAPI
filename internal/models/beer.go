package models

import "time"

type Beer struct {
	Base
	Name string `json:"name" bson:"name" gorm:"index"`
}

func (Beer) TableName() string {
	return "vault_secret"
}

type BeerModel struct {
	Id        uint32 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"type:varchar(255);unique;not null"`
	Brewery   string `gorm:"type:varchar(255);unique;not null"`
	Country   string `gorm:"type:varchar(255);unique;not null"`
	Price     int32  `gorm:"type:int;unique;not null"`
	Currency  int32  `gorm:"type:int;unique;not null"`
}
