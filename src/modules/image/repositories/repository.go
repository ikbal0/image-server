package repositories

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	tx := &repository{db: db}
	return tx
}
