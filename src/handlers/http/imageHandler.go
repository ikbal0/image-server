package handler

import (
	"image-server/database"
	"image-server/internals/utils"
	"image-server/src/modules/image/dto"
	"image-server/src/modules/image/entities"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateImage(ctx *gin.Context) {
	var db = database.GetDB()
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}

	var image entities.Image

	if err := db.First(&image, "Id = ?", ctx.Param("imageId")).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	imageName := image.Name
	if err := ctx.SaveUploadedFile(file, "image/"+imageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	var input entities.Image
	input.Name = imageName

	ctx.ShouldBindJSON(&input)

	db.Model(&image).Updates(&input)

	ctx.JSON(http.StatusOK, gin.H{"data": image})
}

func DeleteImage(ctx *gin.Context) {
	db := database.GetDB()
	imageDelete := entities.Image{}
	imageId, _ := strconv.Atoi(ctx.Param("imageId"))

	err := db.First(&imageDelete, "Id = ?", imageId).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record has not found!"})
		return
	}

	errRemove := os.Remove("image/" + imageDelete.Name)

	if errRemove != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   errRemove.Error(),
			"message": "file has not found!",
		})
		return
	}

	db.Delete(&imageDelete)

	ctx.JSON(http.StatusOK, gin.H{"message": "image deleted!"})
}

func UploadImage(ctx *gin.Context) {
	db := database.GetDB()
	Image := entities.Image{}

	file, err := ctx.FormFile("file")
	// name := ctx.PostForm("name")
	userId := ctx.PostForm("userId")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}

	timeStamp := utils.MakeTimeStamp()
	newName := strconv.Itoa(int(timeStamp)) + file.Filename

	// if err := ctx.SaveUploadedFile(file, "image/"+name); err != nil {
	if err := ctx.SaveUploadedFile(file, "image/"+newName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	num, errConv := strconv.Atoi(userId)

	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to convert data",
		})
		return
	}

	Image.UserID = uint(num)
	// Image.Name = name
	Image.Name = newName
	// Image.ImageUrl = "http://localhost:3000/image/" + name
	Image.ImageUrl = "http://localhost:3000/image/" + newName

	errCreate := db.Debug().Create(&Image).Error

	if errCreate != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		// "name":      name,
		"image_id":  Image.ID,
		"user_id":   Image.UserID,
		"image_url": Image.ImageUrl,
	})
}

func (h HttpHandlerImpl) UploadImageVer2(ctx *gin.Context) {
	userId := ctx.PostForm("userId")
	num, errConv := strconv.Atoi(userId)

	if errConv != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to convert data",
		})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}

	timeStamp := utils.MakeTimeStamp()
	newName := strconv.Itoa(int(timeStamp)) + file.Filename
	imageUrl := "http://localhost:3000/image/" + newName

	// save file
	if err := ctx.SaveUploadedFile(file, "uploads/image/"+newName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	imageRequestBody := dto.ImageRequestBody{
		Name:     newName,
		ImageUrl: imageUrl,
		UserID:   uint(num),
	}

	image, err := h.InsertImage(imageRequestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		// "name":      name,
		"image_id":  image.ID,
		"user_id":   image.UserID,
		"image_url": image.ImageUrl,
	})
}
