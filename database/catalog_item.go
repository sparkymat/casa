package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

func (s *Service) CreateCatalogItem(_ context.Context, title string, description string, imageUrl string) (*model.CatalogItem, error) {
	item := model.CatalogItem{
		Title:       title,
		Description: description,
		ImageUrl:    imageUrl,
	}

	if result := s.db.Create(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil

}
