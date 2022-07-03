package main

import (
	"fmt"

	"github.com/theognis1002/go-rotten-tomatoes/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Movie{})

	// Create
	db.Create(&models.Movie{Title: "D42", RottenTomatoScore: 100})

	// Read
	var movie models.Movie
	var movies []models.Movie
	// db.First(&movie, 1)                  // find movie with integer primary key
	db.First(&movie, "title = ?", "D42") // find movie with code D42

	db.Find(&movies)
	fmt.Println(movies)
	// // Update - update movie's RottenTomatoScore to 200
	// db.Model(&movie).Update("RottenTomatoScore", 200)
	// // Update - update multiple fields
	// db.Model(&movie).Updates(models.Movie{RottenTomatoScore: 200}) // non-zero fields

	// // Delete - delete movie
	// db.Delete(&movie, 1)
}
