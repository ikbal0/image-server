package repositories

import (
	"fmt"
	"image-server/src/modules/image/entities"
)

type RepositoryImageCommand interface {
	InsertImage(Image entities.Image) (entities.Image, error)
	UpdateImage(Image entities.Image) (entities.Image, error)
	Delete(Image entities.Image) error
	ImageByID(ID int) (entities.Image, error)
}

func (r *repository) ImageByID(ID int) (entities.Image, error) {
	var image entities.Image
	if err := r.db.First(&image, "Id = ?", ID).Error; err != nil {
		return image, err
	}
	return image, nil
}

func (r *repository) Delete(Image entities.Image) error {
	err := r.db.Debug().Delete(Image).Error

	return err
}

func (r *repository) UpdateImage(Image entities.Image) (entities.Image, error) {
	var input entities.Image

	input.Name = Image.Name
	input.ImageUrl = Image.ImageUrl
	input.UserID = Image.UserID

	r.db.Model(&Image).Updates(&input)

	fmt.Println(input.Name, input.ImageUrl, input.UserID)

	return Image, nil
}

func (r *repository) InsertImage(Image entities.Image) (entities.Image, error) {
	err := r.db.Debug().Create(&Image).Error

	return Image, err
}
