package models

import "time"

type Wishlist struct {
	ID        uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	ToyId     int  `json:"toy_id"`
	Toy       Toy  `gorm:"foreignKey:ToyId;constraint:OnDelete:CASCADE;"`
	UserId    int  `json:"user_id"`
	User      User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
}
