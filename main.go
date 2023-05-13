package main

import (
	"image-server/database"
	"image-server/internals/protocols/http/routers"
)

func main() {
	var PORT = "3000"
	database.StartDB()
	r := routers.StartApp()
	defer r.Run(":" + PORT)

	// http.HandleFunc("/image", controllers.UploadImage)

	// fmt.Println("listen to port 8080")

	// http.ListenAndServe(":8080", nil)

	// router := gin.Default()

	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	// imageRouter := router.Group("/image")
	// {
	// 	imageRouter.POST("/", controllers.UploadImage)
	// }
	// router.Run(":8080")
}
