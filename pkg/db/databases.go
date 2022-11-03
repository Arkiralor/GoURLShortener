package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/arkiralor/GoURLShortener/pkg/models"
)

func GetDBAttributes() (string, string, string, string, string) {
	db_name := os.Getenv("PGDATABASE")
	username := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")

	return db_name, username, password, host, port
}

func Init() *gorm.DB {
	db_name, username, password, host, port := GetDBAttributes()
	dbUrl := fmt.Sprintf("postgres://%s:%ss@%s:%s/%s", username, password, host, port, db_name)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error:\t%v", err)
	}

	db.AutoMigrate(&models.UniversalResourceLocator{}, &models.User{})
	return db
}
