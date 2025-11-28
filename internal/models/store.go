package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name    string `json:"name"`
	Author  string `json:"author"`
	Price   int    `json:"price"`
	FileURL string `json:"fileURL"`
}
