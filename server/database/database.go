package database

import (
	"gallery/server/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

func OpenDB(database string) *gorm.DB {

	databaseDriver := os.Getenv("DATABASE_DRIVER")

	db, err := gorm.Open(databaseDriver, database)
	if err != nil {
		log.Fatalf("%s", err)
	}
	if err := AutoMigrate(db); err != nil {
		panic(err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) error {

	return db.AutoMigrate(&domain.ImageInfo{}).Error

}
