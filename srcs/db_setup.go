// db_setup.go

package main

import (
	"database/sql"
	"fmt"
)

// ! setup only for timestamp-go, will need more work for others
// ? here maybe call setupDbTimestamp, setupDbFccProject2
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

	// Create the table if it doesn't exist
	_, err = db.Exec(`
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
	if err != nil {
		fmt.Println("Create Table", err)
		return err
	}

	fmt.Println("Database and table initialization completed.")
	return nil
}
