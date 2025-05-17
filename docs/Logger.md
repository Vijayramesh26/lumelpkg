# 🔧 Logger Utility for Go Applications

This utility provides a structured way to generate logs with timestamps, log levels, steps, and unique request IDs using Go's built-in `log` package and `uuid`.

## 📦 Package: `utils`

### ✅ Features

* Logs include timestamp, level, request ID, step, and custom messages.
* Logs are written to a file and can also appear in the console.
* Unique request ID (`ReqID`) per request or operation.
* Simple integration across your project.

---

## 📁 Directory Structure Example

```
project/
├── main.go
├── log/
│   └── logfileDDMMYYYY.HH.MM.SS.nanoseconds.txt
├── utils/
│   └── logger.go
```

---

## 🧩 Setup

### 1. Create the `log` Directory (if not exists)

```bash
mkdir log
```

### 2. Import and Initialize Logger

In your `main.go` or server bootstrap file:

```go
package main

import (
	"project/utils"
)

func main() {
	utils.InitLogger()
	// your application logic
}
```

---

## 🛠️ Usage in Any File

```go
package example

import (
	"project/utils"
)

func SomeFunction() {
	// Create a logger instance
	log := new(utils.Logger)
	
	// Assign a unique request ID
	log.SetReqID()

	// Log messages at various steps
	log.Log("INFO", "1", "Function started")
	log.Log("DEBUG", "2", "Some internal state:", 42)
	log.Log("ERROR", "3", "Something went wrong")
}
```

---

## 🔤 Log Levels Convention

You can use the following levels for consistency:

```go
const (
	INFO  = "INFO"
	DEBUG = "DEBUG"
	ERROR = "ERROR"
)
```

---

## 🧪 Sample Output (inside log file)

```
2025/05/16 10:21:03 [INFO] [ReqID: 3f6c8911-7c44-4a59-927d-87a21672c503] INFO [Step 1] Function started
2025/05/16 10:21:03 [DEBUG] [ReqID: 3f6c8911-7c44-4a59-927d-87a21672c503] DEBUG [Step 2] Some internal state: 42
2025/05/16 10:21:03 [ERROR] [ReqID: 3f6c8911-7c44-4a59-927d-87a21672c503] ERROR [Step 3] Something went wrong
```

---

## 🔁 Best Practice

* Use `SetReqID()` per request or unit of work.
* Use meaningful `step` values (e.g., `"1"`, `"start"`, `"DB_CALL"`) for tracing.
* Log key actions to trace full execution paths with `ReqID`.

---

## 📌 Notes

* Logs are stored in the `./log/` directory with timestamped filenames.
* You can customize log file naming or format in `InitLogger()`.

---

