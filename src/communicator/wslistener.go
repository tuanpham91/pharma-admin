package communicator
// security is not required, because local.
import (
	"net/http"
	"fmt"
	"recordType"
	"encoding/json"
	"strings"
	"dbconnector"
	"io/ioutil"
)

var (
	allRecords		="/record"
	addRecord		="/addrecord"
	removeRecord	="/rmrecord"
	allInventory 	="/inventory"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc(allRecords, getRecordsWithQuery)
	http.HandleFunc(addRecord, addRecordRequestRoute)
	http.HandleFunc(allInventory, getInventoryWithQuery)
}

func getRecordRoute(w http.ResponseWriter, r *http.Request) {
	testRecord := recordType.ItemInventory{1, "Test", 1, "1.1.2020"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testRecord)
}


// To add a new record to database, must be POST
func addRecordRequestRoute(w http.ResponseWriter, r *http.Request) {
	/*
	if (r.Method == "GET") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	*/
	//Check if content is JSON
	if (!strings.Contains(r.Header.Get("Content-Type"),"json")){
		http.Error(w, "A json format is required", http.StatusBadRequest)
		return
	}	
	//Read arguments from request
	decoder := json.NewDecoder(r.Body)
	var record recordType.Record
	err := decoder.Decode(&record)
	if (err != nil) {
		http.Error(w, "An internal error occurred", http.StatusInternalServerError)
		return
	}
	dbconnector.HandleRecord(record)
	w.WriteHeader(http.StatusOK)
}

// TODO: There should be some criterias to filter 
// Function which handles Request for Records.
func getRecordsWithQuery(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	
	// TODO : Check request Body for DBFilter:
	query:= buildQueryWithFilters("med_record", r)
	w.Header().Set("Content-Type", "application/json")
	results := dbconnector.GetRecordDataFromDBWithFilter(query)
	fmt.Println("Request return with " + string(len(results)))
	json.NewEncoder(w).Encode(results)
}


func getDBFilterFromJsonPayload(payload []byte) []dbconnector.DBFilter {
	var filters [][] string
	err := json.Unmarshal(payload, &filters)
	if (err != nil) {
		fmt.Printf("Error while parsing DBFilters %s\n ", err.Error())
		return make([]dbconnector.DBFilter, 0)
	} else {
		// TODO : Loop through the array to convert to DBFilter
		var dbFilters =  make([]dbconnector.DBFilter, len(filters))
		for i:=0 ; i<len(filters); i++ {
			dbFilters[i] = dbconnector.DBFilter{filters[i][0], filters[i][1]}
		}
		return dbFilters
	}
}

func arrayToDBFilter(ar []string) dbconnector.DBFilter {
	return dbconnector.DBFilter{ar[0], ar[1]}
}

func getInventoryWithQuery(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "POST") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	fmt.Printf("Receive a request to Inventory \n")
	
	w.Header().Set("Content-Type", "application/json")
	query:= buildQueryWithFilters("med_record", r)
	results := dbconnector.GetInventoryDataFromDBWithFilter(query)
	json.NewEncoder(w).Encode(results)
}

func buildQueryWithFilters(tableName string, r *http.Request ) string {
	requestBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("Receive a request with content :%s\n", string(requestBody))
	var query string
	if (err != nil) {
		query = dbconnector.BaseQueryBuilder(tableName)
		fmt.Printf(err.Error())
	} else {
		// Convert Request Body to array of DBFilter
		filters := getDBFilterFromJsonPayload(requestBody)
		//fmt.Printf("%v",filters)
 		query = dbconnector.BaseQueryBuilder(tableName, filters...)
	}
	fmt.Printf("Getting Record with Query :%s \n", query)

	return query
}

func StartWebserver() {
	fmt.Printf("Start Webserver at local host and port 8080 right now\n")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}