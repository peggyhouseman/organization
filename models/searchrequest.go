package models

import "strings"

type SortOrder int

const (
	ASC SortOrder = iota
	DSC
)

type OrganizationSearchRequest struct {
	Id       int
	Name     string
	City     string
	State    string
	Postal   string
	Category string
	OrderBy  string
	Sorting  SortOrder
}

func (r *OrganizationSearchRequest) HasSearchParam() bool {
	return r.Id > 0 || r.Name != "" || r.City != "" || r.State != "" || r.Postal != "" || r.Category != ""
}

func (r *OrganizationSearchRequest) IsMatching(data Organization) bool {

	if r.Id > 0 && r.Id != data.Id {
		return false
	}

	if r.Name != "" && !strings.EqualFold(r.Name, data.Name) {
		return false
	}

	if r.City != "" && !strings.EqualFold(r.City, data.City) {
		return false
	}

	if r.State != "" && !strings.EqualFold(r.State, data.State) {
		return false
	}

	if r.Postal != "" && !strings.EqualFold(r.Postal, data.Postal) {
		return false
	}

	if r.Category != "" && !strings.EqualFold(r.Category, data.Category) {
		return false
	}

	return true
}
