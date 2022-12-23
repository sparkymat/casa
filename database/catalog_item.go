package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
	"gorm.io/gorm"
)

func (s *Service) ListCategories(_ context.Context) ([]string, error) {
	categories := []string{}

	if result := s.db.Model(&model.CatalogItem{}).Distinct().Pluck("Category", &categories); result.Error != nil {
		return nil, result.Error
	}

	return categories, nil
}

func (s *Service) CreateCatalogItem(_ context.Context, title string, url string, category string, description string, imageURL string) (*model.CatalogItem, error) {
	item := model.CatalogItem{
		Title:       title,
		URL:         url,
		Cataegory:   category,
		Description: description,
		ImageURL:    imageURL,
	}

	if result := s.db.Create(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (s *Service) DestroyCatalogItem(_ context.Context, id uint) error {
	item := model.CatalogItem{
		Model: gorm.Model{
			ID: id,
		},
	}

	result := s.db.Delete(&item)

	return result.Error
}

func (s *Service) ListCatalogItems(_ context.Context) ([]model.CatalogItem, error) {
	catalogItems := []model.CatalogItem{}

	if result := s.db.Order("title desc").Find(&catalogItems); result.Error != nil {
		return nil, result.Error
	}

	return catalogItems, nil
}
