// The Persistence package provides functions to persists wlog data.
package persistence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"s-stark.net/code/wlog/types"
)

// Read data from json nfile.
func Read(file string) (types.Week, error) {
	var week types.Week

	b, err := ioutil.ReadFile(file)

	if err == nil {
		err = json.Unmarshal(b, &week)

		if err != nil {
			return week, fmt.Errorf("Failed to unmarshal data from %v, %w", file, err)
		}
	}

	return week, nil
}

// write data to json file.
func Write(file string, week types.Week) error {
	dir := filepath.Dir(file)

	info, err := os.Stat(dir)

	if err != nil || info.IsDir() {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			fmt.Errorf("Failed to create data directory %v, %w", dir, err)
		}
	}

	b, err := json.MarshalIndent(week, "", " ")

	if err != nil {
		return fmt.Errorf("Failed to marshal data, %w", err)
	}

	err = ioutil.WriteFile(file, b, 0644)

	if err != nil {
		return fmt.Errorf("Failed to write data to %v, %w", file, err)
	}

	return nil
}
