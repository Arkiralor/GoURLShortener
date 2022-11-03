package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UniversalResourceLocator struct {
	gorm.Model

	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	FullURL   string    `json:"fullUrl"`
	ShortUrl  string    `json:"shortUrl"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type User struct {
	gorm.Model
	ID         uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid()"`
	FirstName  string    `json:"fullName"`
	MiddleName string    `json:"middleName"`
	LastName   string    `json:"lastName"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
