package api

import (
	"go-structure-folder-clean/cmd/entity"
	"time"
)

type User = entity.User

type UserFormatterRegister struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
}

type UserFormaterReturnLogin struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	Occupation     string    `json:"occupation"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           int       `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FormatUserRegister(user User, token string) UserFormatterRegister {
	formatter := UserFormatterRegister{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Occupation: user.Occupation,
		Token:      token,
	}

	return formatter
}

func FormatUserLogin(user User, token string) UserFormaterReturnLogin {
	formatter := UserFormaterReturnLogin{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		Occupation:     user.Occupation,
		AvatarFileName: user.Occupation,
		Role:           user.Role,
		Token:          token,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}

	return formatter
}
