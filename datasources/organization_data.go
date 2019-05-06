package datasources

import (
	"errors"
	"fmt"
	"organization/datasources/converters"
	"organization/datasources/utils"
	"organization/models"
	"sort"
	"strings"
)

type OrganizationDataRetrieverForCsv struct {
	organizations []models.Organization
}

func NewOrganizationDataRetrieverForCsv(filePath string) OrganizationDataRetrieverForCsv {
	orgs, err := getAllOrganizationsFromCsv(filePath)
	if err != nil {
		panic(errors.New(fmt.Sprintf("failed to initialize org data retriever: %v", err)))
	}
	return OrganizationDataRetrieverForCsv{
		organizations: orgs,
	}
}

func getAllOrganizationsFromCsv(filePath string) ([]models.Organization, error) {
	data, err := utils.GetCsvDataFromFile(filePath)
	if err != nil {
		return nil, err
	}
	orgs, err := converters.ConvertCsvToOrganization(data)
	if err != nil {
		return nil, err
	}
	return orgs, nil
}

func (r OrganizationDataRetrieverForCsv) Search(req models.OrganizationSearchRequest) []models.Organization {
	result := []models.Organization{}

	if !req.HasSearchParam() {
		return result
	}

	for _, org := range r.organizations {
		if req.IsMatching(org) {
			result = append(result, org)
		}
	}

	sort.Slice(result[:], func(i, j int) bool {
		switch strings.ToLower(req.OrderBy) {
		case "id":
			if req.Sorting == models.ASC {
				return result[i].Id < result[j].Id
			} else {
				return result[i].Id > result[j].Id
			}
		case "name":
			if req.Sorting == models.ASC {
				return result[i].Name < result[j].Name
			} else {
				return result[i].Name > result[j].Name
			}
		case "city":
			if req.Sorting == models.ASC {
				return result[i].City < result[j].City
			} else {
				return result[i].City > result[j].City
			}
		case "state":
			if req.Sorting == models.ASC {
				return result[i].State < result[j].State
			} else {
				return result[i].State > result[j].State
			}
		case "postal":
			if req.Sorting == models.ASC {
				return result[i].Postal < result[j].Postal
			} else {
				return result[i].Postal > result[j].Postal
			}
		default:
			if req.Sorting == models.ASC {
				return result[i].Category < result[j].Category
			} else {
				return result[i].Category > result[j].Category
			}
		}
	})

	return result
}
