package model

import "gorm.io/gorm"

type Car struct {
	Id    int    `gorm:"type:int;primary_key"`
	Name  string `gorm:"type:varchar(50)"`
	Color string `gorm:"type:varchar(50)"`
	*gorm.Model
}
