package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

func (s *Service) CreateCatalogItem(_ context.Context, title string, url string, description string, imageURL string) (*model.CatalogItem, error) {
	item := model.CatalogItem{
		Title:       title,
		URL:         url,
		Description: description,
		ImageURL:    imageURL,
	}

	if result := s.db.Create(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (s *Service) ListCatalogItems(_ context.Context) ([]model.CatalogItem, error) {
	catalogItems := []model.CatalogItem{}

	if result := s.db.Order("title desc").Find(&catalogItems); result.Error != nil {
		return nil, result.Error
	}

	return catalogItems, nil
}
