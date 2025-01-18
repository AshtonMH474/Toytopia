package models

type Toy struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	ProductType string `json:"product_type" gorm:"not null"`
	Theme       string `json:"theme" gorm:"not null"`
	Count       int    `json:"count" gorm:"not null;default:0"`
	Available   bool   `json:"available" gorm:"not null;default:false"`
}
