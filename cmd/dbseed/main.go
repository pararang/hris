package main

import (
	"log"

	"github.com/pararang/hris/config"
	"github.com/pararang/hris/dbase"
	"github.com/pararang/hris/dbase/seeders"
)

func main() {

	dbConf := config.New().DB

	// Initialize database connection
	db, err := dbase.NewPostgresDB(dbConf)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := seeders.SeedUser(db); err != nil {
		log.Fatalf("Failed to seed data: %v", err)
	}

	log.Println("Data seeding completed successfully!")
}
