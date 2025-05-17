package api

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Step-by-step comments for DBConnection:
// Step 1: Define the connection parameters including username, password, database name, and host.
// Step 2: Construct the Data Source Name (DSN) using the connection parameters.
// Step 3: Establish the database connection using the sql.Open function with the MySQL driver.
// Step 4: If the connection fails, log an error message and return the error.
// Step 5: Set the maximum number of open connections to the database to 20 to manage connection pooling.
// Step 6: Set the maximum number of idle connections to 10 to optimize resource usage.
// Step 7: Set the maximum time (in seconds) that a connection can remain idle before being closed to 60 seconds.
// Step 8: Check if the connection is successful by pinging the database.
// Step 9: If the ping fails, log an error message and return the error.
// Step 10: Return the established database connection.

func DBConnection() (*sql.DB, error) {
	// Define the connection parameters
	lDBUser := "your_username"
	lDBPassword := "your_password"
	lDBName := "your_database_name"
	lDBHost := "localhost" // or your database host

	// Construct the Data Source Name (DSN)
	lDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s", lDBUser, lDBPassword, lDBHost, lDBName)

	// Establish the database connection
	lDB, lErr := sql.Open("mysql", lDSN)
	if lErr != nil {
		// Log an error message if the connection fails
		log.Println("DBConnection - Error connecting to the database:", lErr)
		return nil, lErr
	}

	// Set the maximum number of open connections to the database
	lDB.SetMaxOpenConns(20)

	// Set the maximum number of idle connections to the database
	lDB.SetMaxIdleConns(10)

	// Set the maximum time (in seconds) that a connection can remain idle before being closed
	lDB.SetConnMaxIdleTime(60 * time.Second)

	// Check if the connection is successful
	lErr = lDB.Ping()
	if lErr != nil {
		// Log an error message if the ping fails
		log.Println("DBConnection - Error pinging database:", lErr)
		return nil, lErr
	}

	// Return the established database connection
	return lDB, nil
}
