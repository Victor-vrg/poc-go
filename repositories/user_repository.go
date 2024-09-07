package repositories

import (
    "github.com/Victor-vrg/poc-go/models"
    "gorm.io/gorm"
)

type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
    return &UserRepository{db}
}

// Cria um novo usuário
func (repo *UserRepository) Create(user *models.User) error {
    return repo.db.Create(user).Error
}

// Busca um usuário por login
func (repo *UserRepository) FindByLogin(companyLogin string) (*models.User, error) {
    var user models.User
    err := repo.db.Where("company_login = ?", companyLogin).First(&user).Error
    return &user, err
}
