package seeders

import (
	"log"

	database "github.com/AshtonMH474/Toytopia/db"
	"github.com/AshtonMH474/Toytopia/models"
	"golang.org/x/crypto/bcrypt"
)

func SeedUsers() {
	hash1, err1 := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	hash2, err2 := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	hash3, err3 := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)

	if err1 != nil || err2 != nil || err3 != nil {
		return
	}

	users := []models.User{
		{FirstName: "Demo", LastName: "User", Email: "demo@gmail.com", Password: string(hash1), Username: "DemoUser"},
		{FirstName: "What", LastName: "Ever", Email: "What@gmail.com", Password: string(hash2), Username: "WhatEver"},
		{FirstName: "User2", LastName: "Wow", Email: "Wow@gmail.com", Password: string(hash3), Username: "User2Wow"},
	}

	for _, user := range users {
		if err := database.Database.Db.Create(&user).Error; err != nil {
			log.Printf("Failed to seed user: %v\n", err)
		}
	}
}

func UndoAllUsers() {
	// Delete all records in the users table
	if err := database.Database.Db.Exec("DELETE FROM users").Error; err != nil {
		log.Printf("Failed to delete all users: %v\n", err)
	} else {
		log.Println("Successfully deleted all users from the table.")
	}

	// Reset the auto-increment ID for the products table
	switch database.Database.Db.Dialector.Name() {
	case "sqlite":
		// SQLite-specific reset
		if err := database.Database.Db.Exec("DELETE FROM sqlite_sequence WHERE name = 'users'").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (SQLite): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for users table (SQLite).")
		}
	case "postgres":
		// PostgreSQL-specific reset
		sequenceName := "users_id_seq" // Default naming convention in PostgreSQL
		if err := database.Database.Db.Exec("ALTER SEQUENCE " + sequenceName + " RESTART WITH 1").Error; err != nil {
			log.Printf("Failed to reset auto-increment ID (PostgreSQL): %v\n", err)
		} else {
			log.Println("Successfully reset auto-increment ID for users table (PostgreSQL).")
		}
	default:
		log.Println("Unsupported database type. Auto-increment ID reset skipped.")
	}
}
