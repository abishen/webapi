package webapi

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection and runs migrations
func InitDB(dbPath string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate the Album model
	if err := DB.AutoMigrate(&Album{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Seed initial data if table is empty
	if err := seedInitialData(); err != nil {
		log.Printf("warning: failed to seed initial data: %v", err)
	}

	return nil
}

// seedInitialData populates the database with initial albums
func seedInitialData() error {
	var count int64
	DB.Model(&Album{}).Count(&count)

	// Only seed if table is empty
	if count > 0 {
		return nil
	}

	initialAlbums := []Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	}

	return DB.Create(&initialAlbums).Error
}

// CloseDB closes the database connection
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
