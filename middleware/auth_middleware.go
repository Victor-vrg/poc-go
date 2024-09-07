package middleware

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/dgrijalva/jwt-go"
	"os"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Captura o cabeçalho Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Missing or malformed JWT"})
	}

	// Verifica se o token começa com "Bearer "
	tokenString := strings.Split(authHeader, "Bearer ")
	if len(tokenString) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid token format"})
	}

	// Parse o token JWT
	token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Unexpected signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Avança para a próxima etapa se o token for válido
	return c.Next()
}
