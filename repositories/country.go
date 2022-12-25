package repositories

import (
	"fmt"
	"gorm-imp/models"

	"gorm.io/gorm"
)

type CountryRepository interface {
	FindCountries() ([]models.Country, error)
	FindCountry(id int) (models.Country, error)
	CreateCountry(trip models.Country) (models.Country, error)
	UpdateCountry(trip models.Country, id int) (models.Country, error)
	DeleteCountry(trip models.Country, id int) (models.Country, error)
}

func RepositoryCountry(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindCountries() ([]models.Country, error) {
	var country []models.Country
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Find(&country).Error

	return country, err
}

// this func begin handling database items with object relation models
func (r *repository) FindCountry(id int) (models.Country, error) {
	var country models.Country
	err := r.db.First(&country, id).Error // Using Find method

	fmt.Println(country)

	return country, err
}

func (r *repository) CreateCountry(country models.Country) (models.Country, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Debug().Preload("Country").Create(&country).Error

	return country, err

}

func (r *repository) UpdateCountry(country models.Country, id int) (models.Country, error) {

	// NOTES: sementara untuk supy tidak error
	err := r.db.Save(&country).Error

	// err := r.db.Debug().Raw(`"UPDATE trips SET title=?, country_id=?, accomodation=?,transportation=?, eat=?, day=?, night=?, dateTrip=?, price=?, quota=?, description=?, image=? WHERE id=?"`, trip.Title, trip.Country, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	// err := r.db.Debug().Raw("UPDATE trips SET title=?, accomodation=?,transportation=?, eat=?, day=?, night=?, date_trip=?, price=?, quota=?, description=?, image=? WHERE id=?", trip.Title, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	return country, err
}

func (r *repository) DeleteCountry(country models.Country, ID int) (models.Country, error) {

	// err := r.db.Preload("Country").Delete(&trip).Error
	err := r.db.Raw("DELETE FROM countries WHERE id=?", ID).Scan(&country).Error
	return country, err
}
