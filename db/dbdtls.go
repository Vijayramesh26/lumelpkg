package db

import (
	"fmt"
	"lumelpkg/common"
	"lumelpkg/config"
	"lumelpkg/utils"
)

const (
	SQLDB = "localDB"
)

// Init loads the database configurations from the toml config file.
// It logs errors and returns an error if loading fails.
func (pDb *AllUsedDatabases) Init(log *utils.Logger) error {
	if lErr := config.GetAndAssignTomlValue("dbconfig", "DbName", &pDb.DbName); lErr != nil {
		log.Log(common.ERROR, "Init", fmt.Sprintf("loading lDbName failed: %v", lErr))
		return fmt.Errorf("loading lDbName: %w", lErr)
	}
	log.Log(common.INFO, "Init", "lDbName loaded successfully")
	return nil
}
