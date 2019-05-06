package converters

import (
	"errors"
	"fmt"
	"organization/models"
	"strconv"
)

func ConvertCsvToOrganization(csvData [][]string) ([]models.Organization, error) {
	orgs := []models.Organization{}
	for i, record := range csvData {

		// skip headers
		if i == 0 {
			continue
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, errors.New(fmt.Sprintf("Error converting %n row, id column value: %v", i, record[0]))
		}

		orgs = append(orgs, models.Organization{
			Id:       id,
			Name:     record[1],
			City:     record[2],
			State:    record[3],
			Postal:   record[4],
			Category: record[5],
		})
	}
	return orgs, nil
}
