package validator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	model "github.com/lakshay88/golang-task/models"
)

func Validate(c *gin.Context, p model.PersonCreate) error {
	if p.Name == "" {
		return c.Error(fmt.Errorf("name is required"))
	}
	if p.PhoneNumber == "" {
		return c.Error(fmt.Errorf("phone number is required"))
	}
	if p.City == "" {
		return c.Error(fmt.Errorf("city is required"))
	}
	if p.State == "" {
		return c.Error(fmt.Errorf("state is required"))
	}
	if p.Street1 == "" {
		return c.Error(fmt.Errorf("street1 is required"))
	}
	if p.ZipCode == "" {
		return c.Error(fmt.Errorf("zip code is required"))
	}
	return nil
}
