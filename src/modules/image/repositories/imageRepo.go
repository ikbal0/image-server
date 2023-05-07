package repositories

import (
	"image-server/internals/utils"
	"image-server/src/modules/image/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryImageCommand interface {
	InsertImage(ctx *gin.Context) (entities.Image, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	tx := &repository{db: db}
	return tx
}

func (r *repository) InsertImage(ctx *gin.Context) (entities.Image, error) {
	Image := entities.Image{}
	file, err := ctx.FormFile("file")
	// name := ctx.PostForm("name")
	userId := ctx.PostForm("userId")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return Image, err
	}

	timeStamp := utils.MakeTimeStamp()
	newName := strconv.Itoa(int(timeStamp)) + file.Filename

	// if err := ctx.SaveUploadedFile(file, "image/"+name); err != nil {
	if err := ctx.SaveUploadedFile(file, "image/"+newName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return Image, err
	}

	num, errConv := strconv.Atoi(userId)

	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to convert data",
		})
		return Image, err
	}

	Image.UserID = uint(num)
	Image.Name = newName
	Image.ImageUrl = "http://localhost:3000/image/" + newName

	// Image.Name = name
	// Image.ImageUrl = "http://localhost:3000/image/" + name
	// errCreate := db.Debug().Create(&Image).Error

	errCreate := r.db.Debug().Create(&Image).Error

	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return Image, errCreate
	}

	return Image, nil
}
