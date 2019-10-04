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

func check(error interface{}) {
	if error != nil {
		log.Fatal(error)
		panic(error)
	}
}

func main() {

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)
	fmt.Println("Database ping succesfull")

	rows, err := db.Query("SELECT employee_id,first_name FROM employee")
	check(err)
	defer rows.Close()
	for rows.Next() {
		var employee_id int
		var first_name string
		err = rows.Scan(&employee_id, &first_name)
		if err != nil {
			// handle this error
			panic(err)
		}
		fmt.Println(employee_id, first_name)
	}

}
