package main

import (
	"fmt"

	"github.com/AlanKabolov/prpr-v1/config"
	"github.com/AlanKabolov/prpr-v1/db"
	"github.com/AlanKabolov/prpr-v1/handlers"
	"github.com/AlanKabolov/prpr-v1/repositories"
	"github.com/gofiber/fiber/v2"
)

func main() {
	envConfig := config.NewEnvConfig()
	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "PrPr",
		ServerHeader: "Fiber",
	})

	//rep
	EventRepository := repositories.NewEventRepository(db)

	//rout
	server := app.Group("/api")

	//hand
	handlers.NewEventHandler(server.Group("/event"), EventRepository)
	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
