package main

import (
	"fmt"
	_ "util"
	_ "time"
	"recordType"
	"dbconnector"
	"database/sql"
)

var pathToDatabase = "./record.db"

func main() {
	dbconnector.Database, _ = sql.Open("sqlite3", pathToDatabase)
	fmt.Println("Welcome to Smart Pharma")
	// st := int64(time.Now().Unix())
	dbconnector.GenerateDatabases()
	// rec := createARecord("tuan",10,99.99,"1.1.2010", "2.2.2020")
	// dbconnector.AddRecordToDatabase(rec)
	// fmt.Println(util.ConvertToFormattedDate(st))
	//rows := dbconnector.SelectQueryFromDB("Select * From med_record")
	//res := dbconnector.RowsToRecord(rows)
	//fmt.Println(res)
}

func createARecord(name string, quantity int, price float32, expirationDate string, dateOfRecord string) recordType.Record {
	return recordType.Record{0, name, quantity, price, expirationDate, dateOfRecord}
}


