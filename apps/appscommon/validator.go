package appscommon

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"lumelpkg/common"
	"lumelpkg/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

/*
Purpose : This method is used to Collect request from user and unmarshall and assign to given struture.
Parameter : log - *utils.Logger,pHttpRequest *http.Request, pRequestData *any
Response :

On Success:
===========
In case of a successful execution of this method, you will get response details.

On Error:
===========
In case of any exception during the execution of this method you will get the error details. The calling program should handle the error.

Author : VIJAY
Date : 17-05-2025
*/
func CollectRequest(log *utils.Logger, pHttpRequest *http.Request, pRequestData any) error {

	log.Log(common.INFO, "CollectRequest (+)")

	lBody, lErr := io.ReadAll(pHttpRequest.Body)
	if lErr != nil {
		log.Log(common.ERROR, "CCR:001", lErr)
		return (lErr)
	}

	log.Log(common.DEBUG, "Raw Body: ", string(lBody))

	if lErr = json.Unmarshal(lBody, &pRequestData); lErr != nil {
		log.Log(common.ERROR, "CCR:002", lErr)
		return (lErr)
	}

	log.Log(common.INFO, "CollectRequest (-)")
	return nil
}

/*
   Purpose : This method is used to validate request data from the client and validate token is valid or not.
   Parameter : pDebug - *helpers.HelperStruct, pRequestData RequestStruct
   Response :

   On Success:
   ===========
	In case of a successful execution of this method, you will get nil as a response.

   On Error:
   ===========
   	In case of any exception during the execution of this method you will send the error details.

   Author : VIJAY
   Date : 01-04-2025
*/
// ValidateRequest validates input data
func ValidateRequest(log *utils.Logger, pRequestData any, pHttpRequest *http.Request) error {
	log.Log(common.INFO, "GlobalDBInit (+)")
	lValidate := validator.New()
	lErr := lValidate.Struct(pRequestData)
	if lErr != nil {
		var validationErrors string
		for _, err := range lErr.(validator.ValidationErrors) {
			validationErrors += fmt.Sprintf("The field '%s' failed validation: it must satisfy the '%s' rule%s.",
				err.Field(),
				err.ActualTag(),
				func() string {
					if err.Param() != "" {
						return fmt.Sprintf(" with parameter '%s'", err.Param())
					}
					return ""
				}(),
			)
		}
		log.Log(common.ERROR, "Validation failed: ", validationErrors)
		log.Log(common.ERROR, "ValidateRequest : 001 (PCVR-001) request validation failed", lErr.Error())
		return errors.New(validationErrors)
	}
	log.Log(common.INFO, "ValidateRequest (-)")
	return nil
}

/*
   Purpose : This method is used to marshall data and sent it to requester.
   Parameter : pDebug - *helpers.HelperStruct, pResponseRec ClientDetailsResp, pHttpWriter http.ResponseWriter
   Response : Writes the response to the HTTP writer.
   Author : VIJAY
   Created Date : 11-04-2025
*/
// CompleteAndMarshall sends the final response
func CompleteAndMarshall(log *utils.Logger, pResponseRec common.CommonResp, pHttpWriter http.ResponseWriter) {
	log.Log(common.INFO, "CompleteAndMarshall (+)")
	lData, lErr := json.Marshal(pResponseRec)
	if lErr != nil {
		http.Error(pHttpWriter, "Error marshaling response: "+lErr.Error(), http.StatusInternalServerError)
		return
	}
	pHttpWriter.WriteHeader(http.StatusOK)
	pHttpWriter.Write(lData)
	log.Log(common.INFO, "CompleteAndMarshall (-)")
}

// CompareDates checks if fromDate is after toDate.
// Returns an error if the date format is invalid or if fromDate > toDate.
func CompareDates(fromDateStr, toDateStr, layout string) error {
	fromDate, err := time.Parse(layout, fromDateStr)
	if err != nil {
		return errors.New("invalid FromDate format")
	}

	toDate, err := time.Parse(layout, toDateStr)
	if err != nil {
		return errors.New("invalid ToDate format")
	}

	if fromDate.After(toDate) {
		return errors.New("ToDate should be greater than FromDate")
	}

	return nil
}
