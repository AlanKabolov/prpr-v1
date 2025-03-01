package main

import (
	"github.com/AlanKabolov/prpr-v1/handlers"
	"github.com/AlanKabolov/prpr-v1/repositories"
	"github.com/gofiber/fiber/v2"
)


func main() {
	app:=fiber.New(fiber.Config{
		AppName:"PrPr",
		ServerHeader:"Fiber",
	})


	//rep
	EventRepository:=repositories.NewEventRepository(nil)

	//rout
	server:=app.Group("/api")

	//hand
	handlers.NewEventHandler(server.Group("/event"),EventRepository)
	app.Listen(":8080")
}