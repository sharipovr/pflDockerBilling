package storage

import (
	"log"

	"github.com/sharipovr/pflDockerBilling/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=rustemsharipov password=postgres dbname=pfl_docker_billing port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Subscription{})
	DB = db
}
