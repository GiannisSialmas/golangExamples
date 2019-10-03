package main


// We are importing the package so that it can register its drivers with the database/sql package, and we use the _ identifier
// to tell Go that we still want this included even though we will never directly reference the package in our code.
import (
	"database/sql"
	"fmt"
	"log"
	
	_ "github.com/lib/pq"

)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "chinook"
	dbname   = "postgres"
  )

func check(error interface{})  {
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	check(err)
	err = db.Ping()
	check(err)
	fmt.Println("Database ping succesfull")

	// rows, err := db.QueryContext(ctx, "SELECT * FROM users WHERE age = $1", age)
	// check(err)
	// defer rows.Close()

}