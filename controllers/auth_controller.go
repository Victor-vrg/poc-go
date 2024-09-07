package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/Victor-vrg/poc-go/models"
    "github.com/Victor-vrg/poc-go/services"
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

    err := services.NewUserService(nil).RegisterUser(&user)  // Adapte o service aqui com o repositório real
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created successfully"})
}
