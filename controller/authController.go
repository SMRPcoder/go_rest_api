package controller

import (
	"github.com/SMRPcoder/go_rest_api/database"
	"github.com/SMRPcoder/go_rest_api/functions"
	"github.com/SMRPcoder/go_rest_api/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
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

	var requser models.User
	if err := c.BodyParser(&requser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	var user models.User
	result := database.DB.Where("username = ?", requser.Username).First(&user)
	if result.Error != nil {
		// log.Fatal(result.Error)
		return c.Status(200).JSON(fiber.Map{"message": "User Not Found", "status": false})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requser.Password))
	if err != nil {
		return c.Status(200).JSON(fiber.Map{"message": "Password Missmatch", "status": false})
	}

	token, err := functions.EncodeJwt(functions.JWTUser{ID: user.ID, Username: user.Username, Name: user.Name})
	if err != nil {
		return c.Status(200).JSON(fiber.Map{"message": err, "status": false})
	}
	c.Status(200).JSON(fiber.Map{"message": "logged in successfully", "token": "Bearar " + token, "status": true})
	return nil

}
