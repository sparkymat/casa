package model

import "gorm.io/gorm"

type HomeItem struct {
	gorm.Model
	UserID        uint
	User          User
	CatalogItemID uint
	CatalogItem   CatalogItem
}
