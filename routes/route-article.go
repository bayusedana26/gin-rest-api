package routes

import (
	"github.com/bayusedana26/gin-rest-api.git/config"
	"github.com/bayusedana26/gin-rest-api.git/models"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetHome(ctx *gin.Context) {
	items := []models.Article{}

	config.DB.Find(&items)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    items,
	})
}

func GetArticle(ctx *gin.Context) {
	slug := ctx.Param("slug")
	var item models.Article

	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
		ctx.JSON(400, gin.H{
			"message": "Record not found",
		})
		ctx.Abort()
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    item,
	})
}

func PostArticle(ctx *gin.Context) {
	item := models.Article{
		Title: ctx.PostForm("title"),
		Desc:  ctx.PostForm("desc"),
		Slug:  slug.Make(ctx.PostForm("title")),
	}
	config.DB.Create(&item)

	ctx.JSON(200, gin.H{
		"message": "Success",
		"data":    item,
	})
}
