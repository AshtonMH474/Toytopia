package routes

import (
	"time"

	"github.com/AshtonMH474/Toytopia/models"
)

type ToySerial struct {
	// not model Toy, see this as serialzer
	ID          uint       `json:"id"`
	ReleaseDate time.Time  `json:"release_date"`
	Price       float64    `json:"price"`
	ProductType string     `json:"product_type"`
	Theme       string     `json:"theme"`
	Count       int        `json:"count"`
	Available   bool       `json:"available"`
	User        UserSerial `json:"user"`
}

func CreateResToy(toy models.Toy, user UserSerial) ToySerial {
	return ToySerial{ID: toy.ID, ReleaseDate: toy.ReleaseDate, Price: toy.Price, ProductType: toy.ProductType, Theme: toy.Theme, Count: toy.Count, Available: toy.Available, User: user}
}
