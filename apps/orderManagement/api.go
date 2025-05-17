package ordermanagement

// import (
// 	"fmt"
// 	"lumelpkg/common"
// 	"lumelpkg/utils"
// 	"net/http"
// 	"strings"
// )

// /*
//    Purpose : This API is fetch Client Basic Details from EKYC Table.
//    Parameter : lHttpWriter http.ResponseWriter, lHttpRequest *http.Request
//    Response : Success or Error response with status code and error message.
//    Author : VIJAY
//    Created Date : 11-04-2025
// */
// // FetchExchangeDetails handles ClientDetailsResp
// func FetchExchangeDetails(lHttpWriter http.ResponseWriter, lHttpRequest *http.Request) {
// 	// Initialize logger with a unique request ID for tracing
// 	log := new(utils.Logger)
// 	log.SetReqID()
// 	log.Log(common.INFO, "FetchExchangeDetails (+)")
// 	lHttpWriter.Header().Set("Access-Control-Allow-Origin", "*")
// 	lHttpWriter.Header().Set("Access-Control-Allow-Credentials", "true")
// 	lHttpWriter.Header().Set("Access-Control-Allow-Methods", fmt.Sprintf("%s, %s", http.MethodGet, http.MethodOptions))
// 	lHttpWriter.Header().Set("Access-Control-Allow-Headers", "CLIENTID,USER, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 	var lResponseRec common.CommonResp
// 	if strings.EqualFold(lHttpRequest.Method, http.MethodGet) {
// 		// Here we validate is that client is valid and user is in valid session or not
// 		lClientID, lErr := mode ValidateClient(log, lHttpRequest, "CLIENTID", "USER")
// 		if lErr != nil {
// 			lResponseRec.Status = common.ErrorCode
// 			lResponseRec.ErrMsg = "FetchExchangeDetails : 001 - Validation failed: " + lErr.Error()
// 			goto Complete
// 		}
// 		// Here we communicate with DB and Produce Client details
// 		lResponseRec.DetailsArr, lErr = CommunicateWithDB(log, lClientID, GetExchangeDeatils)
// 		if lErr != nil {
// 			lResponseRec.Status = common.ErrorCode
// 			lResponseRec.ErrMsg = "FetchExchangeDetails : 002 - Communication failure: " + lErr.Error()
// 			goto Complete
// 		}
// 		// Finally we set status as success
// 		lResponseRec.Status = common.SuccessCode
// 	}
// Complete:
// 	// Here we marshall data and sent it to requester
// 	CompleteAndMarshall(log, lResponseRec, lHttpWriter)
// 	log.Log(common.INFO, "FetchExchangeDetails (-)")
// }

// /*
//    Purpose : This method is used to fetch client exchange details.
//    Parameter : pDebug - *helpers.HelperStruct,   lClientID - string
//    Response :

//    On Success:
//    ===========
//    In case of a successful execution of this method, you will get response details.

//    On Error:
//    ===========
//    In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

//    Author : VIJAY
//    Date : 11-04-2025
// */

// func getExchangeDetails(pDebug *helpers.HelperStruct, lReqId string) (lClientDetails ExchangeDetails, lErr error) {
// 	pDebug.Log(helpers.Statement, "getExchangeDetails (+)")

// 	lCoreString := `SELECT
// 						NVL(MAX(CASE WHEN STAGE = 'CVLKRA' AND STATUS = 'AC' THEN 'Y' ELSE 'N' END),'NA') CVLKRA,
// 						NVL(MAX(CASE WHEN STAGE = 'CDSL' AND STATUS = 'DA' THEN 'Y' ELSE 'N' END),'NA') CDSL,
// 						NVL(MAX(CASE WHEN STAGE = 'NSE' AND STATUS = 'NA' THEN 'Y' ELSE 'N' END),'NA') NSE,
// 						NVL(MAX(CASE WHEN STAGE = 'BSE' AND STATUS = 'BA' THEN 'Y' ELSE 'N' END),'NA') BSE,
// 						NVL(MAX(CASE WHEN STAGE = 'MCX' AND STATUS = 'MA' THEN 'Y' ELSE 'N' END),'NA') MCX,
// 						NVL(MAX(CASE WHEN STAGE = 'BACKOFFICE' AND STATUS = 'AB' THEN 'Y' ELSE 'N' END),'NA') BACKOFFICE,
// 						NVL(MAX(CASE WHEN STAGE = 'BSE MF' AND STATUS = 'BMA' THEN 'Y' ELSE 'N' END),'NA') BSE_MF
// 					FROM
// 						NEWEKYC_INTEGRATION_HISTORY NIH
// 					WHERE
// 						1 = 1
// 					AND NIH.REQUESTUID = ?`
// 	lStmt, lErr := ftdb.G_MariaEKYCPRD_Instance.Prepare(lCoreString)
// 	if lErr != nil {
// 		pDebug.Log(helpers.Elog, "GED-001", lErr.Error())
// 		return lClientDetails, fmt.Errorf("getExchangeDetails - (GED-001) " + lErr.Error())
// 	}
// 	defer lStmt.Close()

// 	lRows, lErr := lStmt.Query(lReqId)
// 	if lErr != nil {
// 		pDebug.Log(helpers.Elog, "GED-002", lErr.Error())
// 		return lClientDetails, fmt.Errorf("getExchangeDetails - (GED-002) " + lErr.Error())
// 	}
// 	defer lRows.Close()

// 	for lRows.Next() {
// 		lRows.Columns()
// 		lErr := lRows.Scan(&lClientDetails.Cvlkra, &lClientDetails.Cdsl, &lClientDetails.Nse, &lClientDetails.Bse, &lClientDetails.Mcx, &lClientDetails.Backoffice, &lClientDetails.Bsemf)
// 		if lErr != nil {
// 			pDebug.Log(helpers.Elog, "GED-003", lErr.Error())
// 			return lClientDetails, fmt.Errorf("getExchangeDetails - (GED-003) " + lErr.Error())
// 		} else {
// 			//  Your Logic Here
// 			pDebug.Log(helpers.Details, "getExchangeDetails  ", lClientDetails)
// 		}
// 	}

// 	pDebug.Log(helpers.Statement, "getExchangeDetails (-)")
// 	return lClientDetails, nil
// }
