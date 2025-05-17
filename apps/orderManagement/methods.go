package ordermanagement

import (
	"errors"
	"fmt"
	ordercommon "lumelpkg/apps/orderManagement/common"
	"lumelpkg/common"
	"lumelpkg/db"
	"lumelpkg/utils"
)

/*
   Purpose : This method is used to communicate and fetch client Basic details, CRM Details, Exchange Details.
   Parameter : pDebug - *common.commontruct,  pClientId, FetchDetails string
   Response : Response array and error if any.
   Author : VIJAY
   Created Date : 11-04-2025
*/
// CommunicateWithDB retrieves external data
func CommunicateWithDB(log *utils.Logger, pReqRec ordercommon.RequestStruct, pKeyToFetch string) (any, error) {
	log.Log(common.INFO, "CommunicateWithDB (+)")

	switch pKeyToFetch {
	case ordercommon.GetTotalRevenue:
		// Here we fetch total revenue
		lTotalRevenue, lErr := GetTotalRevenue(log, pReqRec)
		if lErr != nil {
			log.Log(common.ERROR, "CommunicateWithDB:002 -", lErr.Error()+" Error While fetching Client Basic Details")
			return lTotalRevenue, errors.New(" Error While fetching Client Basic Details" + lErr.Error())
		} else {
			return lTotalRevenue, nil
		}
	case ordercommon.GetCategoryRevenue:
		// Here we fetch revenue by category
		lCategoryRevenue, lErr := GetCategoryRevenue(log, pReqRec)
		if lErr != nil {
			log.Log(common.ERROR, "CommunicateWithDB:003 -", lErr.Error()+" Error While fetching Client CRM Details")
			return lCategoryRevenue, errors.New(" Error While fetching Client CRM Details" + lErr.Error())
		} else {
			return lCategoryRevenue, nil
		}
	case ordercommon.GetProductRevenue:
		// Here we fetch revenue by product
		lProductRevenue, lErr := GetProductRevenue(log, pReqRec)
		if lErr != nil {
			log.Log(common.ERROR, "CommunicateWithDB:004 -", lErr.Error()+" Error While fetching Client Exchange Details")
			return lProductRevenue, errors.New(" Error While fetching Client Exchange Details" + lErr.Error())
		} else {
			return lProductRevenue, nil
		}
	case ordercommon.GetRegionRevenue:
		// Here we fetch revenue by region
		lRegionRevenue, lErr := GetRegionRevenue(log, pReqRec)
		if lErr != nil {
			log.Log(common.ERROR, "CommunicateWithDB:005 -", lErr.Error()+" Error While fetching Client form stages Details")
			return lRegionRevenue, errors.New(" Error While fetching Client Exchange Details" + lErr.Error())
		} else {
			return lRegionRevenue, nil
		}
	}

	log.Log(common.INFO, "CommunicateWithDB (-)")
	return nil, nil
}

/*
   Purpose : This method is used to fetch client Request Id.
   Parameter : pDebug - *common.commontruct,   lClientID - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 11-04-2025
*/

