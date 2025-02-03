package main

import (
	"log"
	"net/http"
	"scheduling/database"
	"scheduling/routes"
)

func main() {
	uri := "mongodb+srv://jbiazonferreira:qwerty123456@cluster0.82ixr.mongodb.net/"

	_, err := database.InitDB(uri)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	router := routes.ScheduleRoutes()

	log.Println("Starting server on :3001")
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
