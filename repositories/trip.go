package repositories

import (
	"fmt"
	"gorm-imp/models"

	"gorm.io/gorm"
)

type TripRepository interface {
	FindTrip() ([]models.Trip, error)
	FindSingleTrip(id int) (models.Trip, error)
	CreateTrip(trip models.Trip) (models.Trip, error)
	UpdateTrip(trip models.Trip, id int) (models.Trip, error)
	DeleteTrip(trip models.Trip, id int) (models.Trip, error)
}

func RepositoryTrip(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrip() ([]models.Trip, error) {
	var trips []models.Trip
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Preload("Country").Find(&trips).Error

	return trips, err
}

// this func begin handling database items with object relation models
func (r *repository) FindSingleTrip(id int) (models.Trip, error) {
	var trip models.Trip
	err := r.db.Preload("Country").First(&trip, id).Error // Using Find method

	fmt.Println(trip)

	return trip, err
}

func (r *repository) CreateTrip(trip models.Trip) (models.Trip, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Preload("Country").Create(&trip).Error

	return trip, err

}

func (r *repository) UpdateTrip(trip models.Trip, ID int) (models.Trip, error) {

	err := r.db.Preload("Country").Save(&trip).Error
	return trip, err
}

func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {

	err := r.db.Preload("Country").Delete(&trip).Error
	return trip, err
}
