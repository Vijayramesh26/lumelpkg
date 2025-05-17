package api

import (
	"database/sql"
	"fmt"
	"log"
	"lumelpkg/common"
)

// Step-by-step comments for SelectRecordsMethod:
// Step 1: Log the start of the function to track its execution.
// Step 2: Establish a connection to the database using a predefined connection key.
// Step 3: If the connection fails, log the error and return it with a specific error code.
// Step 4: Ensure the database connection is closed when the function exits to avoid leaks.
// Step 5: Prepare an SQL query to select records from the database.
// Step 6: If preparing the SQL statement fails, log the error and return it.
// Step 7: Ensure the prepared statement is closed after its use to free resources.
// Step 8: Execute the prepared SQL query with the given parameter.
// Step 9: If the query execution fails, log the error and return it.
// Step 10: Ensure the result set is closed after its use to free resources.
// Step 11: Process the result set by iterating over it.
// Step 12: For each row, scan the result into a variable and handle any scanning errors.
// Step 13: Log the end of the function.
// Step 14: Return the result and any error encountered.

// This method is used to fetch data for this purpose from this table.
func SelectRecordsMethod(pParameterName string) (lVariable string, lErr error) {
	// Log the start of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (+)")

	// Connect to the database using a connection key
	lDb, lErr := DBConnection()
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("ASRM:001", lErr.Error())
		// Return the local variable and an error with a specific code and the error message
		return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-001) " + lErr.Error())
	} else {
		// Ensure the database connection is closed when the function exits
		defer lDb.Close()

		// Prepare the SQL statement to retrieve data from the table (replace with actual query)
		lCoreString := `// Enter Select Query`
		lStmt, lErr := lDb.Prepare(lCoreString)
		if lErr != nil {
			// Log an error message if the statement preparation fails
			log.Println("ASRM:002", lErr.Error())
			// Return the local variable and an error with a specific code and the error message
			return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-002) " + lErr.Error())
		} else {
			// Ensure the prepared statement is closed when the function exits
			defer lStmt.Close()

			// Execute the prepared statement with the provided parameter
			lRows, lErr := lStmt.Query(pParameterName)
			if lErr != nil {
				// Log an error message if the query execution fails
				log.Println("ASRM:003", lErr.Error())
				// Return the local variable and an error with a specific code and the error message
				return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-003) " + lErr.Error())
			} else {
				// Ensure the result set is closed when the function exits
				defer lRows.Close()

				// Process the result set
				for lRows.Next() {
					// Scan the row and store the result in the local variable
					lErr := lRows.Scan(&lVariable)
					if lErr != nil {
						// Log an error message if scanning the row fails
						log.Println("ASRM:004", lErr.Error())
						// Return the local variable and an error with a specific code and the error message
						return lVariable, fmt.Errorf("SelectRecordsMethod - (ASRM-004) " + lErr.Error())
					} else {
						// Additional logic for processing the result can be added here
					}
				}
			}
		}
	}

	// Log the end of the SelectRecordsMethod function
	log.Println("SelectRecordsMethod (-)")

	// Return the local variable and no error
	return lVariable, nil
}

// Step-by-step comments for InsertUpdateMethod:
// Step 1: Log the start of the function to track its execution.
// Step 2: Establish a connection to the database using a predefined connection key.
// Step 3: If the connection fails, log the error and return it with a specific error code.
// Step 4: Ensure the database connection is closed when the function exits to avoid leaks.
// Step 5: Initialize variables for the SQL query and execution result.
// Step 6: Determine the SQL query based on the provided flag (INSERT or UPDATE).
// Step 7: If the flag is INSERT, prepare and execute the insert query with the given parameter.
// Step 8: If the flag is UPDATE, prepare and execute the update query with the given parameter.
// Step 9: If executing the query fails, log the error and return it with a specific error code.
// Step 10: Retrieve and log the number of rows affected by the query.
// Step 11: If retrieving the affected rows count fails, log the error.
// Step 12: Log the end of the function and return any error encountered.

// This method is used to insert or update records based on the flag.
func InsertUpdateMethod(pParameterName, pFlag string) error {
	// Log the start of the InsertUpdateMethod function
	log.Println("InsertUpdateMethod (+)")

	// Establish a connection to the database using a connection key
	lDb, lErr := DBConnection()
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("AIUM-001 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertUpdateMethod - (AIUM-001) " + lErr.Error())
	}
	// Ensure the database connection is closed when the function exits
	defer lDb.Close()

	var lCorestring string
	var lExecResult sql.Result

	// Prepare the SQL statement based on the flag (INSERT or UPDATE)
	switch {
	case pFlag == common.INSERT:
		// Define the SQL insert query string (replace with actual query)
		lCorestring = `Enter Insert Query`
		// Execute the SQL insert query with the provided parameter
		lExecResult, lErr = lDb.Exec(lCorestring, pParameterName)
	case pFlag == common.UPDATE:
		// Define the SQL update query string (replace with actual query)
		lCorestring = `Enter Update Query`
		// Execute the SQL update query with the provided parameter
		lExecResult, lErr = lDb.Exec(lCorestring, pParameterName)
	}

	// Check if there was an error executing the query
	if lErr != nil {
		// Log an error message if the query execution fails
		log.Println("AIUM-002 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertUpdateMethod - (AIUM-002) " + lErr.Error())
	} else {
		// Check the number of rows affected by the insert or update query
		lRowsAffected, _ := lExecResult.RowsAffected()
		if lErr != nil {
			// Log an error message if fetching the affected rows count fails
			log.Println("AIUM-003 ", lErr.Error())
		} else {
			// Log the number of rows affected and a success message
			log.Printf("InsertUpdateMethod Rows affected: %d\n", lRowsAffected)
			log.Println("Record Inserted or Updated successfully")
		}
	}

	// Log the end of the InsertUpdateMethod function
	log.Println("InsertUpdateMethod (-)")
	return nil
}

