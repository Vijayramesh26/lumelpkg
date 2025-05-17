# 🛠️ TOML Configuration Loader (Go)

This package provides utilities for loading `.toml` files and safely accessing configuration values (including nested values) using a dot-separated key path. It supports unmarshalling directly into native Go types or structs.

## ✅ Features

- Load all `.toml` files from a folder into a global `ConfigMap`
- Access single-level and nested config keys using dot-separated paths
- Thread-safe (uses sync.RWMutex)
- Assign values directly to Go variables using JSON-based conversion

---

## 📦 Loading Configs

```go
import "yourmodule/config"

func main() {
    config.LoadTOMLFile() // Loads all .toml files under ./toml/
}
```

> This must be called once during application initialization.

---

## 🔑 Accessing Config Values

### Method: `GetAndAssignTomlValue`

```go
func GetAndAssignTomlValue(filename, key string, out any) error
```

- `filename`: name of the TOML file without the `.toml` extension
- `key`: dot-separated key path (e.g., `database.host`, `credentials.username`)
- `out`: pointer to a Go variable (e.g., `&myString`, `&myStruct`)

---

## 🧪 Example Usage

### app.toml

```toml
appname = "MyApp"

[database]
  host = "localhost"
  port = 5432

  [database.credentials]
    username = "admin"
    password = "secret"
```

### Go Code

```go
var appName string
_ = config.GetAndAssignTomlValue("app", "appname", &appName)

var host string
_ = config.GetAndAssignTomlValue("app", "database.host", &host)

var creds struct {
    Username string
    Password string
}
_ = config.GetAndAssignTomlValue("app", "database.credentials", &creds)
```

---

## ⚠️ Notes

- If the file or key is not found, an error is returned.
- The `out` parameter must be a pointer (`&myVar`).
- Thread-safe for concurrent access.