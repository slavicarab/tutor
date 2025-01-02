package models

import "gorm.io/gorm"

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

// Save saves the user instance to the database
func (u *Users) Save(db *gorm.DB) error {
	return db.Create(u).Error
}

// Update updates the user instance in the database
func (u *Users) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

// Delete deletes the user instance from the database
func (u *Users) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}

// FindByID retrieves a user by ID
func (u *Users) FindByID(db *gorm.DB, id uint) error {
	return db.First(u, id).Error
}
