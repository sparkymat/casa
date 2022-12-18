package dbiface

import "git.orion.home/oxhead/casa/model"

type DatabaseAPI interface {
	AutoMigrate() error
	GetUser() (*model.User, error)
}
