package database

import (
	"go-auth/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:@/go_auth"
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DryRun: false,
	})

	if err != nil {
		panic("Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{}, &models.Role{})
}
