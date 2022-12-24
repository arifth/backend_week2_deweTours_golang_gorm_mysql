package models

import (
	"gorm.io/gorm"
)

// type User struct {
// 	ID        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	Email     string    `json:"email"`
// 	Password  string    `json:"password"`
// 	CreatedAt time.Time `json:"created_at"`
// 	UpdatedAt time.Time `json:"updated_at"`
// }

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    int    `json:"phone"`
	Address  string `json:"address"`
}
