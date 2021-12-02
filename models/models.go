// store the database schema
// The User struct is a representation of users table
package models

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv" // package used to read the .env file
	"log"
	"os"
)

type User struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Location string `json:"location"`
	Age int64 `json:"age"`
}


func CreateTableUsers() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// open the connection
	db, err := sql.Open("postgres", url)
	defer db.Close()

	if err != nil {
		panic(err)
	}

	errPing := db.Ping()

	if errPing != nil {
		panic(err)
	}

	sqlStatement := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS public.users (
		userid serial NOT NULL,
		name text NOT NULL,
		age integer NOT NULL,
		location text NOT NULL,
		CONSTRAINT userspk PRIMARY KEY (userid)
	)`)

	_, errCreating := db.Query(sqlStatement)

	if errCreating != nil {
		log.Fatalf("Unable to execute the query. %v", errCreating)
	}

	return
}