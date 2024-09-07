package models

import (
    "time"
)

type User struct {
    ID            uint      `json:"id" gorm:"primaryKey"`
    ClientID      string    `json:"client_id"`
    ClientSecret  string    `json:"client_secret"`
    CompanyLogin  string    `json:"company_login"`
    CompanyPassword string  `json:"company_password"`
    CreatedAt     time.Time `json:"created_at"`
    UpdatedAt     time.Time `json:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at"`
}
