package dbiface

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

type DatabaseAPI interface {
	AutoMigrate() error
	GetUser(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, email string, name string) (*model.User, error)
	CreateCatalogItem(ctx context.Context, title string, description string, imageUrl string) (*model.CatalogItem, error)
	ListCatalogItems(ctx context.Context) ([]model.CatalogItem, error)
}
