package database

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBInit() error {
	var err error

	// connect to DB
	if DB, err = gorm.Open(sqlite.Open("Shopify-Challenge.db"), &gorm.Config{}); err != nil {
		return err
	}

	// initialize schema from go structs
	err = DB.AutoMigrate(&Inventory{}, &Item{}, &Shipment{})
	return err
}
