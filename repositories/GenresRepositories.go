package repositories

import (
	"fmt"
	"golang1/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GenresHandler(fungsi *gin.Context) {
	id := fungsi.Param("id")

	fungsi.JSON(http.StatusOK, gin.H{"id": id})
}

func GenresQueryHandler(fungsi *gin.Context) {

	// single params query example query?title=test
	title := fungsi.Query("title")
	// double params query example query?title=test&description=query
	description := fungsi.Query("description")

	fungsi.JSON(http.StatusOK, gin.H{"title": title, "description": description})
}

func PostGenresHandler(fungsi *gin.Context){
	var GenresType models.GenresType

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