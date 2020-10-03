package domain

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"time"
)


func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv("./../.env"))
	if err != nil {
		log.Fatal("Ensure that the env is located in the server root. ", err)
	}

	os.Exit(m.Run())
}

//func DBConn() *gorm.DB {
//
//	return LocalDatabase()
//}

//Local DB
func DBConn() (*gorm.DB, error) {
	dbdriver := os.Getenv("TEST_DB_DRIVER")
	host := os.Getenv("TEST_DB_HOST")
	password := os.Getenv("TEST_DB_PASSWORD")
	user := os.Getenv("TEST_DB_USER")
	dbname := os.Getenv("TEST_DB_NAME")
	port := os.Getenv("TEST_DB_PORT")

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	conn, err := gorm.Open(dbdriver, DBURL)
	if err != nil {
		return nil, err
	} else {
		log.Println("CONNECTED TO: ", dbdriver)
	}

	err = conn.DropTableIfExists(&ImageInfo{}).Error
	if err != nil {
		return nil, err
	}
	err = conn.Debug().AutoMigrate(
		ImageInfo{},
	).Error
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func SeedImageInfo(db *gorm.DB) (*ImageInfo, error) {
	info := &ImageInfo{
		ID:        1,
		ImageId:    1,
		UserIp:    "123.09.45",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := db.Create(&info).Error
	if err != nil {
		return nil, err
	}
	return info, nil
}

func SeedManyImageInfo(db *gorm.DB) ([]ImageInfo, error) {
	data := []ImageInfo{
		{
			ID:        1,
			ImageId:    1,
			UserIp:    "123.09.45",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			ImageId:    2,
			UserIp:    "990.09.45",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	for _, v := range data {
		err := db.Create(&v).Error
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}