package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"tutor/backend/db"
	"tutor/backend/routes"
)

func main() {
	// Initialize the database
	database := db.InitDB()

	// Create the router
	r := mux.NewRouter()
	// Setup routes
	routes.RegisterRoutes(r, database)

	// Define CORS middleware
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Allow all origins, or specify allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	// Wrap the router with CORS middleware
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", corsMiddleware(r)))
}
