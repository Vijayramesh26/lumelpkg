package config

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"lumelpkg/common"
	"lumelpkg/utils"
	"path/filepath"
	"strings"
	"sync"

	"github.com/BurntSushi/toml"
)

func Init(logger *utils.Logger) {
	logger.Log(common.DEBUG, "3", "LoadTOMLFile (+) ")
	// Load configs from the directory on server start
	LoadAllTOMLConfigs("./toml")
	logger.Log(common.DEBUG, "3", "LoadTOMLFile (-) ")
}

var (
	// ConfigMap stores all TOML configs with the filename (without .toml) as key
	ConfigMap = make(map[string]map[string]any)

	// mu ensures thread-safe access to ConfigMap
	mu sync.RWMutex

	// once ensures config is loaded only once
	once sync.Once
)

// LoadAllTOMLConfigs loads all .toml files from the given folder into ConfigMap
func LoadAllTOMLConfigs(folderPath string) {
	once.Do(func() {
		err := filepath.WalkDir(folderPath, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if strings.HasSuffix(d.Name(), ".toml") {
				var data map[string]any
				if _, err := toml.DecodeFile(path, &data); err != nil {
					log.Fatalf("Failed to decode TOML file %s: %v", path, err)
				}

				filename := strings.TrimSuffix(d.Name(), ".toml")

				mu.Lock()
				ConfigMap[filename] = data
				mu.Unlock()
			}
			return nil
		})

		if err != nil {
			log.Fatalf("Error reading TOML config files from folder %s: %v", folderPath, err)
		}
	})
}

// GetConfig returns the entire config map for a given TOML filename (without extension)
func GetConfig(filename string) (map[string]any, bool) {
	mu.RLock()
	defer mu.RUnlock()
	data, ok := ConfigMap[filename]
	return data, ok
}

/* GetAndAssignTomlValue retrieves a value from a TOML config file (either top-level or nested)
and assigns it to the provided output variable.

Parameters:
  - filename: The name of the TOML file (without `.toml` extension).
  - key: A dot-separated path to the desired key (e.g., "database.host", "appname").
  - out: A pointer to the variable where the result should be stored.

This method uses JSON marshaling/unmarshaling to convert the raw TOML value into the desired type.
The `out` parameter must be a pointer to the expected Go type.

Example TOML:
  appname = "MyApp"
  [database.credentials]
  username = "admin"
  password = "secret"

Example usage:
  var appName string
  err := GetAndAssignTomlValue("app", "appname", &appName)

  var username string
  err := GetAndAssignTomlValue("app", "database.credentials.username", &username)

  var creds struct {
      Username string
      Password string
  }
  err := GetAndAssignTomlValue("app", "database.credentials", &creds)

Returns:
  - nil on success
  - error if file/key not found or type mismatch
*/

// GetAndAssignTomlValue retrieves a value (top-level or nested) from the TOML config
// and assigns it to the output variable (must be a pointer).
func GetAndAssignTomlValue(filename, key string, out any) error {
	mu.RLock()
	defer mu.RUnlock()

	data, ok := ConfigMap[filename]
	if !ok {
		return fmt.Errorf("config file not found: %s", filename)
	}

	keys := strings.Split(key, ".")
	var current any = data

	for _, k := range keys {
		m, ok := current.(map[string]any)
		if !ok {
			return fmt.Errorf("intermediate value is not a map at key: %s", k)
		}
		current, ok = m[k]
		if !ok {
			return fmt.Errorf("key not found: %s", k)
		}
	}

	// Convert to desired type using JSON marshal/unmarshal
	bytes, err := json.Marshal(current)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	if err := json.Unmarshal(bytes, out); err != nil {
		return fmt.Errorf("unmarshal error: %w", err)
	}

	return nil
}
