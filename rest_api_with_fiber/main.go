package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	RESTapi := fiber.New()
	RESTapi.Use(cors.New())
	customers := RESTapi.Group("/customers")
	customers.Get("/", func(context *fiber.Ctx) error {
		return context.SendString("Customers are being loaded")
	})
	log.Fatal(RESTapi.Listen(":5000"))
}
