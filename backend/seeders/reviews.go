package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedReviews() {
	reviews := []models.Review{
		{Review: "Sucky", Rating: 1.1, UserId: 1},
		{Review: "Amazing", Rating: 4.8, UserId: 1},
		{Review: "Not Bad", Rating: 3.0, UserId: 1},
	}

	for _, review := range reviews {
		if err := database.Database.Db.Create(&review).Error; err != nil {
			log.Printf("Failed to seed Review: %v\n", err)
		}
	}
}

func UndoReviews() {
	// Delete all records in the toys table
	if err := database.Database.Db.Exec("DELETE FROM reviews").Error; err != nil {
		log.Printf("Failed to delete all reviews: %v\n", err)
	} else {
		log.Println("Successfully deleted all reviews from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'reviews'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "reviews_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for reviews table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
