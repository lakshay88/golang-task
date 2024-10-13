package database

import model "github.com/lakshay88/golang-task/models"

type Database interface {
	FetchPersonInfo(personID int) (*model.PersonInfo, error)
	CreatePerson(person model.PersonCreate) error
	Close() error
}
