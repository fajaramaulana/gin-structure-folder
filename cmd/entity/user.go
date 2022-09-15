package entity

import "time"

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Occupation     string    `json:"occupation"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"password_hash"`
	AvatarFileName string    `json:"avatar_file_name"`
	Role           int       `json:"role"`
	Token          string    `json:"token"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
