package main

import (
	"encoding/json"
	"fmt"
	"log"
	
	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"06-chinook-gorm/models"
)

func init() {
	log.Println("Program initialazing")
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
			fmt.Println(string(b))
	}
	return
}

func PrintJsonShort(v interface{}) (err error) {
	b, err := json.Marshal(v)
	if err == nil {
		fmt.Println(string(b))
	}
	return
}


func main() {

	// connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := gorm.Open("postgres", connStr)
	// check(err)
	// defer db.Close()
	// db.SingularTable(true)

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

	// fmt.Println("Customer*****************************************************************************************************************************************")
	// var customer models.Customer
	
	// // db.First(&Customer)
	// db.Preload("Invoices.InvoiceLines").First(&customer)
	// PrettyPrint(customer)

	customer := models.CustomerInvoicesInvoiceItems(2)
	PrintJsonShort(customer)
	// PrettyPrint(customer)


}