package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedWishlists() {
	wishslists := []models.Wishlist{
		{UserId: 1, ToyId: 1},
		{UserId: 1, ToyId: 3},
		{UserId: 1, ToyId: 6},
		{UserId: 2, ToyId: 1},
		{UserId: 2, ToyId: 3},
	}

	for _, wishlist := range wishslists {
		if err := database.Database.Db.Create(&wishlist).Error; err != nil {
			log.Printf("Failed to seed wishlist: %v\n", err)
		}
	}
}

func UndoWishlists() {
	// Delete all records in the users table
	if err := database.Database.Db.Exec("DELETE FROM wishlists").Error; err != nil {
		log.Printf("Failed to delete all wishlists: %v\n", err)
	} else {
		log.Println("Successfully deleted all wishlists from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'wishlists'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for wishlists table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "wishlists_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for wishlists table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
