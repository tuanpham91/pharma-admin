package dbconnector

import (
	"recordType"
	"database/sql"
	"fmt"
)

var addUpdateInventoryQuery = "UPDATE med_inventory SET quantity = quantity + ? WHERE name = ? AND expirationDate = ? "

var addInventoryQuery = "INSERT INTO med_inventory (name, quantity, expirationDate) VALUES (?, ?, ?)"

var checkInventoryQuery = "SELECT * FROM med_record where name = ? and expirationDate = ?"
var checkIfItemExistsInventory = "SELECT * FROM med_inventory WHERE name = ? AND expirationDate = ?"

var checkInventoryQueryByName = "SELECT * FROM med_inventory where name = ?"

func CheckIfItemExistsInInventory(name string, date string) bool {
	rows, _ := Database.Query(checkIfItemExistsInventory, name, date)
	res := RowsToInventoryItems(rows)
	return (len(res) > 0 )
}

func UpdateInventoryInDatabase(name string, mhd string, change int) {
	statement, _ := Database.Prepare(addUpdateInventoryQuery)
	statement.Exec(change, name, mhd)
}

func AddInventoryToDatabase(inv_record recordType.ItemInventory) {
	statement, err := Database.Prepare(addInventoryQuery)
	if (err != nil) {
		fmt.Println(err)
	}
	statement.Exec(inv_record.Name, inv_record.Quantity, inv_record.ExpirationDate)
}

// A Method to check the inventory of an item with a certain best-before-date
// The return value is the total amount of that item
// - name: The name of the item
// - date: The best-before date of the item
func CheckInventoryForItemTotalNumber(name string, date string) int {
	rows, _ := Database.Query(checkIfItemExistsInventory, name, date)
	res := RowsToInventoryItems(rows)
	var total int = 0
	for _, item  := range res {
		fmt.Printf("Found an inventory: %d \n ", item.Quantity)
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

// A Method to check the inventory of an item, regardless of their best-before-date
// The return value is the total amount of that item
// - name: The name of the item
func CheckInventoryForAllDates(name string) int {
	rows, _ := Database.Query(checkInventoryQueryByName, name)
	res := RowsToInventoryItems(rows)
	var total int = 0
	for _, item  := range res {
		fmt.Printf("Found an inventory: %d \n ", item.Quantity)
		total += item.Quantity
	}
	return total
}