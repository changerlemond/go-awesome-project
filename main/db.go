package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func connectToDatabase() {
	dsn := "host=localhost user=postgres password=test dbname=postgres port=5432 sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %v", err))
	}
	db.AutoMigrate(&User{})
}

type User struct {
	Id        uint           `json:"id" gorm:"primary_key"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Role      string         `json:"role"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
