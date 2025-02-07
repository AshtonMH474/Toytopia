package models

import "time"

type ToyImage struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	PrimaryImg bool   `json:"primary_img"`
	ImgUrl     string `json:"img_url" gorm:"not null"`
	ToyId      int    `json:"toy_id"`
	Toy        Toy    `gorm:"foreignKey:ToyId;constraint:OnDelete:CASCADE;"`
}
