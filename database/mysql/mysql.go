package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lakshay88/golang-task/config"
	"github.com/lakshay88/golang-task/database"
	model "github.com/lakshay88/golang-task/models"
)

type MySQLDB struct {
	connection *sql.DB
}

func (db *MySQLDB) FetchPersonInfo(personID int) (*model.PersonInfo, error) {
	var personInfo model.PersonInfo

	query := `
		SELECT p.name, ph.number, a.city, a.state, a.street1, a.street2, a.zip_code
		FROM person p
		INNER JOIN phone ph ON p.id = ph.person_id
		INNER JOIN address_join aj ON p.id = aj.person_id
		INNER JOIN address a ON aj.address_id = a.id
		WHERE p.id = ?;
	`
	row := db.connection.QueryRow(query, personID)

	// Appending data
	err := row.Scan(
		&personInfo.Name,
		&personInfo.PhoneNumber,
		&personInfo.City,
		&personInfo.State,
		&personInfo.Street1,
		&personInfo.Street2,
		&personInfo.ZipCode,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no person found ID %d", personID)
		}
		return nil, err
	}

	return &personInfo, nil
}

func (db *MySQLDB) CreatePerson(person model.PersonCreate) error {

	// Creating One transation connection
	tx, err := db.connection.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// First Person
	personInsertQuery := "INSERT INTO person (name, age) VALUES (?, ?)"
	result, err := tx.Exec(personInsertQuery, person.Name, person.Age)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert person: %w", err)
	}

	personID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retriev last insert ID: %w", err)
	}

	// Second Phone number
	phoneInsertQuery := "INSERT INTO phone (number, person_id) VALUES (?, ?)"
	_, err = tx.Exec(phoneInsertQuery, person.PhoneNumber, personID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert phone number: %w", err)
	}

	// Third Address
	addressInsertQuery := "INSERT INTO address (city, state, street1, street2, zip_code) VALUES (?, ?, ?, ?, ?)"
	addressResult, err := tx.Exec(addressInsertQuery, person.City, person.State, person.Street1, person.Street2, person.ZipCode)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert address: %w", err)
	}

	// Get the inserted address's ID
	addressID, err := addressResult.LastInsertId()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to retrieve address insert ID: %w", err)
	}

	// Join address to person
	addressJoinInsertQuery := "INSERT INTO address_join (person_id, address_id) VALUES (?, ?)"
	_, err = tx.Exec(addressJoinInsertQuery, personID, addressID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into address_join: %w", err)
	}

	// checking transation status
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func ConnectionToMySQL(cfg config.DatabaseConfig) (database.Database, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)

	connection, err := sql.Open(cfg.Driver, connectionString)
	if err != nil {
		log.Fatalf("Failed to connect with db error -> %s", err)
		return nil, fmt.Errorf("failed to open connection: %w", err)
	}

	// Adding Connection Pooling
	connection.SetMaxOpenConns(5)
	connection.SetMaxIdleConns(5)
	connection.SetConnMaxLifetime(5 * time.Minute)

	if err := connection.Ping(); err != nil {
		log.Fatalf("Failed to ping database error : %s", err)
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &MySQLDB{connection: connection}, nil
}

func (db *MySQLDB) Close() error {
	return db.connection.Close()
}
