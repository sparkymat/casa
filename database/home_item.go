package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
	"gorm.io/gorm"
)

func (s *Service) CreateHomeItem(_ context.Context, userID uint, appID uint) (*model.HomeItem, error) {
	item := model.HomeItem{
		UserID:        userID,
		CatalogItemID: appID,
	}

	if result := s.db.Create(&item); result.Error != nil {
		return nil, result.Error
	}

	return &item, nil
}

func (s *Service) DestroyHomeItem(_ context.Context, id uint) error {
	item := model.HomeItem{
		Model: gorm.Model{
			ID: id,
		},
	}

	result := s.db.Delete(&item)

	return result.Error
}

func (s *Service) ListHomeItems(_ context.Context, userID uint) ([]model.HomeItem, error) {
	homeItems := []model.HomeItem{}

	if result := s.db.Joins("CatalogItem").Order("CatalogItem.title desc").Find(&homeItems, "home_items.user_id = ?", userID); result.Error != nil {
		return nil, result.Error
	}

	return homeItems, nil
}
