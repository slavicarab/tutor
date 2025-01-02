package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tutor/backend/models"

	"gorm.io/gorm"
)

func GetAppointments(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var app []models.Appointment
		if err := db.Find(&app).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	}
}

func CreateAppointment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var app models.Appointment
		if err := json.NewDecoder(r.Body).Decode(&app); err != nil {
			fmt.Println(app)
			log.Printf("Error decoding JSON: %v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := db.Create(&app).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	}
}
