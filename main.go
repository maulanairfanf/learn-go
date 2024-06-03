package main

import (
	"log"
	"myapi/db"
	"myapi/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    db.Init()

    r := mux.NewRouter()
    r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
    
    r.HandleFunc("/products", handlers.GetProducts).Methods("GET")
    r.HandleFunc("/products/{id}", handlers.GetProduct).Methods("GET")
    r.HandleFunc("/products", handlers.CreateProduct).Methods("POST")

    r.HandleFunc("/category", handlers.CreateCategory).Methods("POST")
    r.HandleFunc("/category", handlers.GetCategories).Methods("GET")
    r.HandleFunc("/category/{id}", handlers.GetCategories).Methods("GET")

    log.Fatal(http.ListenAndServe(":8000", r))
}
