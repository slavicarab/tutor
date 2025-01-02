package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	ID          uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	UserName    string `json:"userName"`
	UserClass   string `json:"userClass"`
	UserAddress string `json:"userAddress"`
	UserNumber  string `json:"userNumber"`
	UserEmail   string `json:"userEmail"`
	UserCourse  string `json:"userCourse"`
	UserStatus  bool   `json:"userStatus"`
}

var db *gorm.DB

func initDB() {
	dsn := "tutor:tutor@tcp(127.0.0.1:3306)/tutor?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err := db.AutoMigrate(&Users{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	var users []Users
	if err := db.Find(&users).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user Users
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Println(user)
		log.Printf("Error decoding JSON: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := db.Create(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Update a Todo
func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user Users
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Decode the updated data
	var updatedUser Users
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		fmt.Println(updatedUser)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the fields
	user.UserName = updatedUser.UserName
	user.UserClass = updatedUser.UserClass
	user.UserAddress = updatedUser.UserAddress
	user.UserEmail = updatedUser.UserEmail
	user.UserNumber = updatedUser.UserNumber
	user.UserCourse = updatedUser.UserCourse
	user.UserStatus = updatedUser.UserStatus

	// Save the changes to the database
	if err := db.Save(&user).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Get a single user by ID
func getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var user Users
	if err := db.First(&user, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Delete a user by ID
func deleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := db.Delete(&Users{}, id).Error; err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	initDB()

	// Create the router
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", updateUser).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", getUserByID).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", deleteUserByID).Methods("DELETE")

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
