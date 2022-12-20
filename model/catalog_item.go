package model

import "gorm.io/gorm"

type CatalogItem struct {
	gorm.Model
	Title       string
	Description string
	ImageURL    string
	URL         string
}
