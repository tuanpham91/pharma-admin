package main

import (
	"fmt"
	_ "util"
	_ "time"
	"recordType"
	"communicator"
	"dbconnector"
)

var pathToDatabase = "./record.db"

func main() {

	fmt.Println("Welcome to Smart Pharma")
	dbconnector.GenerateDatabases(pathToDatabase)

	dbconnector.AddRecordToDatabase(recordType.Record{1, "Tuan", 1,11.111,"tuan","tuan"})
	dbconnector.AddInventoryToDatabase(recordType.ItemInventory{1, "TypeA", 10 , "1.1.2020" })	
	// Electron here we go
	communicator.StartWebserver()
}

func createARecord(name string, quantity int, price float32, expirationDate string, dateOfRecord string) recordType.Record {
	return recordType.Record{0, name, quantity, price, expirationDate, dateOfRecord}
}


