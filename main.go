package main

import (
	"fmt"

	"github.com/SMRPcoder/go_rest_api/controller"
	"github.com/SMRPcoder/go_rest_api/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("Hii Hello World")
	fmt.Println("Hii Hello World2")

	database.Connetion()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	auth := app.Group("/auth")
	auth.Post("/register", controller.Register)
	auth.Post("/login", controller.Login)

	app.Listen(":7000")
}
