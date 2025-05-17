package appscommon

import (
	"lumelpkg/common"
	"lumelpkg/utils"
	"net/http"
	"strings"
)

func Ready(lHttpWriter http.ResponseWriter, lHttpRequest *http.Request) {
	// Initialize logger with a unique request ID for tracing
	log := new(utils.Logger)
	log.SetReqID()
	log.Log(common.INFO, "Ready", "Started")
	// Set HTTP headers for CORS and request handling
	lHttpWriter.Header().Set("Access-Control-Allow-Origin", "*")
	lHttpWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	lHttpWriter.Header().Set("Access-Control-Allow-Methods", http.MethodGet)
	lHttpWriter.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if strings.EqualFold(lHttpRequest.Method, http.MethodGet) {
		log.Log(common.DEBUG, "STATUS => ", http.StatusOK)
		lHttpWriter.WriteHeader(http.StatusOK)
	}
	log.Log(common.INFO, "Ready", "Finished")

}
