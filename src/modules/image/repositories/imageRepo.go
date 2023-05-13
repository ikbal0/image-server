package repositories

import (
	"image-server/src/modules/image/entities"

	"gorm.io/gorm"
)

type RepositoryImageCommand interface {
	InsertImage(Image entities.Image) (entities.Image, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	tx := &repository{db: db}
	return tx
}

func (r *repository) InsertImage(Image entities.Image) (entities.Image, error) {
	err := r.db.Debug().Create(&Image).Error

	if err != nil {
		return Image, err
	}

	return Image, nil
}
