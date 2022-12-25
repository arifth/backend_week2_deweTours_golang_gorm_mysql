package repositories

import (
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

// // CreateTrip implements TripRepository
// func (*repository) CreateTrip(trip models.Trip) (models.Trip, error) {
// 	panic("unimplemented")
// }
