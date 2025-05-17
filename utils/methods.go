package utils

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/gocarina/gocsv"
)

// LoadCSV reads a CSV file with a custom delimiter and unmarshals its content
// into a slice of structs of type T. The CSV is expected to have a header row.
// Parameters:
//   - filename: path to the CSV file
//   - delimiter: rune character to use as field separator (e.g., ',', ';', '~')
//
// Returns:
//   - slice of T containing the unmarshaled CSV records
//   - error if any occurs during file reading or unmarshaling
func LoadCSV[T any](filename string, delimiter rune) ([]T, error) {
	// Open the CSV file in read-only mode
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Configure gocsv to use a custom CSV reader with specified delimiter
	// and enable LazyQuotes to handle fields with irregular quoting
	gocsv.SetCSVReader(func(r io.Reader) gocsv.CSVReader {
		cr := csv.NewReader(r)
		cr.Comma = delimiter
		cr.LazyQuotes = true
		return cr
	})

	var records []T

	// Unmarshal the CSV file content into the records slice
	// This expects the CSV header to match struct field tags
	if err := gocsv.UnmarshalFile(file, &records); err != nil {
		return nil, err
	}

	// Return the successfully parsed records
	return records, nil
}
