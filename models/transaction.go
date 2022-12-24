package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CounterQty int    `json:"counter_qty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachement"`
	TripId     int    `json:"trip_id"`
	Trip       Trip   `json:"trip"`
	UserId     int    `json:"user_id"`
	User       User   `json:"user"`
}
