package communicator

import (
	"net/http"
	"fmt"
	"recordType"
	"encoding/json"
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
	http.HandleFunc("/record",getRecordRoute )
}

func getRecordRoute(w http.ResponseWriter, r *http.Request) {
	//TODO read arguements from request
	testRecord := recordType.ItemInventory{1, "Test", 1, "1.1.2020"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(testRecord)
}

func StartWebserver() {
	fmt.Printf("Start Webserver at local host and port 8080 \n")
	setUpRoutes()
	http.ListenAndServe(":8080", nil)
}