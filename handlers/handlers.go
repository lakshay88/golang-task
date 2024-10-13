package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lakshay88/golang-task/database"
	"github.com/lakshay88/golang-task/handlers/validator"
	model "github.com/lakshay88/golang-task/models"
)

func GetPersonInfo(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Getting Person Id
		personIDSting := c.Param("person_id")
		personID, err := strconv.Atoi(personIDSting)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid person_id format"})
			return
		}

		// feteching person info -
		personInfo, err := db.FetchPersonInfo(personID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch person info"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name":         personInfo.Name,
			"phone_number": personInfo.PhoneNumber,
			"city":         personInfo.City,
			"state":        personInfo.State,
			"street1":      personInfo.Street1,
			"street2":      personInfo.Street2,
			"zip_code":     personInfo.ZipCode,
		})
	}
}

func CreatePerson(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var person model.PersonCreate

		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid input format",
			})
			return
		}

		if err := validator.Validate(c, person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// creating person db
		err := db.CreatePerson(person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create person",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Person created successfully",
		})
	}
}
