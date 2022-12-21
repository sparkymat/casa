package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
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
