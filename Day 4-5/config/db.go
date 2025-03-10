package config

import (
	"day4/models"
	"fmt"
	"log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectDatabase() {
	dsn := "root:root@tcp(127.0.0.1:3306)/retailer?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db

	fmt.Println("Database connected successfully!")

	DB.AutoMigrate(&models.Customer{}, &models.Product{}, &models.Order{}, &models.OrderItem{}, &models.Transaction{})

	fmt.Println("DB Migrated applied successfully!")
}
