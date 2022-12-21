package dbiface

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

type DatabaseAPI interface {
	AutoMigrate() error
	GetUser(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, email string, name string) (*model.User, error)
	CreateCatalogItem(ctx context.Context, title string, url string, description string, imageURL string) (*model.CatalogItem, error)
	DestroyCatalogItem(ctx context.Context, id uint) error
	ListCatalogItems(ctx context.Context) ([]model.CatalogItem, error)
	CreateHomeItem(ctx context.Context, userID uint, appID uint) (*model.HomeItem, error)
}
