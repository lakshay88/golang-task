package routers

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lakshay88/golang-task/config"
	"github.com/lakshay88/golang-task/database"
	"github.com/lakshay88/golang-task/handlers"
)

func RouterRegister(cfg config.AppConfig, db database.Database) error {
	router := gin.Default()

	// Add routs
	router.GET("/person/:person_id/info", handlers.GetPersonInfo(db))
	router.POST("/person/create", handlers.CreatePerson(db))

	log.Printf("REST Server Started running on port: %s", cfg.ServerConfig.Port)

	if err := router.Run(fmt.Sprintf(":%s", cfg.ServerConfig.Port)); err != nil {
		log.Fatalf("Failed to start REST Server", err)
		return err
	}

	return nil
}
