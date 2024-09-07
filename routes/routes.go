package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"github.com/Victor-vrg/poc-go/controllers"
	"github.com/Victor-vrg/poc-go/middleware"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// Grupo de rotas de autenticação
	auth := app.Group("/auth")
	auth.Post("/register", controllers.Register)
	auth.Post("/login", controllers.Login)

	// Rota protegida com middleware de autenticação
	protected := app.Group("/user", middleware.AuthMiddleware)
	protected.Get("/me", func(c *fiber.Ctx) error { return controllers.GetProfile(c, db) })
}
