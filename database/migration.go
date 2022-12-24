package database

import (
	"fmt"
	"gorm-imp/models"
	"gorm-imp/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {

	// NOTE: migrate parent models first ,then child models after, otherwise it wont create the relation table between both

	err := mysql.DB.AutoMigrate(&models.User{}, &models.Transaction{}, &models.Country{}, &models.Trip{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Succes")
}
