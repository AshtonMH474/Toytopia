package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
)

func SeedToyImages() {
	images := []models.ToyImage{
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/Baby-Yoda-Plush-Stuffed-Toy-Plushie-Cute-Animal-Pillow-Grogu-Baby-Yoda-Star-Wars-Kids-Doll-Gift-10-inch-26-cm_b0514180-f418-49bf-aaae-05fd87088b9f.c00883bd8f234eb2d97ea171a678024f.avif", PrimaryImg: true, ToyId: 8},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/redranger.jpg", PrimaryImg: true, ToyId: 7},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/luke.jpg", PrimaryImg: true, ToyId: 6},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/Radiator-Springs-Lightning-McQueen-Diecast-Car-Disney-Cars_33243daa-d660-4423-a1a2-5cdce7d1cf4d.5079ef90c615abdf4886a7500f76e767.avif", PrimaryImg: true, ToyId: 5},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/BST-AXN_TMNT_IDW_Michelangelo-Battle-Ready_Package_1.webp", PrimaryImg: true, ToyId: 4},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/mellianflacon.jpg", PrimaryImg: true, ToyId: 3},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/buzz.jpg", PrimaryImg: true, ToyId: 2},
		{ImgUrl: "https://toytopiaimages.s3.us-east-2.amazonaws.com/images_toys/optimus.png", PrimaryImg: true, ToyId: 1},
	}
	for _, image := range images {
		if err := database.Database.Db.Create(&image).Error; err != nil {
			log.Printf("Failed to seed image: %v\n", err)
		}
	}
}

func UndoToyImages() {
	if err := database.Database.Db.Exec("DELETE FROM toy_images").Error; err != nil {
		log.Printf("Failed to delete all toy_images: %v\n", err)
	} else {
		log.Println("Successfully deleted all toy_images from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'toy_images'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toy_images table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "toy_images_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for toy_images table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
