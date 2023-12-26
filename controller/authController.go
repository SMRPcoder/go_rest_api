package controller

import (
	"fmt"

	"github.com/SMRPcoder/go_rest_api/database"
	"github.com/SMRPcoder/go_rest_api/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusNoContent).JSON(fiber.Map{"error": "Error While Creating A User"})
	}

	c.Status(200).JSON(fiber.Map{"message": "data saved", "data": user, "status": true})
	return nil
}

func Login(c *fiber.Ctx) error {

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	result := database.DB.Where("username = ?", user.Username).First(&user)
	fmt.Println(result)
	c.Status(200).JSON(fiber.Map{"message": "logged in successfully", "user": user, "status": true})
	return nil

}
