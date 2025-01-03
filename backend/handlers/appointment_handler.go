package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"tutor/backend/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetAppointments(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var app []models.Appointment
		if err := db.Preload("User").Find(&app).Error; err != nil {
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

// Update a Appointment
func UpdateAppointment(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var app models.Appointment
		if err := db.First(&app, id).Error; err != nil {
			http.Error(w, "Appointments not found", http.StatusNotFound)
			return
		}

		// Decode the updated data
		var updatedApp models.Appointment
		if err := json.NewDecoder(r.Body).Decode(&updatedApp); err != nil {
			fmt.Println(updatedApp)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Update the fields
		app.AppDate = updatedApp.AppDate
		app.AppTime = updatedApp.AppTime
		app.UserID = updatedApp.UserID

		// Save the changes to the database
		if err := db.Save(&app).Error; err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	}
}

// Get Appointments user by DATE
func GetAppointmentsByDate(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		appDate := vars["date"]

		var app []models.Appointment
		//fmt.Println(db.Debug().Preload("User").Where("app_date = ?", appDate).Find(&app))
		if err := db.Preload("User").Where("app_date = ?", appDate).Find(&app).Error; err != nil {
			http.Error(w, "Appointmets not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	}
}

// Get a single appointment by ID
func GetAppointmentByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var app models.Appointment
		if err := db.Preload("User").First(&app, id).Error; err != nil {
			http.Error(w, "Appointment not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	}
}

// Delete a user by ID
func DeleteAppointmentByID(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		if err := db.Delete(&models.Appointment{}, id).Error; err != nil {
			http.Error(w, "Appointment not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
