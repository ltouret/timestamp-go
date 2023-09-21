// db_setup.go

package main

import (
	"database/sql"
	"fmt"
)

// Add dbs for urlShortener
// we need one for analytics and one for the service itself
// urlShortener db :
// original Url
// uuid of short Url

// analytics db
func setupTimestampDB(db *sql.DB) error {
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS timestamp (
            id INT AUTO_INCREMENT PRIMARY KEY,
            userAgent VARCHAR(255) DEFAULT NULL,
            clientIP  CHAR(39) DEFAULT NULL,
            timestamp CHAR(36) DEFAULT NULL,
            responseTime CHAR(30) DEFAULT NULL,
            queryParameters VARCHAR(255) DEFAULT NULL,
            statusCode INT DEFAULT NULL,
            userUuid CHAR(36) DEFAULT NULL
        )
    `)
	return err
}

// analytics db
func setupHeaderParserDB(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS headerParser (
			id INT AUTO_INCREMENT PRIMARY KEY,
			userAgent VARCHAR(255) DEFAULT NULL,
			clientIP  CHAR(39) DEFAULT NULL,
            timestamp CHAR(36) DEFAULT NULL,
			language CHAR(39) DEFAULT NULL,
            responseTime CHAR(30) DEFAULT NULL,
			userUuid CHAR(36) DEFAULT NULL
		)
	`)
	return err
}

func SetupUrlShortenerDb(db *sql.DB) error {
	// Check if the database exists
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS urlShortener")
	if err != nil {
		fmt.Println("Create Db", err)
		return err
	}

	// Switch to the newly created database
	_, err = db.Exec("USE urlShortener")
	if err != nil {
		fmt.Println("Use Db", err)
		return err
	}

	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS headerParser (
		id INT AUTO_INCREMENT PRIMARY KEY,
		userAgent VARCHAR(255) DEFAULT NULL,
		clientIP  CHAR(39) DEFAULT NULL,
		timestamp CHAR(36) DEFAULT NULL,
		language CHAR(39) DEFAULT NULL,
		responseTime CHAR(30) DEFAULT NULL,
		userUuid CHAR(36) DEFAULT NULL
	)
	`)
	if err != nil {
		fmt.Println("Create Table", err)
		return err
	}
	return nil
}

func SetupDb(db *sql.DB) error {
	// Check if the database exists
	_, err := db.Exec("CREATE DATABASE IF NOT EXISTS analytics")
	if err != nil {
		fmt.Println("Create Db", err)
		return err
	}

	// Switch to the newly created database
	_, err = db.Exec("USE analytics")
	if err != nil {
		fmt.Println("Use Db", err)
		return err
	}

	// Create the tables if they dont exist
	err = setupTimestampDB(db)
	if err != nil {
		fmt.Println("Create Table", err)
		return err
	}

	err = setupHeaderParserDB(db)
	if err != nil {
		fmt.Println("Create Table", err)
		return err
	}

	fmt.Println("Database and table initialization completed.")
	return nil
}
