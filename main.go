package main

import (
	"log"
	"myapi/db"
	"myapi/routes"
	"net/http"
)

func main() {
    db.Init()

	router := routes.InitializeRoutes()

    log.Fatal(http.ListenAndServe(":8000", router))
}
