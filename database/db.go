package database

import (
	"fmt"
	"image-server/src/modules/image/entities"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = os.Getenv("PGHOST")
	port     = os.Getenv("PGPORT")
	user     = os.Getenv("PGUSER")
	password = os.Getenv("PGPASSWORD")
	dbname   = os.Getenv("PGDATABASE")
	db       *gorm.DB
	err      error
)

// var (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgresql"
// 	dbname   = "image-server-go"
// 	db       *gorm.DB
// 	err      error
// )

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode= disable", host, user, password, dbname, port)
	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	db.Debug().AutoMigrate(entities.Image{})
}

func GetDB() *gorm.DB {
	return db
}
