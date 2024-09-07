package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Victor-vrg/poc-go/models"
	"github.com/Victor-vrg/poc-go/services"
	"gorm.io/gorm"
)

func Register(c *fiber.Ctx) error {
	type RegisterRequest struct {
		ClientID       string `json:"client_id"`
		ClientSecret   string `json:"client_secret"`
		CompanyLogin   string `json:"company_login"`
		CompanyPassword string `json:"company_password"`
	}

	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	user := models.User{
		ClientID:       req.ClientID,
		ClientSecret:   req.ClientSecret,
		CompanyLogin:   req.CompanyLogin,
		CompanyPassword: req.CompanyPassword,
	}

	err := services.NewUserService(nil).RegisterUser(&user)  
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}

func Login(c *fiber.Ctx) error {
	// Implementar lógica de login
	return nil
}

func GetProfile(c *fiber.Ctx, db *gorm.DB) error {
	// Implementar lógica para obter o perfil do usuário
	return nil
}
