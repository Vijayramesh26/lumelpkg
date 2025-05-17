package main

import (
	"fmt"
	"lumelpkg/apps/appscommon"
	scheduler "lumelpkg/apps/orderManagement/Scheduler"
	"lumelpkg/apps/orderManagement/api"
	"lumelpkg/config"
	"lumelpkg/db"
	"lumelpkg/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the logger for app-wide logging
	utils.InitLogger()
	logger := &utils.Logger{}
	logger.SetReqID()

	// Load all global DB instance
	db.GlobalDBInit(logger)

	// Load all toml data in a global variable
	config.Init(logger)

	// Load CSV File Data
	scheduler.SchedularInit()

	// Set up the router
	router := mux.NewRouter()

	// Define the /ready route (GET method)
	router.HandleFunc("/ready", appscommon.Ready).Methods(http.MethodGet)
	router.HandleFunc("/orders/totalrevenue", api.FetchTotalRevenue).Methods(http.MethodPost)
	router.HandleFunc("/orders/prodrevenue", api.FetchProductRevenue).Methods(http.MethodPost)
	router.HandleFunc("/orders/categrevenue", api.FetchCategoryRevenue).Methods(http.MethodPost)
	router.HandleFunc("/orders/regionrevenue", api.FetchRegionRevenue).Methods(http.MethodPost)

	// Start the server
	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":26301", router)
}
