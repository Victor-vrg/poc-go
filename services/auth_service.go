package services

import (
    "github.com/dgrijalva/jwt-go"
    "time"
    "os"
)

type Claims struct {
    UserID int `json:"user_id"`
    jwt.StandardClaims
}

// Função para gerar tokens JWT
func GenerateJWT(userID int, isRefresh bool) (string, error) {
    secret := os.Getenv("JWT_SECRET")
    expiration := time.Hour * time.Duration(1) // Default é 1 hora

    if isRefresh {
        expiration = time.Hour * time.Duration(2)
    }

    claims := &Claims{
        UserID: userID,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(expiration).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(secret))
}
