package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/sharipovr/pflDockerBilling/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	// Контейнерная версия подключения к БД
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	// Обычная (не-контейнерная) версия подключения к БД
	// dsn := "host=localhost user=rustemsharipov password=postgres dbname=pfl_docker_billing port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}

	db.AutoMigrate(&models.User{}, &models.Subscription{})
	DB = db
}
