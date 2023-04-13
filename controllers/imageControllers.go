package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// var (
// 	appJson = "application/json"
// )

func UploadImage(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	name := ctx.PostForm("name")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "No File is received",
			"err":     err.Error(),
		})
		return
	}

	if err := ctx.SaveUploadedFile(file, "image/"+file.Filename); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save file",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

// func UploadImage(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "POST" {
// 		http.Error(w, "", http.StatusBadRequest)
// 		return
// 	}

// 	if err := r.ParseMultipartForm(1024); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	alias := r.FormValue("alias")

// 	uploadedFile, handler, err := r.FormFile("file")

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer uploadedFile.Close()

// 	dir, err := os.Getwd()

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	fileName := handler.Filename

// 	if alias != "" {
// 		fileName = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
// 	}

// 	fileLocation := filepath.Join(dir, "image", fileName)

// 	targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	defer targetFile.Close()

// 	if _, err := io.Copy(targetFile, uploadedFile); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write([]byte("done"))
// }
