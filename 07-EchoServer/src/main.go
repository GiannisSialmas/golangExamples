package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// getListenAdress configures a listen adress or the echo server
func getListenAdress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("Defaulting to port %s", port)
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		log.Fatalln("Invalid port", port)
	}
	return fmt.Sprintf("0.0.0.0:%s", port)
}

func main() {
	app := echo.New()

	// Middleware
	// Recover from panics anywhere in the chain, prints stack trace and handles the control to the centralized HTTPErrorHandler.
	app.Use(middleware.Recover())
	// Log requests
	// app.Use(middleware.Logger())

	// Route => handler
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	app.GET("/genres", func(c echo.Context) error {
		genresString, err := json.Marshal(GenresGet())
		if err != nil {
			panic(err)
		}
		fmt.Printf("%T\n", genresString)
		fmt.Println(genresString)
		// return c.String(http.StatusOK, []byte(genresString))

		return c.String(http.StatusOK, "hello")

	})

	// Start server
	listenAdress := getListenAdress()
	app.Logger.Fatal(app.Start(listenAdress))
}

// GenresGet returns all the genres in the database
type Genre struct {
	GenreId int `gorm:"primary_key"`
	Name    string
}

func GenresGet() []Genre {

	log.Println("Opening session to database")

	var port = 5432
	var host = os.Getenv("DB_HOST")
	var user = os.Getenv("DB_USER")
	var password = os.Getenv("DB_PASSWORD")
	var dbname = os.Getenv("DB")

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	db.SingularTable(true)

	var genres []Genre
	db.Find(&genres)
	return genres
}
