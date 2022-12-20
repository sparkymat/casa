package model

import "gorm.io/gorm"

type CatalogItem struct {
	gorm.Model
	Title       string
	Description string
	ImageUrl    string
	Url         string
}
