package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

// Logger struct to hold the request ID (ReqID)
type Logger struct {
	ReqID string
}

// Initialize the logger (This will configure the log file and output)
func InitLogger() {
	// Open log file for appending
	logFile, err := os.OpenFile("./log/logfile"+time.Now().Format("02012006.15.04.05.000000000")+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Set the log output to both file and console
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags) // Log with timestamp
}

// Generate a new unique Request ID (ReqID)
func GenerateReqID() string {
	return uuid.New().String() // Generate and return a new UUID as Request ID
}

// Set the Request ID in the Logger instance
func (l *Logger) SetReqID() {
	// Generate a new ReqID for each request
	l.ReqID = GenerateReqID()
}

// Log method to print logs with the Request ID and message, including log level
func (l *Logger) Log(level, step string, message ...any) {
	// Format log with timestamp, log level, and ReqID
	log.Printf("%s [%s] [ReqID: %s] %s [Step %s] %s", time.Now().Format("2006/01/02 15:04:05"), level, l.ReqID, level, step, fmt.Sprintf("%v", message))
}
