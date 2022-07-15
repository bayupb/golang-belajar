package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	// fungsi route default gin github
	r := gin.Default()

	// TestHandler push query 
	r.GET("/" ,  TestHandler)
	r.GET("/genres/:id" ,  GenresHandler)
	r.GET("/genres/query/" ,  GenresQueryHandler)
	
	// post
	r.POST("/genres/simpan" , PostGenresHandler)

	// active port in localhost:2000
	r.Run(":2000")
}

func TestHandler(fungsi *gin.Context){
	fungsi.JSON(http.StatusOK, gin.H{
		"user_id" : 1,
		"name" : "Bayu Priyambabada",
		"bio" : "Junior Software Development",
	})
}

func GenresHandler(fungsi *gin.Context){
	id := fungsi.Param("id")

	fungsi.JSON(http.StatusOK, gin.H{"id": id})
}

func GenresQueryHandler(fungsi *gin.Context){
	
	// single params query example query?title=test
	title := fungsi.Query("title")
	// double params query example query?title=test&description=query
	description := fungsi.Query("description")

	fungsi.JSON(http.StatusOK, gin.H{"title": title , "description" : description})
}

// queries PostGenresHandler with type

// validator golang required in fields column json
type GenresType struct{
	Title string `json:"title" binding:"required"`
	// int number equal example 20
	// Fee int `json:"fee" binding:"required,number"`

	// json_number as string "20" same 20 not same "text here"
	Fee json.Number `json:"fee" binding:"required,number"`
}

func PostGenresHandler(fungsi *gin.Context){
	var GenresType GenresType

	// func should on Json not use Form-Data
	errors := fungsi.ShouldBindJSON(&GenresType)

	// if queries error or null check data [1].
	// if errors != nil {
	// 	fungsi.JSON(http.StatusBadRequest, errors)
	// 	return
	// }

	if errors != nil{

		// message error many validate
		errorMessagess := []string{}
		for _, error := range errors.(validator.ValidationErrors){
				errorMessages := fmt.Sprintf(" Column %s cannot be empty", error.Field())
				errorMessagess = append(errorMessagess, errorMessages)
			}
 
		fungsi.JSON(http.StatusBadRequest, gin.H {
			"errors" : errorMessagess,
		})
		return
		// if make an errors validation show 1 data validate
		// for _, error := range errors.(validator.ValidationErrors){
		// 	errorMessage :=  fmt.Sprintf("Error fields column %s, cannot be empty", error.Field())
			// fungsi.JSON(http.StatusBadRequest, errorMessage)
			// return
		// }
	}
	// queries post data with data title,description,fee
	fungsi.JSON(http.StatusOK, gin.H{
		"title": GenresType.Title, 
		"fee": GenresType.Fee,
	})
}

