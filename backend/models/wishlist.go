package models

import "time"

type Wishlist struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserId      int   `json:"user_id" gorm:"not null"`
	User        User  `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE;"`
	Toys        []Toy `gorm:"many2many:toys_in_wishlists;" json:"toys"`
}
