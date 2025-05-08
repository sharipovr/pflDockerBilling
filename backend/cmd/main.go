package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/sharipovr/pflDockerBilling/routes"
	"github.com/sharipovr/pflDockerBilling/storage"
)

func main() {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load("../.env.local"); err != nil {
		log.Println("Файл .env.local не найден, используются переменные окружения системы")
		if err := godotenv.Load(".env.docker"); err != nil {
			log.Println("Файл .env.docker не найден, используются переменные окружения системы")
		}
	}
	storage.Init()

	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:: ", err)
	}
}
