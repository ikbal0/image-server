package handler

import (
	"image-server/internals/utils"
	"image-server/src/modules/image/dto"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h HttpHandlerImpl) UpdateImageHandler(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}

	getId := ctx.Param("imageId")
	id, err := strconv.Atoi(getId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to convert string to int",
		})
		return
	}

	image, err := h.ImageByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	imageName := image.Name
	if err := ctx.SaveUploadedFile(file, "uploads/image/"+imageName); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	imageRequestBody := dto.ImageRequestBody{
		Name:     imageName,
		ImageUrl: image.ImageUrl,
		UserID:   image.UserID,
	}

	updatedImage, err := h.UpdateImage(id, imageRequestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedImage})
}

func (h HttpHandlerImpl) DeleteImageHandler(ctx *gin.Context) {
	getId := ctx.Param("imageId")
	id, err := strconv.Atoi(getId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to convert string to int",
		})
		return
	}

	image, err := h.ImageByID(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	if err := os.Remove("uploads/image/" + image.Name); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "file has not found!",
		})
		return
	}

	h.Delete(id)

	ctx.JSON(http.StatusOK, gin.H{"message": "image deleted!"})
}

func (h HttpHandlerImpl) UploadImageHandler(ctx *gin.Context) {
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
