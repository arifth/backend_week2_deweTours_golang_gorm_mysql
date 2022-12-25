package models

import (
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name string `json:"name"`
}

type CountryResponse struct {
	gorm.Model
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "Countries"
}
