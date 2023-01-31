package dto

import (
	"github.com/google/uuid"
	"vladmsnk/billing/internal/entity"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) FromDto(newId uuid.UUID) entity.User {
	return entity.User{
		ID:       newId,
		Name:     u.Name,
		Username: u.Username,
		Email:    u.Email,
		Password: u.Password,
	}
}
