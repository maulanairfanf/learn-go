package main

import (
	"log"
	"myapi/db"
	"myapi/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()

    // Initialize the database connection
    db.InitDB("root:@tcp(127.0.0.1:3306)/learn_go")

    // Define routes
    router.HandleFunc("/users", handlers.GetUsers).Methods("GET")

    // Start the server
    log.Fatal(http.ListenAndServe(":8000", router))
}
