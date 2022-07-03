package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/theognis1002/go-rotten-tomatoes/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func checkRottenTomatoScores() {
	const RottenTomatoUrl = "https://www.rottentomatoes.com/browse/movies_in_theaters/sort:newest?page=1"

	// Request the HTML page.
	res, err := http.Get(RottenTomatoUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("div .discovery-tiles score-pairs").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the criticScore
		movieTitle := strings.TrimSpace(s.Next().Text())
		criticScore := s.AttrOr("criticsscore", "-")
		criticSentiment := s.AttrOr("criticsentiment", "-")
		// audienceScore := s.AttrOr("audiencescore", "-")
		// audienceSentiment := s.AttrOr("audiencesentiment", "-")

		fmt.Printf("[%d] %s - Score: %s%% Sentiment: %s\n", i+1, movieTitle, criticScore, criticSentiment)
	})
}

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
	// checkRottenTomatoScores()
}
