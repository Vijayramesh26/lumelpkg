package appscommon

import (
	"log"
	"lumelpkg/common"
	"lumelpkg/config"
	"lumelpkg/utils"
	"net/http"
)

func ResetToml(lHttpWriter http.ResponseWriter, lHttpRequest *http.Request) {
	log.Println("ResetToml (+)")
	// Initialize logger with a unique request ID for tracing
	log := new(utils.Logger)
	log.SetReqID()
	log.Log(common.INFO, "ResetToml", "Started")

	// Set HTTP headers for CORS and request handling
	lHttpWriter.Header().Set("Access-Control-Allow-Origin", "*")
	lHttpWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	lHttpWriter.Header().Set("Access-Control-Allow-Methods", "POST")
	lHttpWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if lHttpRequest.Method == http.MethodGet {
		// Global toml Values Read
		config.Init(log)
	}
	lHttpWriter.WriteHeader(200)

	log.Log(common.INFO, "ResetToml", "Finished")
}
