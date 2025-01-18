package models

import "time"

type User struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"not null"`
	Username  string `json:"username" gorm:"not null"`
}
