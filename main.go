package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/joho/godotenv"
	"github.com/theognis1002/go-rotten-tomatoes/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
  
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
  
	return os.Getenv(key)
  }

func databaseInit() (db *gorm.DB) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.Movie{})
	return db
}

func checkRottenTomatoScores() {
	const rottenTomatoUrl = "https://www.rottentomatoes.com/browse/movies_in_theaters/critics:certified_fresh~sort:popular?page=1"

	// Request the HTML page.
	res, err := http.Get(rottenTomatoUrl)
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

	db := databaseInit()
	// Find the review items
	doc.Find("div .discovery-tiles score-pairs").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the criticScore
		movieTitle := strings.TrimSpace(s.Next().Text())
		criticScore, _ := strconv.Atoi(s.AttrOr("criticsscore", "-"))
		criticSentiment, _ := strconv.Atoi(s.AttrOr("criticsentiment", "-"))
		audienceScore, _ := strconv.Atoi(s.AttrOr("audiencescore", "-"))
		audienceSentiment, _ := strconv.Atoi(s.AttrOr("audiencesentiment", "-"))

		movie := models.Movie{Title: movieTitle, CriticScore: criticScore, CriticSentiment: criticSentiment, AudienceScore: audienceScore, AudienceSentiment: audienceSentiment}
		if db.Model(&movie).Where("title = ?", movieTitle).Updates(&movie).RowsAffected == 0 {
			db.Create(&movie)
		}
		fmt.Printf("[%d] %s - Score: %d%% Sentiment: %d\n", i+1, movieTitle, criticScore, criticSentiment)
	})
}

func checkAmcTheatreNowPlaying() {
	const amcTheatreUrl = "https://www.amctheatres.com/movies"
	// Request the HTML page.
	res, err := http.Get(amcTheatreUrl)
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
	doc.Find("div .PosterContent h3").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the criticScore
		movieTitle := strings.TrimSpace(s.Text())
		fmt.Printf("[%d] %s\n", i+1, movieTitle)
	})
}

func main() {
	TWILIO_ACCOUNT_SID := goDotEnvVariable("TWILIO_ACCOUNT_SID")
	TWILIO_AUTH_TOKEN := goDotEnvVariable("TWILIO_AUTH_TOKEN")
	fmt.Printf("TWILIO_ACCOUNT_SID: %s", TWILIO_ACCOUNT_SID)
	fmt.Printf("TWILIO_AUTH_TOKEN: %s", TWILIO_AUTH_TOKEN)

	db := databaseInit()

	// Migrate the schema
	db.AutoMigrate(&models.Movie{})

	// Create
	// db.Create(&models.Movie{Title: "D42", RottenTomatoScore: 100})

	// Read
	var movie models.Movie
	var movies []models.Movie
	// db.First(&movie, 1)                  // find movie with integer primary key
	db.First(&movie, "title = ?", "D42") // find movie with code D42

	db.Find(&movies)
	fmt.Println(movie)
	fmt.Println(movies)

	// Update - update movie's RottenTomatoScore to 200
	// db.Model(&movie).Update("RottenTomatoScore", 200)
	// Update - update multiple fields
	// db.Model(&movie).Updates(models.Movie{RottenTomatoScore: 200}) // non-zero fields

	// Delete - delete movie
	// db.Delete(&movie, 1)

	checkRottenTomatoScores()
	// checkAmcTheatreNowPlaying()
}
