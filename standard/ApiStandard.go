package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"lumelpkg/common"
	"net/http"
	"strings"
)

// TestApi is an HTTP handler for testing the API.
func TestApi(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	(w).Header().Set("Access-Control-Allow-Credentials", "true")
	(w).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(w).Header().Set("Access-Control-Allow-Headers", "USER,Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")
	log.Println("TestApi(+)")

	// Create a response struct
	var lRespRec ResponseStruct

	// Handle POST requests
	if strings.EqualFold(r.Method, http.MethodPost) {
		// Set default success status
		lRespRec.Status = common.SuccessCode

		// Get USER header
		lUser := r.Header.Get("USER")
		if lUser != "" {
			// Read request body
			lBody, lErr := io.ReadAll(r.Body)
			if lErr != nil {
				log.Println("TestApi:001" + lErr.Error())
				lRespRec.Status = common.ErrorCode
				lRespRec.ErrMsg = "TestApi:001" + lErr.Error()
			} else {
				// Unmarshal response data
				lErr := json.Unmarshal(lBody, &lRespRec)
				if lErr != nil {
					log.Println("TestApi:001" + lErr.Error())
					lRespRec.Status = common.ErrorCode
					lRespRec.ErrMsg = "TestApi:001" + lErr.Error()
				}
			}
		} else {
			// Invalid User
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = "Invalid User"
		}
	} else {
		// Invalid Method
		lRespRec.Status = common.ErrorCode
		lRespRec.ErrMsg = "Invalid Method"
	}

	// Marshal response data
	lData, lErr := json.Marshal(lRespRec)
	if lErr != nil {
		fmt.Fprintf(w, "Error taking data"+lErr.Error())
	} else {
		// Write response data
		fmt.Fprintf(w, string(lData))
	}

	log.Println("TestApi(-)")
}
