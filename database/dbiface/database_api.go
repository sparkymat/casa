package dbiface

type DatabaseAPI interface {
	AutoMigrate() error
}