// Step-by-step comments for InsertRecords:
// Step 1: Log the start of the function to track its execution.
// Step 2: Establish a connection to the database using a predefined connection key.
// Step 3: If the connection fails, log the error and return it with a specific error code.
// Step 4: Ensure the database connection is closed when the function exits to avoid leaks.
// Step 5: Prepare the SQL insert query string (replace with actual query).
// Step 6: Execute the SQL insert query with the given parameter.
// Step 7: If executing the query fails, log the error and return it with a specific error code.
// Step 8: Retrieve and log the number of rows affected by the insert query.
// Step 9: If retrieving the affected rows count fails, log the error.
// Step 10: Log the end of the function and return any error encountered.

// This method is used to insert records into the database.
func InsertRecords(pParameterName string) error {
	// Log the start of the InsertRecords function
	log.Println("InsertRecords (+)")

	// Establish a connection to the database using a connection key
	lDb, lErr := DBConnection()
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("AIR-001 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("InsertRecords - (AIR-001) " + lErr.Error())
	} else {
		// Ensure the database connection is closed when the function exits
		defer lDb.Close()

		// Define the SQL insert query string (replace with actual query)
		lSqlString := `//Enter Insert Query`

		// Execute the SQL insert query with the provided parameter
		lExecResult, lErr := lDb.Exec(lSqlString, pParameterName)
		if lErr != nil {
			// Log an error message if the query execution fails
			log.Println("AIR-002 ", lErr.Error())
			// Return an error with a specific code and the error message
			return fmt.Errorf("InsertRecords - (AIR-002) " + lErr.Error())
		} else {
			// Check the number of rows affected by the insert query
			lRowsAffected, lErr := lExecResult.RowsAffected()
			if lErr != nil {
				// Log an error message if fetching the affected rows count fails
				log.Println("AIR-003 ", lErr.Error())
			} else {
				// Log the number of rows affected and a success message
				log.Printf("InsertRecords Rows affected: %d\n", lRowsAffected)
				log.Println("Record Inserted successfully")
			}
		}
	}

	// Log the end of the InsertRecords function
	log.Println("InsertRecords (-)")
	return nil
}

// Step-by-step comments for UpdateRecords:
// Step 1: Log the start of the function to track its execution.
// Step 2: Establish a connection to the database using a predefined connection key.
// Step 3: If the connection fails, log the error and return it with a specific error code.
// Step 4: Ensure the database connection is closed when the function exits to avoid leaks.
// Step 5: Prepare the SQL update query string (replace with actual query).
// Step 6: Execute the SQL update query with the given parameter.
// Step 7: If executing the query fails, log the error and return it with a specific error code.
// Step 8: Retrieve and log the number of rows affected by the update query.
// Step 9: If retrieving the affected rows count fails, log the error.
// Step 10: Log the end of the function and return any error encountered.

func UpdateRecords(pParameterName string) error {
	// Log the start of the UpdateRecords function
	log.Println("UpdateRecords (+)")

	// Establish a connection to the database using a connection key
	lDb, lErr := DBConnection()
	if lErr != nil {
		// Log an error message if the database connection fails
		log.Println("AUR-001 ", lErr.Error())
		// Return an error with a specific code and the error message
		return fmt.Errorf("UpdateRecords - (AUR-001) " + lErr.Error())
	} else {
		// Ensure the database connection is closed when the function exits
		defer lDb.Close()

		// Define the SQL update query string (replace with actual query)
		lSqlString := `//Enter Update Query`

		// Execute the SQL update query with the provided parameter
		lExecResult, lErr := lDb.Exec(lSqlString, pParameterName)
		if lErr != nil {
			// Log an error message if the query execution fails
			log.Println("AUR-002 ", lErr.Error())
			// Return an error with a specific code and the error message
			return fmt.Errorf("UpdateRecords - (AUR-002) " + lErr.Error())
		} else {
			// Check the number of rows affected by the update query
			lRowsAffected, lErr := lExecResult.RowsAffected()
			if lErr != nil {
				// Log an error message if fetching the affected rows count fails
				log.Println("AUR-003 ", lErr.Error())
			} else {
				// Log the number of rows affected and a success message
				log.Printf("UpdateRecords Rows affected: %d\n", lRowsAffected)
				log.Println("Record Updated successfully")
			}
		}
	}

	// Log the end of the UpdateRecords function
	log.Println("UpdateRecords (-)")
	return nil
}
