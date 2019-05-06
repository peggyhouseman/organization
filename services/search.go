package services

import (
	"encoding/json"
	"net/http"
	"organization/models"
	"strconv"

	"github.com/labstack/gommon/log"
)

type OrganizationDataRetriever interface {
	Search(req models.OrganizationSearchRequest) []models.Organization
}

type SearchOrganizationsHandler struct {
	DataRetriever OrganizationDataRetriever
}

func NewSearchOrganizationsHandler(retriever OrganizationDataRetriever) SearchOrganizationsHandler {
	return SearchOrganizationsHandler{
		DataRetriever: retriever,
	}
}

func (s *SearchOrganizationsHandler) SearchOrganizations(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")
	id := 0
	if len(idStr) > 0 {
		idInt, err := strconv.Atoi(idStr)
		if err != nil {
			log.Error(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id = idInt
	}
	name := r.URL.Query().Get("name")
	city := r.URL.Query().Get("city")
	state := r.URL.Query().Get("state")
	postal := r.URL.Query().Get("postal")
	category := r.URL.Query().Get("category")
	orderby := r.URL.Query().Get("orderby")
	sort := r.URL.Query().Get("direction")
	sortOrder := models.ASC
	if sort == "dsc" {
		sortOrder = models.DSC
	}

	req := models.OrganizationSearchRequest{
		Id:       id,
		Name:     name,
		City:     city,
		State:    state,
		Postal:   postal,
		Category: category,
		OrderBy:  orderby,
		Sorting:  sortOrder,
	}

	resp := s.DataRetriever.Search(req)

	apiResp := models.ApiResponse{Organizations: resp}
	jsonResponse, err := json.Marshal(apiResp)

	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
