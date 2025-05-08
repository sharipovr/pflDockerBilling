package main

import (
	"log"

	"github.com/sharipovr/pflDockerBilling/routes"
	"github.com/sharipovr/pflDockerBilling/storage"
)

func main() {
	storage.Init()
	router := routes.SetupRouter()

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера:: ", err)
	}
}
