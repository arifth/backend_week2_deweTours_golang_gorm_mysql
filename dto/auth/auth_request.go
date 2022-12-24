package authdto

import "gorm.io/gorm"

type RegisterRequest struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
}

type LoginRequest struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}
