package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Victor-vrg/poc-go/metrics"
	"github.com/Victor-vrg/poc-go/models"
	"github.com/Victor-vrg/poc-go/routes"
)

func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Configura a conexão com o banco de dados
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Executa migrations para o modelo User
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	// Cria a instância do Fiber
	app := fiber.New()

	// Middleware
	app.Use(logger.New())       // Logs de requisições
	app.Use(requestid.New())    // Gera um ID único para cada requisição

	// Setup Prometheus Metrics
	metrics.SetupPrometheus(app)

	// Setup Routes
	routes.SetupRoutes(app, db) // Passa o banco de dados para as rotas

	// Inicia o servidor
	log.Fatal(app.Listen(":3000"))
}
