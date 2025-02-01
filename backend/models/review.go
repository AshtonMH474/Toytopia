package models

import "time"

type Review struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Rating    float64 `json:"rating" gorm:"not null;default:0.0"`
	Review    string  `json:"review" gorm:"not null"`
	UserId    int     `json:"user_id"`
	User      User    `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}
