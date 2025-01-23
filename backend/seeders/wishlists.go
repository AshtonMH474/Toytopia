package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedWishlists() {
	wishslists := []models.Wishlist{
		{UserId: 1, Name: "Originals", Description: "All the retro toys"},
		{UserId: 1, Name: "Transformers", Description: "Transformers can transform"},
		{UserId: 3, Name: "I want", Description: "Random"},
	}

	for _, wishlist := range wishslists {
		if err := database.Database.Db.Create(&wishlist).Error; err != nil {
			log.Printf("Failed to seed wishlist: %v\n", err)
		}
	}

	toys_in_wishlists := []models.ToysInWishlist{
		{ToyId: 1, WishlistId: 1},
		{ToyId: 1, WishlistId: 2},
		{ToyId: 4, WishlistId: 1},
		{ToyId: 6, WishlistId: 1},
	}

	for _, toy := range toys_in_wishlists {
		if err := database.Database.Db.Create(&toy).Error; err != nil {
			log.Printf("Failed to seed toys in wishlists: %v\n", err)
		}
	}
}

func UndoWishlistsToys() {

	if err := database.Database.Db.Exec("DELETE FROM toys_in_wishlists").Error; err != nil {
		log.Printf("Failed to delete all toys_in_wishlists: %v\n", err)
	} else {
		log.Println("Successfully deleted all toys_in_wishlists from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'toys_in_wishlists'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys_in_wishlists table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "toys_in_wishlists_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys_in_wishlists table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
	// Delete all records in the wishlist table

}

func UndoAllWishLists() {
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
