package repositories

import (
	"gorm-imp/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTrans() ([]models.Transaction, error)
	FindTran(id int) (models.Transaction, error)
	CreateTrans(trans models.Transaction) (models.Transaction, error)
	UpdateTrans(trans models.Transaction, id int) (models.Transaction, error)
	// DeleteTrip(trip models.Trip, id int) (models.Trip, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTrans() ([]models.Transaction, error) {
	var trans []models.Transaction
	// err := r.db.Raw("SELECT * FROM trips").Scan(&trips).Error

	err := r.db.Preload("Country").Find(&trans).Error

	return trans, err
}

// this func begin handling database items with object relation models
func (r *repository) FindTran(id int) (models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Preload("Country").First(&trans, id).Error // Using Find method

	// fmt.Println(trans)

	return trans, err
}

func (r *repository) CreateTrans(trans models.Transaction) (models.Transaction, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Debug().Preload("Country").Create(&trans).Error

	return trans, err

}

func (r *repository) UpdateTrans(trans models.Transaction, id int) (models.Transaction, error) {

	err := r.db.Debug().Preload("Country").Save(&trans).Error

	// err := r.db.Debug().Raw(`"UPDATE trips SET title=?, country_id=?, accomodation=?,transportation=?, eat=?, day=?, night=?, dateTrip=?, price=?, quota=?, description=?, image=? WHERE id=?"`, trip.Title, trip.Country, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	// err := r.db.Debug().Raw("UPDATE trips SET title=?, accomodation=?,transportation=?, eat=?, day=?, night=?, date_trip=?, price=?, quota=?, description=?, image=? WHERE id=?", trip.Title, trip.Accomodation, trip.Transportation, trip.Eat, trip.Day, trip.Night, trip.DateTrip, trip.Price, trip.Quota, trip.Description, trip.Image, id).Scan(&trip).Error
	return trans, err
}

// func (r *repository) DeleteTrip(trip models.Trip, ID int) (models.Trip, error) {

// 	// err := r.db.Preload("Country").Delete(&trip).Error
// 	err := r.db.Raw("DELETE FROM trips WHERE id=?", ID).Scan(&trip).Error
// 	return trip, err
// }