func GetTotalRevenue(log *utils.Logger, pReqRec ordercommon.RequestStruct) (lReqRec ordercommon.RevenueStruct, lErr error) {
	log.Log(common.INFO, "GetTotalRevenue (+)")

	lCoreString := `SELECT 
			SUM(quantity_sold * unit_price * (1 - discount)) AS RevenueWithDiscount,
			SUM(quantity_sold * unit_price) AS RevenueWithoutDiscount
		FROM order_items oi
		JOIN orders o ON o.order_id = oi.order_id
		WHERE o.date_of_sale BETWEEN ? AND ?`
	lStmt, lErr := db.Global_DB_Instance.Prepare(lCoreString)
	if lErr != nil {
		log.Log(common.ERROR, "GTR-001", lErr.Error())
		return lReqRec, fmt.Errorf("GetTotalRevenue - (GTR-001) " + lErr.Error())
	}
	defer lStmt.Close()

	lRows, lErr := lStmt.Query(pReqRec.FromDate, pReqRec.ToDate)
	if lErr != nil {
		log.Log(common.ERROR, "GTR-002", lErr.Error())
		return lReqRec, fmt.Errorf("GetTotalRevenue - (GTR-002) " + lErr.Error())
	}
	defer lRows.Close()

	for lRows.Next() {
		lErr := lRows.Scan(&lReqRec.RevenueWithDiscount, &lReqRec.RevenueWithDiscount)
		if lErr != nil {
			log.Log(common.ERROR, "GTR-003", lErr.Error())
			return lReqRec, fmt.Errorf("GetTotalRevenue - (GTR-003) " + lErr.Error())
		} else {
			//  Your Logic Here
			log.Log(common.DEBUG, "GetTotalRevenue  ", "lReqId")
		}
	}

	log.Log(common.INFO, "GetTotalRevenue (-)")
	return lReqRec, nil
}

/*
   Purpose : This method is used to fetch client Request Id.
   Parameter : pDebug - *common.commontruct,   lClientID - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 11-04-2025
*/

func GetCategoryRevenue(log *utils.Logger, pReqRec ordercommon.RequestStruct) (lReqArr []ordercommon.RevenueResp, lErr error) {
	log.Log(common.INFO, "GetCategoryRevenue (+)")
	var lReqRec ordercommon.RevenueResp

	lCoreString := `SELECT 
						p.category AS CatagoryName,
						SUM(quantity_sold * unit_price * (1 - discount)) AS RevenueWithDiscount,
						SUM(quantity_sold * unit_price) AS RevenueWithoutDiscount
					FROM order_items oi
					JOIN products p ON oi.product_id = p.product_id
					JOIN orders o ON o.order_id = oi.order_id
					WHERE o.date_of_sale BETWEEN ? AND ?
					GROUP BY p.category`
	lStmt, lErr := db.Global_DB_Instance.Prepare(lCoreString)
	if lErr != nil {
		log.Log(common.ERROR, "GCR-001", lErr.Error())
		return lReqArr, fmt.Errorf("GetCategoryRevenue - (GCR-001) " + lErr.Error())
	}
	defer lStmt.Close()

	lRows, lErr := lStmt.Query(pReqRec.FromDate, pReqRec.ToDate)
	if lErr != nil {
		log.Log(common.ERROR, "GCR-002", lErr.Error())
		return lReqArr, fmt.Errorf("GetCategoryRevenue - (GCR-002) " + lErr.Error())
	}
	defer lRows.Close()

	for lRows.Next() {
		lErr := lRows.Scan(&lReqRec.CatagoryName, &lReqRec.RevenueWithDiscount, &lReqRec.RevenueWithDiscount)
		if lErr != nil {
			log.Log(common.ERROR, "GCR-003", lErr.Error())
			return lReqArr, fmt.Errorf("GetCategoryRevenue - (GCR-003) " + lErr.Error())
		} else {
			//  Your Logic Here
			lReqArr = append(lReqArr, lReqRec)
			log.Log(common.DEBUG, "GetCategoryRevenue  ", "lReqId")
		}
	}

	log.Log(common.INFO, "GetCategoryRevenue (-)")
	return lReqArr, nil
}

/*
   Purpose : This method is used to fetch client Request Id.
   Parameter : pDebug - *common.commontruct,   lClientID - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 11-04-2025
*/

