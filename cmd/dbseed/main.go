package main

import (
	"log"

	"github.com/prrng/dealls/config"
	"github.com/prrng/dealls/dbase"
	"github.com/prrng/dealls/dbase/seeders"
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
