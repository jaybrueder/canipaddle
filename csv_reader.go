package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
)

var (
	lastValue string
	mu        sync.RWMutex
)

// getLastValueFromCSV reads the CSV file and returns the last field of the last record.
func getLastValueFromCSV(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var lastRecord []string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", fmt.Errorf("failed to read CSV record: %w", err)
		}
		lastRecord = record
	}

	if len(lastRecord) == 0 {
		return "", fmt.Errorf("CSV file is empty or has no records")
	}

	// Return the last field of the last record
	return lastRecord[len(lastRecord)-1], nil
}

// updateLastValue periodically reads the CSV and updates the lastValue.
// For this example, it's called once at startup. In a real app,
// you might run this in a goroutine with a ticker.
func updateLastValue(filePath string) error {
	val, err := getLastValueFromCSV(filePath)
	if err != nil {
		return fmt.Errorf("failed to get last value from CSV: %w", err)
	}

	mu.Lock()
	lastValue = val
	mu.Unlock()
	fmt.Printf("Updated lastValue to: %s\n", val)
	return nil
}

// GetStoredValue returns the currently stored last value.
func GetStoredValue() string {
	mu.RLock()
	defer mu.RUnlock()
	return lastValue
}
