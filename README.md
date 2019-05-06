TreeTop Commons Organization Service
==========
This service is used for querying organizations by the following filters:
* id
* name
* city
* state
* postal
* category
* orderby (value = field name listed above)
* direction (values = ASC, DSC)

# Desired Improvements
* move running port and .csv filepath to environment vars
* add dependency injection
* add unit tests
* if data used is changed often, ideally move the data source into a database or other data storage
* if csv is the data source format and is edited often, can add AWS Athena integration to query
* use reflection for sorting by field/direction as right now code checks every field in the search request
* more refactoring such as separating filtering from data retrieval with interfaces 
* move adding response Content-Type header to middleware
* add api versioning "v1"
* add logging
* add security/token auth
* add to build pipeline (i.e. CircleCi)
* add to deployment pipeline
* add health check if needed

# Problem Design And Process
* main potential problem areas are data extraction from .csv and filtering logic
* started first with attempting to extract the data from the .csv file via http.Get
* due to the line endings, these requests in addition to straight copy/read of the .csv to the local file system resulted in a data set with only the last organization listed being returned
* to solve, downloaded a copy of the .csv and manually returned each line - this resolved the issue as now file read recognized each end of line
* with data successfully retieved, added organization and search request dto models
* added struct with loading and caching of the data from the .csv file in the constructor
* added functionality to perform desired filtering on the struct
* added struct to handle http request and response - gets querystring values and creates search request dto to pass to the filterer
* connected http handler struct to filter struct
* added main.go entry with routing on the desired port
* http handler uses filtering interface so the struct that performs the actual filtering can be replaced if the datasource is changed

# Getting Started
## Before Running the Api
* golang will need to be install on the local machine
* installation documentation `https://golang.org/doc/install`
* git will need to be installed on the local machine
* installation documentation `https://git-scm.com/book/en/v2/Getting-Started-Installing-Git`
* get the repository running this command in a chosen directory `git clone https://github.com/peggyhouseman/treetopcommons.git`
## Building and Running
* navigate to the repository's root directory
* run in terminal `go run main.go`
* api is now running on 127.0.0.1:3000
## Sample Request
* GET `http://127.0.0.1:3000/organizations?category=education&orderby=name&direction=dsc`
