package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "05-chinook-gorm/models"
)

// Define the gorm basic model
// type Model struct {
// 	ID        uint `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt *time.Time
// }

var port = 5432
var host = os.Getenv("DB_HOST")
var user = os.Getenv("DB_USER")
var password = os.Getenv("DB_PASSWORD")
var dbname = os.Getenv("DB")

func init() {
	log.Println("Program initialazing")
}

func check(error interface{}) {
	if error != nil {
		log.Fatal(error)
	}
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
			fmt.Println(string(b))
	}
	return
}

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := gorm.Open("postgres", connStr)
	check(err)
	defer db.Close()
	db.SingularTable(true)

	// fmt.Println("**********************************************************************************************************************************************")
	// var artist Artist

	// db.Preload("Albums.Title").First(&artist)
	// // db.First(&artist)
	// fmt.Println(artist)

	// fmt.Println("**********************************************************************************************************************************************")
	// var album Album

	// // db.Preload("Artist").First(&album)
	// db.First(&album)
	// fmt.Println(album)

	// fmt.Println("INVOICE*****************************************************************************************************************************************")
	// var invoice Invoice
	
	// // db.First(&invoice)
	// db.Preload("InvoiceLines").Take(&invoice)
	// PrettyPrint(invoice)

	// fmt.Println("INVOICELISTS*************************************************************************************************************************************")
	// var invoiceLine InvoiceLine
	
	// // db.First(&invoiceLine)
	// db.Preload("Invoice").Take(&invoiceLine)
	// PrettyPrint(invoiceLine)

	// log.Print("Program completed without a problem")

	fmt.Println("Customer*****************************************************************************************************************************************")
	var customer Customer
	
	// db.First(&Customer)
	db.Preload("Invoices.InvoiceLines").Take(&customer)
	PrettyPrint(customer)



}
