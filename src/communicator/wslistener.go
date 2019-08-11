package communicator
// security is not required, because local.
import (
	"net/http"
	"fmt"
	"recordType"
	"encoding/json"
	"strings"
	"dbconnector"
)

var (
	allRecords		="/record"
	addRecord		="/addrecord"
	removeRecord	="/rmrecord"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage")
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/record", getRecordsWithQuery)
	http.HandleFunc(addRecord, addRecordRequestRoute)
}

func getRecordRoute(w http.ResponseWriter, r *http.Request) {
	//TODO read arguements from request
	testRecord := recordType.ItemInventory{1, "Test", 1, "1.1.2020"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testRecord)
}

// To add a new record to database, must be POST
func addRecordRequestRoute(w http.ResponseWriter, r *http.Request) {
	if (r.Method == "GET") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
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

// There should be some criterias to filter 
func getRecordsWithQuery(w http.ResponseWriter, r *http.Request) {
	if (r.Method != "GET") {
		http.Error(w, "This type of Request is not allowed", http.StatusBadRequest)
		return
	} 
	//Check if content is JSON
	if (!strings.Contains(r.Header.Get("Content-Type"),"json")){
		//http.Error(w, "A json format is required", http.StatusBadRequest)
		//return
	}	
	w.Header().Set("Content-Type", "application/json")
	query := dbconnector.BaseQueryBuilder("med_record")
	results := dbconnector.GetRecordDataFromDBWithFilter(query)
	json.NewEncoder(w).Encode(results)
}

func StartWebserver() {
	fmt.Printf("Start Webserver at local host and port 8080 \n")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}