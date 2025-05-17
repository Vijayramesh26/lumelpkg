package api

import (
	"lumelpkg/apps/appscommon"
	ordermanagement "lumelpkg/apps/orderManagement"
	ordercommon "lumelpkg/apps/orderManagement/common"
	"lumelpkg/common"
	"lumelpkg/utils"
	"net/http"
	"strings"
)

func FetchProductRevenue(lHttpWriter http.ResponseWriter, lHttpRequest *http.Request) {
	log := new(utils.Logger)
	log.SetReqID()
	log.Log(common.INFO, "FetchProductRevenue (+)")

	(lHttpWriter).Header().Set("Access-Control-Allow-Origin", "*")
	(lHttpWriter).Header().Set("Access-Control-Allow-Credentials", "true")
	(lHttpWriter).Header().Set("Access-Control-Allow-Methods", "POST,OPTIONS")
	(lHttpWriter).Header().Set("Access-Control-Allow-Headers", "Accept,Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization")

	var lRespRec common.CommonResp
	var lReqRec ordercommon.RequestStruct

	if strings.EqualFold(http.MethodPost, lHttpRequest.Method) {

		lErr := appscommon.CollectRequest(log, lHttpRequest, &lReqRec)
		if lErr != nil {
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = "Error In request Data"
			log.Log(common.ERROR, "Error In request Data", lErr.Error())
			lHttpWriter.WriteHeader(http.StatusBadRequest)
			goto marshal
		}
		lErr = appscommon.ValidateRequest(log, &lReqRec, lHttpRequest)
		if lErr != nil {
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = lErr.Error()
			log.Log(common.ERROR, "Error In Validate Data", lErr.Error())
			lHttpWriter.WriteHeader(http.StatusBadRequest)
			goto marshal
		}
		lErr = appscommon.CompareDates(lReqRec.FromDate, lReqRec.ToDate, common.DateLayout)
		if lErr != nil {
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = lErr.Error()
			log.Log(common.ERROR, "Error In request Data", lErr.Error())
			lHttpWriter.WriteHeader(http.StatusBadRequest)
			goto marshal
		}
		lRespRec.DetailsArr, lErr = ordermanagement.CommunicateWithDB(log, lReqRec, ordercommon.GetProductRevenue)
		if lErr != nil {
			lRespRec.Status = common.ErrorCode
			lRespRec.ErrMsg = lErr.Error()
			log.Log(common.ERROR, "Error In Communicate with DB", lErr.Error())
			lHttpWriter.WriteHeader(http.StatusInternalServerError)
			goto marshal
		}
		lRespRec.Status = common.SuccessCode
	}
marshal:
	appscommon.CompleteAndMarshall(log, lRespRec, lHttpWriter)
	log.Log(common.INFO, "FetchProductRevenue (-)")

}
