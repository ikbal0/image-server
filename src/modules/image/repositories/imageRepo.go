package repositories

import (
	"image-server/src/modules/image/entities"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	tx := &repository{db: db}
	return tx
}

type RepositoryImageCommand interface {
	InsertImage(Image entities.Image) (entities.Image, error)
	ImageByID(ID int) (entities.Image, error)
	UpdateImage(Image entities.Image) (entities.Image, error)
}

func (r *repository) ImageByID(ID int) (entities.Image, error) {
	var image entities.Image
	if err := r.db.First(&image, "Id = ?", ID).Error; err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) UpdateImage(Image entities.Image) (entities.Image, error) {
	var input entities.Image

	imageName := Image.Name
	input.Name = imageName

	r.db.Model(&Image).Updates(&input)

	return Image, nil
}

func (r *repository) InsertImage(Image entities.Image) (entities.Image, error) {
	err := r.db.Debug().Create(&Image).Error

	if err != nil {
		return Image, err
	}

	return Image, nil
}
