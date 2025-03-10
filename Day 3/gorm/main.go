package main

import (
	"fmt"
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type User struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"type:varchar(100)"`
	Age  int
}

func main() {
	dsn := "root:demo123@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// Insert a new user
	user := User{Name: "John Doe", Age: 30}
	if err := db.Create(&user).Error; err != nil {
		log.Fatal("Error inserting record: ", err)
	}

	// Retrieve all users from the database
	var users []User
	if err := db.Find(&users).Error; err != nil {
		log.Fatal("Error retrieving records: ", err)
	}

	// Print all users
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}
}
