package repositories

import (
	"image-server/src/modules/image/entities"
)

type RepositoryImageCommand interface {
	InsertImage(Image entities.Image) (entities.Image, error)
	UpdateImage(ID int, Image entities.Image) (entities.Image, error)
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

func (r *repository) UpdateImage(ID int, Image entities.Image) (entities.Image, error) {
	var input entities.Image

	input.Name = Image.Name
	input.ImageUrl = Image.ImageUrl
	input.UserID = Image.UserID

	err := r.db.Debug().Model(&Image).Where("Id = ?", ID).Updates(&input).Error

	// fmt.Println(input.Name, input.ImageUrl, input.UserID)

	return Image, err
}

func (r *repository) InsertImage(Image entities.Image) (entities.Image, error) {
	err := r.db.Debug().Create(&Image).Error

	return Image, err
}
