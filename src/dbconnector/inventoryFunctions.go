package dbconnector

import (
	"recordType"
	"database/sql"
	"fmt"
)

var addInventoryQuery = "UPDATE med_inventory SET quantity = quantity + ? WHERE name = ? AND expirationDate = ? "
var subtractInventoryQuery = "UPDATE med_inventory SET quantity = quantity - ? WHERE name = ? AND expirationDate = ? "

var checkInventoryQuery = "SELECT * FROM med_record where name = ? and expirationDate = ?"
var checkIfItemExistsInventory = "SELECT * FROM med_inventory WHERE name = ? AND expirationDate = ?"

func CheckIfItemExistsInInventory(name string, date string) bool {
	rows, _ := Database.Query(checkIfItemExistsInventory, name, date)
	res := RowsToInventoryItems(rows)
	return (len(res) > 0 )
}

func UpdateInventoryInDatabase(product string, mhd string, change int) {
	var query string
	var absChange int
	if (change >= 0) {
		query = addInventoryQuery
		absChange = change
	} else {
		query = subtractInventoryQuery
		absChange = change*(-1)
	}
	statement, _ := Database.Prepare(query)
	statement.Exec(absChange, product, mhd)
}

func AddInventoryToDatabase(inv_record recordType.ItemInventory) {
	statement, _ := Database.Prepare(addInventoryQuery)
	statement.Exec(inv_record.Name, inv_record.Quantity, inv_record.ExpirationDate)
}

func CheckInventoryForItemTotalNumber(name string, date string) int {
	rows, _ := Database.Query(checkIfItemExistsInventory, name, date)
	res := RowsToInventoryItems(rows)
	var total int = 0
	for _, item  := range res {
		total += item.Quantity
	}
	return total
}

func RowsToInventoryItems(rows *sql.Rows) []recordType.ItemInventory {
	var res []recordType.ItemInventory
	if (rows == nil) {
		fmt.Println("This is an empty inventory")
		return nil 
	}
	for rows.Next() {
		var r recordType.ItemInventory
		var i int 
		rows.Scan(&i, &r.Name, &r.Quantity, &r.ExpirationDate)
		res = append(res, r)
	}
	return res
}

// TODO: THis might not be necessary one day
func CheckInventory(name string, dbPathArgs ...string) int {
	var dbPath string
	if (len(dbPathArgs) > 0) {
		dbPath = dbPathArgs[0]
	} else {
		dbPath = databaseName
	}
	database, err  := sql.Open("sqlite3", dbPath)
	if (err != nil) {
		fmt.Println("Something is wrong while opening the database")
	}

	statement, err := database.Prepare(checkInventoryQuery)
	if (err != nil) {
		fmt.Println("Something is wrong while Preparing the query " + err.Error())
	}
	rows, err := statement.Query(name)
	if (err != nil) {
		fmt.Println("Something is wrong while executing the statement")
	}
	
	records := RowsToRecord(rows)
	fmt.Printf("There are %d rows", len(records))
	var res int = 0
	for _, e := range records {
		res += e.Quantity
	}
	return res
}