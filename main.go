package main

import (
	"log"

	"github.com/lakshay88/golang-task/config"
	"github.com/lakshay88/golang-task/database"
	"github.com/lakshay88/golang-task/database/mysql"
	"github.com/lakshay88/golang-task/routers"
)

func main() {

	log.Println("Setting started")
	// Getting Configuration data
	cfg, err := config.LoadConfiguration("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		return
	}

	// Creating data base connection
	log.Println("Setting Database connection")
	var db database.Database
	switch cfg.Database.Driver {
	case "mysql":
		db, err = mysql.ConnectionToMySQL(cfg.Database)
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
	}

	// Closing Connection grace full shut down
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Failed closing database connection: %v", err)
		}
	}()

	// Establish router connection
	log.Println("Setting Routers")
	if err = routers.RouterRegister(*cfg, db); err != nil {
		log.Fatalf("Failed to Resgister routs: %v", err)
		return
	}

	log.Println("Setting Completed")
}
