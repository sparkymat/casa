package database

import (
	"path"

	"git.orion.home/oxhead/casa/model"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	DataPath string
}

func New(cfg Config) (*Service, error) {
	dbPath := path.Join(cfg.DataPath, "store.db")

	gormDB, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		//nolint:wrapcheck
		return nil, err
	}

	return &Service{db: gormDB}, nil
}

type Service struct {
	db *gorm.DB
}

func (s *Service) AutoMigrate() error {
	//nolint:wrapcheck
	return s.db.AutoMigrate(
		&model.User{},
	)
}
