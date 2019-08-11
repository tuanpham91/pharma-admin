package main

import (
	"fmt"
	_ "util"
	_ "time"
	"recordType"
	"communicator"
)

var pathToDatabase = "./record.db"

func main() {
	fmt.Println("Welcome to Smart Pharma")
	// Electron here we go
	communicator.StartWebserver()
}

func createARecord(name string, quantity int, price float32, expirationDate string, dateOfRecord string) recordType.Record {
	return recordType.Record{0, name, quantity, price, expirationDate, dateOfRecord}
}


