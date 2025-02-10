package models

import "time"

type Toy struct {
	ID          uint `json:"id" gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	ReleaseDate time.Time `json:"release_date" gorm:"not null"`
	Price       float64   `json:"price" gorm:"not null;default:0.0"`
	Company     string    `json:"company"`
	ProductType string    `json:"product_type" gorm:"not null"`
	Theme       string    `json:"theme" gorm:"not null"`
	Count       int       `json:"count" gorm:"not null;default:0"`
	Available   bool      `json:"available" gorm:"not null;default:false"`
	Rating      float64   `json:"rating" gorm:"not null;default:0.0"`
	UserId      int       `json:"user_id"`
	User        User      `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}
