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

	err := r.db.Debug().Preload("Trip").Preload("Trip.Country").Find(&trans).Error

	return trans, err
}

// this func begin handling database items with object relation models
func (r *repository) FindTran(id int) (models.Transaction, error) {
	var trans models.Transaction
	err := r.db.Preload("Trip").Preload("Trip.Country").First(&trans, id).Error // Using Find method

	// fmt.Println(trans)

	return trans, err
}

func (r *repository) CreateTrans(trans models.Transaction) (models.Transaction, error) {
	// err := r.db.Find("INSERT INTO trips(title,country,accomodation,transportation,eat,day,night,dateTrip,price,quota,description,image) ,
	err := r.db.Debug().Preload("Trip").Preload("Trip.Country").Create(&trans).Error

	return trans, err

}

func (r *repository) UpdateTrans(trans models.Transaction, id int) (models.Transaction, error) {

	// err := r.db.Debug().Model(&trans).Updates(trans).Error

	err := r.db.Debug().Raw("UPDATE transactions SET counter_qty=?, total=?,status=?, attachment=?, trip_id=? WHERE id=?", trans.CounterQty, trans.Total, trans.Status, trans.Attachment, trans.TripId, id).Scan(&trans).Error

	return trans, err
}
