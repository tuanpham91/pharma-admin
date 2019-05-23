package main

import (
	"fmt"
	"util"
	"time"
	"recordType"
	"dbconnector"
)


func main() {
	fmt.Println("Welcome to Smart Pharma")
	st := int64(time.Now().Unix())
	
	fmt.Println(util.ConvertToFormattedDate(st))
}

func createARecord(name string, quantity int, price float32, expirationDate int64, dateOfRecord int64) recordType.Record {
	return recordType.Record{name, quantity, price, expirationDate, dateOfRecord}
}
