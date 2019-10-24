package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Define the gorm basic model
// type Model struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
// }

type Artist struct {
	// gorm.Model // Extend our own model with the properties given at the gorm basic model
	ArtistId int `gorm:"primary_key"`
	Name     string
	Albums   []Album `gorm:"foreignkey:ArtistId"`
}

type Album struct {
	AlbumId  int `gorm:"primary_key"`
	Title    string
	ArtistId int
	Artist   Artist `gorm:"foreignkey:ArtistId"`
}

type Customer struct {
	CustomerId int `gorm:"primary_key"`
	Email      string
	Invoices   []Invoice `gorm:"foreignkey:CustomerId"`
}

type Invoice struct {
	InvoiceId    int      `gorm:"primary_key"`
	Customer     Customer `gorm:"foreignkey:CustomerId"`
	CustomerId   int
	Total        float32
	InvoiceLines []InvoiceLine `gorm:"foreignkey:InvoiceId"`
}

type InvoiceLine struct {
	InvoiceLineId int     `gorm:"primary_key"`
	Invoice       Invoice `gorm:"foreignkey:InvoiceId"`
	InvoiceId     int
	TrackId       int
	UnitPrice     float32
	Quantity      int
}

func check(error interface{}) {
	if error != nil {
		log.Fatal(error)
	}
}

var port = 5432
var host = os.Getenv("DB_HOST")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var dbname = os.Getenv("DB")

var db *gorm.DB

func init() {
	log.Println("Opening session to database")

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = gorm.Open("postgres", connStr)
	check(err)
	db.SingularTable(true)

}

func CustomerInvoicesInvoiceItems(customerId uint) Customer {

	fmt.Println("Customer*****************************************************************************************************************************************")
	var customer Customer

	// db.First(&Customer)
	db.Preload("Invoices.InvoiceLines").First(&customer, customerId)
	return customer

}