func GetProductRevenue(log *utils.Logger, pReqRec ordercommon.RequestStruct) (lReqArr []ordercommon.RevenueResp, lErr error) {
	log.Log(common.INFO, "GetProductRevenue (+)")
	var lReqRec ordercommon.RevenueResp

	lCoreString := `SELECT 
						p.name AS ProductName,
						SUM(quantity_sold * unit_price * (1 - discount)) AS RevenueWithDiscount,
						SUM(quantity_sold * unit_price) AS RevenueWithoutDiscount
					FROM order_items oi
					JOIN products p ON oi.product_id = p.product_id
					JOIN orders o ON o.order_id = oi.order_id
					WHERE o.date_of_sale BETWEEN ? AND ?
					GROUP BY p.product_id, p.name`
	lStmt, lErr := db.Global_DB_Instance.Prepare(lCoreString)
	if lErr != nil {
		log.Log(common.ERROR, "GPR-001", lErr.Error())
		return lReqArr, fmt.Errorf("GetProductRevenue - (GPR-001) " + lErr.Error())
	}
	defer lStmt.Close()

	lRows, lErr := lStmt.Query(pReqRec.FromDate, pReqRec.ToDate)
	if lErr != nil {
		log.Log(common.ERROR, "GPR-002", lErr.Error())
		return lReqArr, fmt.Errorf("GetProductRevenue - (GPR-002) " + lErr.Error())
	}
	defer lRows.Close()

	for lRows.Next() {
		lErr := lRows.Scan(&lReqRec.ProductName, &lReqRec.RevenueWithDiscount, &lReqRec.RevenueWithDiscount)
		if lErr != nil {
			log.Log(common.ERROR, "GPR-003", lErr.Error())
			return lReqArr, fmt.Errorf("GetProductRevenue - (GPR-003) " + lErr.Error())
		} else {
			//  Your Logic Here
			lReqArr = append(lReqArr, lReqRec)
			log.Log(common.DEBUG, "GetProductRevenue  ", "lReqId")
		}
	}

	log.Log(common.INFO, "GetProductRevenue (-)")
	return lReqArr, nil
}

/*
   Purpose : This method is used to fetch client Request Id.
   Parameter : pDebug - *common.commontruct,   lClientID - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 11-04-2025
*/

func GetRegionRevenue(log *utils.Logger, pReqRec ordercommon.RequestStruct) (lReqArr []ordercommon.RevenueResp, lErr error) {
	log.Log(common.INFO, "GetProductRevenue (+)")
	var lReqRec ordercommon.RevenueResp

	lCoreString := `SELECT 
						o.region AS RegionName,
						SUM(quantity_sold * unit_price * (1 - discount)) AS RevenueWithDiscount,
						SUM(quantity_sold * unit_price) AS RevenueWithoutDiscount
					FROM order_items oi
					JOIN products p ON oi.product_id = p.product_id
					JOIN orders o ON o.order_id = oi.order_id
					WHERE o.date_of_sale BETWEEN ? AND ?
					GROUP BY o.region`
	lStmt, lErr := db.Global_DB_Instance.Prepare(lCoreString)
	if lErr != nil {
		log.Log(common.ERROR, "GRR-001", lErr.Error())
		return lReqArr, fmt.Errorf("GetProductRevenue - (GRR-001) " + lErr.Error())
	}
	defer lStmt.Close()

	lRows, lErr := lStmt.Query(pReqRec.FromDate, pReqRec.ToDate)
	if lErr != nil {
		log.Log(common.ERROR, "GRR-002", lErr.Error())
		return lReqArr, fmt.Errorf("GetProductRevenue - (GRR-002) " + lErr.Error())
	}
	defer lRows.Close()

	for lRows.Next() {
		lErr := lRows.Scan(&lReqRec.RegionName, &lReqRec.RevenueWithDiscount, &lReqRec.RevenueWithDiscount)
		if lErr != nil {
			log.Log(common.ERROR, "GRR-003", lErr.Error())
			return lReqArr, fmt.Errorf("GetProductRevenue - (GRR-003) " + lErr.Error())
		} else {
			//  Your Logic Here
			lReqArr = append(lReqArr, lReqRec)
			log.Log(common.DEBUG, "GetProductRevenue  ", "lReqId")
		}
	}

	log.Log(common.INFO, "GetProductRevenue (-)")
	return lReqArr, nil
}
