package middleware

import (
	"database/sql" // "encoding/json" // package to encode and decode the json into struct and vice versa
	"fmt"          // "github.com/gorilla/mux" // used to get the params from the route
	"log"          // used to access the request and response object of the api
	"os"           // used to read the environment variable

	"github.com/joho/godotenv" // package used to read the .env file
	_ "github.com/lib/pq"      // postgres golang driver
	// models package where User schema is defined
	// "strconv" // package used to covert string into int type
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres database
func createConnection() *sql.DB {
	// load .env file

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// open the connection
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// return the connection
	return db
}
