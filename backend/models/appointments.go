package models

import (
	"gorm.io/gorm"
)

type Appointment struct {
	ID      uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primaryKey"`
	AppDate string `json:"date"`
	AppTime string `json:"time"`
	UserID  uint   `json:"userId" gorm:"not null"`
	User    Users  `json:"user" gorm:"foreignKey:UserID"`
}

// Save saves the Appointment instance to the database
func (a *Appointment) Save(db *gorm.DB) error {
	return db.Create(a).Error
}

// Update updates the Appointment instance in the database
func (a *Appointment) Update(db *gorm.DB) error {
	return db.Save(a).Error
}

// Delete deletes the Appointment instance from the database
func (a *Appointment) Delete(db *gorm.DB) error {
	return db.Delete(a).Error
}

// FindByDate retrieves a Appointemtns by date
//func (a *Appointment) FindByDate(db *gorm.DB, string appDate) error {
//	return db.Where(a.AppDate, " = ", appDate).Find(a).Error
//}
