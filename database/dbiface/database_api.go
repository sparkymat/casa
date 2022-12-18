package dbiface

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

type DatabaseAPI interface {
	AutoMigrate() error
	GetUser(ctx context.Context, email string) (*model.User, error)
	CreateUser(ctx context.Context, email string, name string) (*model.User, error)
}
