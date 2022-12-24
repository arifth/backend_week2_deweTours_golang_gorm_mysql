package models

import (
	"time"

	"gorm.io/gorm"
)

type Trip struct {
	gorm.Model
	Title          string    `json:"title"`
	CountryId      int       `json:"country_id"`
	Country        Country   `json:"country"`
	Accomodation   string    `json:"accomodation"`
	Transportation string    `json:"transportation"`
	Eat            string    `json:"eat"`
	Day            int       `json:"day"`
	Night          int       `json:"night"`
	DateTrip       time.Time `json:"date_trip"`
	Price          int       `json:"price"`
	Quota          int       `json:"quota"`
	Description    string    `json:"description"`
	Image          string    `json:"string"`
}