// Package internal provides utility functions for internal operations
// such as loading and parsing JSON files.
package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var (
	// jsonCache stores parsed JSON data to avoid repeated file reads and parsing
	jsonCache = make(map[string]any)
	// jsonCacheSync provides thread-safe access to the jsonCache
	jsonCacheSync sync.RWMutex
)

// LoadJsonFile loads and parses a JSON file from the specified package and locale
// It caches the results to improve performance for subsequent calls
func LoadJsonFile(packageName, locale, fileName string) (map[string]any, error) {
	cacheKey := fmt.Sprintf("%s:%s:%s", packageName, locale, fileName)

	jsonCacheSync.RLock()
	if cachedData, ok := jsonCache[cacheKey]; ok {
		jsonCacheSync.RUnlock()
		return cachedData.(map[string]any), nil
	}
	jsonCacheSync.RUnlock()

	filePath := filepath.Join("..", "locales", locale, packageName, fileName)

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}

	var result map[string]any
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("failed to parsing json of file %s: %w", fileName, err)
	}

	jsonCacheSync.Lock()
	jsonCache[cacheKey] = result
	jsonCacheSync.Unlock()

	return result, nil
}

// GetStringSlice extracts a string slice from a map by key
// Returns an empty slice if the key doesn't exist or the value is not a slice of strings
func GetStringSlice(data map[string]any, key string) []string {
	if data == nil {
		return []string{}
	}

	value, ok := data[key]
	if !ok {
		return []string{}
	}

	anySlice, ok := value.([]any)
	if !ok {
		return []string{}
	}

	slice := make([]string, len(anySlice))
	for i, v := range anySlice {
		slice[i] = v.(string)
	}

	return slice
}
