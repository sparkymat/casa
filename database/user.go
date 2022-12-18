package database

import (
	"context"

	"git.orion.home/oxhead/casa/model"
)

func (s *Service) GetUser(_ context.Context, email string) (*model.User, error) {
	user := model.User{
		Email: email,
	}

	if result := s.db.First(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (s *Service) CreateUser(_ context.Context, email string, name string) (*model.User, error) {
	user := model.User{
		Email: email,
		Name:  name,
	}

	if result := s.db.Create(&user); result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
