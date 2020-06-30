package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Password struct {
	Username     string `json:"username" gorm:"not_null"`
	PasswordHash []byte `json:"password_hash" gorm:"not_null"`
}

var PswdDB *gorm.DB

func ConnectLoginInfo() {
	database, err := gorm.Open("sqlite3", "login.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Password{})

	PswdDB = database
}
