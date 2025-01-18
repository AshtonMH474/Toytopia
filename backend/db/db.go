package database

import (
	"fmt"
	"log"

	"github.com/AshtonMH474/Toytopia/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBinstance struct {
	Db *gorm.DB
}

var Database DBinstance

func ConnectDB(cfg config.Config) {
	var err error
	var db *gorm.DB

	if cfg.Environment == "development" {
		db, err = gorm.Open(sqlite.Open(cfg.DbFile), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(cfg.DbURL), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Ensure schema exists in production
	if cfg.Environment == "production" && cfg.Schema != "" {
		err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", cfg.Schema)).Error
		if err != nil {
			log.Fatal("Failed to ensure schema: ", err)
		}
	}

	log.Println("Database connection success!")
	db.Logger = logger.Default.LogMode(logger.Info)

	Database = DBinstance{Db: db}
}
