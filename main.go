package main

import (
    "fmt"
    "go-postgres/models"
    "go-postgres/router"
    "log"
    "net/http"
)

func main() {
    models.CreateTableUsers()
    r := router.Router()
    // fs := http.FileServer(http.Dir("build"))
    // http.Handle("/", fs)
    fmt.Println("Starting server on the port 8080...")

    log.Fatal(http.ListenAndServe(":8080", r))
}