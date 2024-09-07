package services

import (
    "github.com/Victor-vrg/poc-go/models"
    "github.com/Victor-vrg/poc-go/repositories"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
    return &UserService{repo}
}

// Cria um novo usuário com senha criptografada
func (s *UserService) RegisterUser(user *models.User) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.CompanyPassword), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.CompanyPassword = string(hashedPassword)
    return s.repo.Create(user)
}

// Verifica login e senha
func (s *UserService) Authenticate(companyLogin, password string) (*models.User, error) {
    user, err := s.repo.FindByLogin(companyLogin)
    if err != nil {
        return nil, err
    }

    // Verifica a senha
    err = bcrypt.CompareHashAndPassword([]byte(user.CompanyPassword), []byte(password))
    if err != nil {
        return nil, err
    }

    return user, nil
}
