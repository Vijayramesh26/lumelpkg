package scheduler

import (
	"fmt"
	ordercommon "lumelpkg/apps/orderManagement/common"
	"lumelpkg/common"
	"lumelpkg/db"
	"lumelpkg/utils"
	"time"
)

func SchedularInit() {
	log := new(utils.Logger)
	log.SetReqID()
	log.Log(common.INFO, "LoadCSVFile ", "Started")

	// File Path and delimeter
	pFilePath := "./uploadedfiles/orderManagemrnt/OrderDetails.csv"
	pDelimeter := ','

	// Run immediately
	lErr := LoadCSVFile(log, pFilePath, pDelimeter)
	if lErr != nil {
		log.Log(common.ERROR, "Error during initial data refresh:", lErr.Error())
	}

	// Set ticker to run every 24 hours
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		log.Log(common.INFO, "Scheduled data refresh")
		lErr := LoadCSVFile(log, pFilePath, pDelimeter)
		if lErr != nil {
			log.Log(common.ERROR, "Error during scheduled data refresh:", lErr.Error())
		}
	}
	// log.Log(common.INFO, "LoadCSVFile ", "Ended")
}

func LoadCSVFile(log *utils.Logger, pFilePath string, pDelimeter rune) error {
	log.Log(common.INFO, "LoadCSVFile ", "Started")

	// Load CSV data to structure
	lCsvData, lErr := utils.LoadCSV[ordercommon.CsvData]("./uploadedfiles/orderManagemrnt/OrderDetails.csv", ',')
	if lErr != nil {
		log.Log(common.ERROR, "LoadCSVFile ", lErr.Error())
		return lErr
	}

	// Initialize instance
	var lCustomerRec ordercommon.Customer
	var lProductRec ordercommon.Product
	var lOrderRec ordercommon.Order
	var lOrderItemRec ordercommon.OrderItem

	for _, record := range lCsvData {
		lCustomerRec = ordercommon.Customer{
			CustomerID:      record.CustomerID,
			CustomerName:    record.CustomerName,
			CustomerEmail:   record.CustomerEmail,
			CustomerAddress: record.CustomerAddress,
		}

		lProductRec = ordercommon.Product{
			ProductID:   record.ProductID,
			ProductName: record.ProductName,
			Category:    record.Category,
		}

		lOrderRec = ordercommon.Order{
			OrderID:       record.OrderID,
			CustomerID:    record.CustomerID,
			Region:        record.Region,
			DateOfSale:    record.DateOfSale,
			ShippingCost:  record.ShippingCost,
			PaymentMethod: record.PaymentMethod,
		}

		lOrderItemRec = ordercommon.OrderItem{
			OrderID:   record.OrderID,
			ProductID: record.ProductID,
			Quantity:  record.Quantity,
			UnitPrice: record.UnitPrice,
			Discount:  record.Discount,
		}

	}
	if lErr = InsertCustomer(log, lCustomerRec); lErr != nil {
		log.Log(common.ERROR, "LoadCSVFile ", "InsertCustomer ", lErr.Error())
		return lErr
	}
	if lErr = InsertOrder(log, lOrderRec); lErr != nil {
		log.Log(common.ERROR, "LoadCSVFile ", "InsertOrder ", lErr.Error())
		return lErr
	}
	if lErr = InsertOrderItem(log, lOrderItemRec); lErr != nil {
		log.Log(common.ERROR, "LoadCSVFile ", "InsertOrderItem ", lErr.Error())
		return lErr
	}
	if lErr = InsertProducts(log, lProductRec); lErr != nil {
		log.Log(common.ERROR, "LoadCSVFile ", "InsertProducts ", lErr.Error())
		return lErr
	}

	log.Log(common.INFO, "LoadCSVFile ", "Started")
	return nil
}

/*
   Purpose : This method is used to describe purpose.
   Parameter : log - *helpers.HelperStruct, log - *helpers.HelperStruct, lRequestId - string, pParameterName - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 17-05-2025
*/

func InsertCustomer(log *utils.Logger, pCustomerData ordercommon.Customer) error {
	log.Log(common.INFO, "InsertCustomer (+)")

	lSqlString := `INSERT INTO customers (CustomerID, CustomerName, CustomerEmail, CustomerAddress)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE CustomerName=VALUES(CustomerName), CustomerEmail=VALUES(CustomerEmail), CustomerAddress=VALUES(CustomerAddress)`

	lExecResult, lErr := db.Global_DB_Instance.Exec(lSqlString, pCustomerData.CustomerID, pCustomerData.CustomerName, pCustomerData.CustomerEmail, pCustomerData.CustomerAddress)
	if lErr != nil {
		log.Log(common.ERROR, "IC-001 ", lErr.Error())
		return fmt.Errorf("InsertCustomer - (IC-001) " + lErr.Error())
	}

	lRowsAffected, lErr := lExecResult.RowsAffected()
	if lErr != nil {
		log.Log(common.ERROR, "IC-002 ", lErr.Error())
	}

	log.Log(common.DEBUG, "InsertCustomer Rows affected: ", lRowsAffected)
	log.Log(common.ERROR, "Record Inserted successfully")

	log.Log(common.INFO, "InsertCustomer (-)")
	return nil
}

