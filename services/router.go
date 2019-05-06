package services

import (
	"fmt"
	"log"
	"net/http"
	"organization/datasources"
	"os"

	"github.com/gorilla/mux"
)

func MapRoutes() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	filePath := fmt.Sprintf("%v/%v", dir, "datasources/resources/organizations.csv")
	datasource := datasources.NewOrganizationDataRetrieverForCsv(filePath)
	retriever := NewSearchOrganizationsHandler(datasource)
	port := 3000

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/organizations", retriever.SearchOrganizations)
	fmt.Printf("Starting api on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
