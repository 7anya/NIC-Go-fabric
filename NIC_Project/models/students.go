package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Student struct {
	ID              uint   `json:"id" gorm:"primary_key; not_null;unique;AUTO_INCREMENT"`
	StudentName     string `json "student_name" gorm:"not_null"`
	BoardName       string `json:"board_name" gorm:"not_null"`
	ExaminationName string `json:"examination_name" gorm:"not_null"`
	SchoolName      string `json:"school_name" gorm:"not_null"`
	YearOfExam      uint   `json:"year_of_exam" gorm:"not_null"`
	StudentRollNo   string `json:"student_roll_no" gorm:"not_null"`
}

var DB *gorm.DB

func ConnectDataBase() {
	database, err := gorm.Open("sqlite3", "student_data.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Student{})

	DB = database
}
