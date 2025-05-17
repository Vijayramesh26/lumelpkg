package db

import (
	"database/sql"
	"fmt"
	"lumelpkg/common"
	"lumelpkg/config"
	"lumelpkg/utils"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// DatabaseType holds individual DB connection details.
type DatabaseType struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
	DBType   string // Database driver type, e.g., "mssql", "mysql", "postgres"
	DB       string // Logical DB name identifier
}
type DBConnectionPool struct {
	DbConMaxIdleTime  int
	DbConMaxOpenConns int
	DbConMaxIdleConns int
}

// AllUsedDatabases groups all database configurations used by the program.
type AllUsedDatabases struct {
	DbName DatabaseType
}

// Global_DB_Instance is a global variable holding the active DB connection.
var Global_DB_Instance *sql.DB

// LocalDbConnect opens a database connection based on the requested DB type.
// It reads DB connection limits from config, sets up connection pooling parameters,
// and returns the opened *sql.DB or an error.
// Logs detailed debug and error information using the custom logger.
func LocalDbConnect(pDbName string) (*sql.DB, error) {
	// Initialize logger with a unique request ID for tracing
	log := new(utils.Logger)
	log.SetReqID()
	log.Log(common.DEBUG, "LocalDbConnect", "Started")

	// Load all DB configurations
	lDbDetails := new(AllUsedDatabases)
	if lErr := lDbDetails.Init(log); lErr != nil {
		log.Log(common.ERROR, "LocalDbConnect", fmt.Sprintf("Failed to init DB details: %v", lErr))
		return nil, lErr
	}

	var lConnString string
	var lDBtype string
	var lDataBaseConnection DatabaseType

	// Match the requested DB type with loaded configuration
	if pDbName == lDbDetails.DbName.DB {
		lDataBaseConnection = lDbDetails.DbName
		lDBtype = lDbDetails.DbName.DBType
	}

	// Build the connection string based on DB driver type
	switch lDBtype {
	case "mssql":
		lConnString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
			lDataBaseConnection.Server, lDataBaseConnection.User, lDataBaseConnection.Password,
			lDataBaseConnection.Port, lDataBaseConnection.Database)

	case "mysql":
		lConnString = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			lDataBaseConnection.User, lDataBaseConnection.Password, lDataBaseConnection.Server,
			lDataBaseConnection.Port, lDataBaseConnection.Database)

	case "postgres":
		lConnString = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
			lDataBaseConnection.Server, lDataBaseConnection.Port, lDataBaseConnection.User,
			lDataBaseConnection.Password, lDataBaseConnection.Database)

	default:
		// Unsupported or missing DB type in config
		return nil, fmt.Errorf("unsupported DB type: %s", lDBtype)
	}

	// Read connection pooling limits from configuration
	var lDBConnectionPool DBConnectionPool
	if lErr := config.GetAndAssignTomlValue("dbconfig", "DBConnectionPool", &lDBConnectionPool); lErr != nil {
		log.Log(common.ERROR, "LocalDbConnect", fmt.Sprintf("Error reading DbConMaxIdleTime: %v", lErr))
	}

	// Attempt to open the database connection
	lDb, lErr := sql.Open(lDBtype, lConnString)
	if lErr != nil {
		log.Log(common.ERROR, "LocalDbConnect", fmt.Sprintf("Failed to open DB connection: %v", lErr))
		return nil, lErr
	}

	// Configure connection pooling parameters
	lDb.SetMaxOpenConns(lDBConnectionPool.DbConMaxOpenConns)
	lDb.SetMaxIdleConns(lDBConnectionPool.DbConMaxIdleConns)
	lDb.SetConnMaxIdleTime(time.Second * time.Duration(lDBConnectionPool.DbConMaxIdleTime))

	log.Log(common.DEBUG, "LocalDbConnect", "Finished successfully")
	return lDb, nil
}
