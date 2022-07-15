package models

import "encoding/json"

type GenresType struct {
	Title string `json:"title" binding:"required"`
	// int number equal example 20
	// Fee int `json:"fee" binding:"required,number"`

	// json_number as string "20" same 20 not same "text here"
	Fee json.Number `json:"fee" binding:"required,number"`
}