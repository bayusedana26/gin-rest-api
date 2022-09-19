package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title string
	Desc  string `sql:"type:text;"`
	Slug  string `gorm:"unique_index"`
}
