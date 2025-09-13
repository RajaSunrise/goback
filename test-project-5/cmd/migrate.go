package main

import (
	"fmt"
	"log"
	"os"

	"github.com/test/test-project-5/internal/config"
	"github.com/test/test-project-5/internal/database"
	"github.com/test/test-project-5/internal/models"
	"gorm.io/gorm"
)

func main() {
	// Load configuration from .env
	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	// Connect to the database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	// Determine migration direction from command-line arguments
	direction := "up"
	if len(os.Args) > 1 && os.Args[1] == "down" {
		direction = "down"
	}

	if direction == "up" {
		log.Println("Running migrations...")
		// GORM's AutoMigrate handles creating/updating tables.
		// Add all your models here.
		err = db.AutoMigrate(
			&models.User{},
			// &models.Product{}, // Add other models here
		)
		if err != nil {
			log.Fatalf("migration failed: %v", err)
		}
		log.Println("Migrations completed successfully.")
	} else {
		log.Println("Rolling back migrations...")
		// GORM doesn't have a built-in down migration.
		// You drop tables manually for a full rollback.
		err = db.Migrator().DropTable(
			&models.User{},
			// &models.Product{}, // Add other models here
		)
		if err != nil {
			log.Fatalf("rollback failed: %v", err)
		}
		log.Println("Rollback completed successfully.")
	}
}
