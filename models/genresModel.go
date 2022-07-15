package models

import "time"

type Genres struct {
	Id          uint
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt 	time.Time
	DeletedAt 	time.Time
}