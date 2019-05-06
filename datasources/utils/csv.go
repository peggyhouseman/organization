package utils

import (
	"encoding/csv"
	"io"
	"os"
)

func GetCsvDataFromFile(filePath string) ([][]string, error) {
	f, err := os.Open(filePath)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return getCsvDataFromReader(f)
}

func getCsvDataFromReader(r io.Reader) ([][]string, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}
