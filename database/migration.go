package database

import (
	"fmt"
	"gorm-imp/models"
	"gorm-imp/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {
	err := mysql.DB.AutoMigrate(&models.User{})

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Succes")
}
