package main

import (
	"log"
	"os"
	"zapi/handler"
	"zapi/repository"
	"zapi/usecase"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	userInstance := os.Getenv("ZAPI_USER_INSTANCE")
	userToken := os.Getenv("ZAPI_USER_TOKEN")
	accountToken := os.Getenv("ZAPI_ACCOUNT_TOKEN")

	if userInstance == "" || userToken == "" || accountToken == "" {
		log.Println("As integrações com o Z-API não estão configuradas corretamente")
	}

	repo := repository.NewRepository()
	usecase := usecase.NewUsecase(repo)
	handler := handler.NewHandler(usecase)

	app.Get("/getchats", handler.GetChats)

	app.Listen(":8080")
}
