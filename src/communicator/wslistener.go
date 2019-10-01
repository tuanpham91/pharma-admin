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
	/*
	if (r.Method != "GET") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	*/
	//Check if content is JSON
	if (!strings.Contains(r.Header.Get("Content-Type"),"json")){
		// TODO : Revert this
		//http.Error(w, "A json format is required", http.StatusBadRequest)
		//return
	}	
	

	// TODO : Check request Body for DBFilter:
	requestBody, err := ioutil.ReadAll(r.Body)
	fmt.Printf("Receive a request to Record with content :%s\n", string(requestBody))
	var query string
	if (err != nil) {
		query = dbconnector.BaseQueryBuilder("med_record")
		fmt.Printf(err.Error())
	} else {
		// Convert Request Body to array of DBFilter
		filters := getDBFilterFromJsonPayload(requestBody)
		//fmt.Printf("%v",filters)
		// TODO Refactor this 
		query = dbconnector.BaseQueryBuilder("med_record", filters...)
	}

	w.Header().Set("Content-Type", "application/json")
	results := dbconnector.GetRecordDataFromDBWithFilter(query)
	json.NewEncoder(w).Encode(results)
}


func getDBFilterFromJsonPayload(payload []byte) []dbconnector.DBFilter {
	var filters []dbconnector.DBFilter
	err := json.Unmarshal(payload, filters)
	if (err != nil) {
		return make([]dbconnector.DBFilter, 0)
	} else {
		return filters
	}
}

func getInventoryWithQuery(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	fmt.Printf("Receive a request to Inventory \n")
	//Check if content is JSON
	if (!strings.Contains(r.Header.Get("Content-Type"),"json")){
		// TODO : Revert this
		//http.Error(w, "A json format is required", http.StatusBadRequest)
		//return
	}	
	
	w.Header().Set("Content-Type", "application/json")
	query := dbconnector.BaseQueryBuilder("med_inventory")
	results := dbconnector.GetInventoryDataFromDBWithFilter(query)
	json.NewEncoder(w).Encode(results)
}

func StartWebserver() {
	fmt.Printf("Start Webserver at local host and port 8080 right now\n")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}