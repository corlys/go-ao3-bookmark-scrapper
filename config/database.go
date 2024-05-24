package config

import (
	"go-scrapper/domain/dao"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectToDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to DB: ", err)
	}
	errMigrateDB := db.AutoMigrate(&dao.Users{}, &dao.Works{}, &dao.Bookmarks{})
	if errMigrateDB != nil {
		log.Fatal("Error migrating db: ", err)
	}
	return db
}
