package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"tutor/backend/handlers"
)

func RegisterRoutes(r *mux.Router, db *gorm.DB) {
	r.HandleFunc("/users", handlers.GetUsers(db)).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser(db)).Methods("POST")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.GetUserByID(db)).Methods("GET")
	r.HandleFunc("/users/{id:[0-9]+}", handlers.DeleteUserByID(db)).Methods("DELETE")

	r.HandleFunc("/appointment", handlers.GetAppointments(db)).Methods("GET")
	r.HandleFunc("/appointment", handlers.CreateAppointment(db)).Methods("POST")
}
