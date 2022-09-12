package database

import (
	"fmt"
	"waysbean/models"
	"waysbean/pkg/mysql"
)

// Automatic Migration if Running App
func RunMigration() {

	err := mysql.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Transaction{},
		&models.Cart{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
