package seeders

import (
	"log"
	"time"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedToys() {
	toys := []models.Toy{
		{ReleaseDate: time.Date(1984, time.January, 18, 0, 0, 0, 0, time.UTC), Price: 1000, ProductType: "Optimus Prime g1 Orginial", Theme: "Transformers", Count: 2, Available: true},
		{ReleaseDate: time.Date(1995, time.December, 15, 0, 0, 0, 0, time.UTC), Price: 250.75, ProductType: "Buzz Lightyear Action Figure", Theme: "Toy Story", Count: 5, Available: true},
		{ReleaseDate: time.Date(2001, time.November, 3, 0, 0, 0, 0, time.UTC), Price: 500, ProductType: "Millennium Falcon", Theme: "Star Wars", Count: 1, Available: true},
		{ReleaseDate: time.Date(1989, time.May, 22, 0, 0, 0, 0, time.UTC), Price: 75, ProductType: "Michelangelo Action Figure", Theme: "Teenage Mutant Ninja Turtles", Count: 12, Available: true},
		{ReleaseDate: time.Date(2010, time.March, 15, 0, 0, 0, 0, time.UTC), Price: 20.99, ProductType: "Lightning McQueen Diecast", Theme: "Cars", Count: 30, Available: true},
		{ReleaseDate: time.Date(1978, time.October, 17, 0, 0, 0, 0, time.UTC), Price: 1000, ProductType: "Luke Skywalker Original Action Figure", Theme: "Star Wars", Count: 3, Available: true},
		{ReleaseDate: time.Date(1998, time.June, 10, 0, 0, 0, 0, time.UTC), Price: 150, ProductType: "Red Ranger Figurine", Theme: "Power Rangers", Count: 8, Available: true},
		{ReleaseDate: time.Date(2020, time.February, 1, 0, 0, 0, 0, time.UTC), Price: 35.5, ProductType: "Baby Yoda Plush", Theme: "The Mandalorian", Count: 50, Available: true},
	}

	for _, toy := range toys {
		if err := database.Database.Db.Create(&toy).Error; err != nil {
			log.Printf("Failed to seed Toy: %v\n", err)
		}
	}
}

func UndoToys() {
	// Delete all records in the toys table
	if err := database.Database.Db.Exec("DELETE FROM toys").Error; err != nil {
		log.Printf("Failed to delete all toys: %v\n", err)
	} else {
		log.Println("Successfully deleted all toys from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'toys'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "toys_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toys table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