/*
   Purpose : This method is used to describe purpose.
   Parameter : log - *helpers.HelperStruct, log - *helpers.HelperStruct, lRequestId - string, pParameterName - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 17-05-2025
*/

func InsertProducts(log *utils.Logger, pProductData ordercommon.Product) error {
	log.Log(common.INFO, "InsertProducts (+)")

	lSqlString := `INSERT INTO products (ProductID, ProductName, Category)
		VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE ProductName=VALUES(ProductName), Category=VALUES(Category)`

	lExecResult, lErr := db.Global_DB_Instance.Exec(lSqlString, pProductData.ProductID, pProductData.ProductName, pProductData.Category)
	if lErr != nil {
		log.Log(common.ERROR, "IC-001 ", lErr.Error())
		return fmt.Errorf("InsertProducts - (IC-001) " + lErr.Error())
	}

	lRowsAffected, lErr := lExecResult.RowsAffected()
	if lErr != nil {
		log.Log(common.ERROR, "IC-002 ", lErr.Error())
	}

	log.Log(common.DEBUG, "InsertProducts Rows affected: ", lRowsAffected)
	log.Log(common.ERROR, "Record Inserted successfully")

	log.Log(common.INFO, "InsertProducts (-)")
	return nil
}

/*
   Purpose : This method is used to describe purpose.
   Parameter : log - *helpers.HelperStruct, log - *helpers.HelperStruct, lRequestId - string, pParameterName - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 17-05-2025
*/

func InsertOrder(log *utils.Logger, pOrderData ordercommon.Order) error {
	log.Log(common.INFO, "InsertOrder (+)")

	lSqlString := `INSERT INTO orders (OrderID, CustomerID, Region, DateOfSale, ShippingCost, PaymentMethod)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE Region=VALUES(Region), DateOfSale=VALUES(DateOfSale), ShippingCost=VALUES(ShippingCost), PaymentMethod=VALUES(PaymentMethod)`

	lExecResult, lErr := db.Global_DB_Instance.Exec(lSqlString, pOrderData.OrderID, pOrderData.CustomerID, pOrderData.Region, pOrderData.DateOfSale, pOrderData.ShippingCost, pOrderData.PaymentMethod)
	if lErr != nil {
		log.Log(common.ERROR, "IC-001 ", lErr.Error())
		return fmt.Errorf("InsertOrder - (IC-001) " + lErr.Error())
	}

	lRowsAffected, lErr := lExecResult.RowsAffected()
	if lErr != nil {
		log.Log(common.ERROR, "IC-002 ", lErr.Error())
	}

	log.Log(common.DEBUG, "InsertOrder Rows affected: ", lRowsAffected)
	log.Log(common.ERROR, "Record Inserted successfully")

	log.Log(common.INFO, "InsertOrder (-)")
	return nil
}

/*
   Purpose : This method is used to describe purpose.
   Parameter : log - *helpers.HelperStruct, log - *helpers.HelperStruct, lRequestId - string, pParameterName - string
   Response :

   On Success:
   ===========
   In case of a successful execution of this method, you will get response details.

   On Error:
   ===========
   In case of any exception during the execution of this method, you will get the error details. The calling program should handle the error.

   Author : VIJAY
   Date : 17-05-2025
*/

func InsertOrderItem(log *utils.Logger, pOrderItems ordercommon.OrderItem) error {
	log.Log(common.INFO, "InsertOrderItem (+)")

	lSqlString := `INSERT INTO order_items (OrderID, ProductID, Quantity, UnitPrice, Discount)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE Quantity=VALUES(Quantity), UnitPrice=VALUES(UnitPrice), Discount=VALUES(Discount)`

	lExecResult, lErr := db.Global_DB_Instance.Exec(lSqlString, pOrderItems.OrderID, pOrderItems.ProductID, pOrderItems.Quantity, pOrderItems.UnitPrice, pOrderItems.Discount)
	if lErr != nil {
		log.Log(common.ERROR, "IC-001 ", lErr.Error())
		return fmt.Errorf("InsertOrderItem - (IC-001) " + lErr.Error())
	}

	lRowsAffected, lErr := lExecResult.RowsAffected()
	if lErr != nil {
		log.Log(common.ERROR, "IC-002 ", lErr.Error())
	}

	log.Log(common.DEBUG, "InsertOrderItem Rows affected: ", lRowsAffected)
	log.Log(common.ERROR, "Record Inserted successfully")

	log.Log(common.INFO, "InsertOrderItem (-)")
	return nil
}
