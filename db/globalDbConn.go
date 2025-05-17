package db

import (
	"fmt"
	"lumelpkg/utils"
)

func GlobalDBInit(log *utils.Logger) {
	log.Log("INFO", "GlobalDBInit (+)")
	var lErr error
	if Global_DB_Instance, lErr = LocalDbConnect(SQLDB); lErr != nil {
		log.Log("ERROR", "GlobalDBInit", fmt.Sprintf("loading lDbName: %v", lErr))
	}
	log.Log("INFO", "GlobalDBInit(-)", "Successfully loaded lDbName config")
}
