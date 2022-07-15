package main

import (
	"golang1/models"
	"golang1/repositories"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// connectiion db
	dsn := "root:@tcp(127.0.0.1:3306)/golang-belajar?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db connection failed")
	}

	db.AutoMigrate(&models.Genres{})

	// fungsi route default gin github
	r := gin.Default()

	v1 := r.Group("/v1")

	// TestHandler push query
	v1.GET("/", TestHandler)
	v1.GET("/genres/:id", repositories.GenresHandler)
	v1.GET("/genres/query/", repositories.GenresQueryHandler)

	// post
	r.POST("/genres/simpan", repositories.PostGenresHandler)

	// active port in localhost:2000
	r.Run(":2000")
}

func TestHandler(fungsi *gin.Context) {
	fungsi.JSON(http.StatusOK, gin.H{
		"user_id": 1,
		"name":    "Bayu Priyambabada",
		"bio":     "Junior Software Development",
	})
}

// Genres move to GenresRepositories

// queries PostGenresHandler with type

// validator golang required in fields column json
// genres Struct Move to models/genresModels
