package config

import (
	"github.com/bayusedana26/gin-rest-api.git/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB() {
	var err error

	DB, err = gorm.Open("mysql", "root:@tcp(localhost:3306)/gin_gorm?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	// Migrate
	DB.AutoMigrate(&models.Article{})
}
