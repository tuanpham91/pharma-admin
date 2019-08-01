package dbconnector
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"recordType"
	"fmt"
)

var databaseName = "record"
var pathToDatabase = "./record.db"
var recordTableName = "med_record"
var inventoryTableName = "med_inventory"

var initializeRecordTable = "CREATE TABLE IF NOT EXISTS med_record (id INTEGER PRIMARY KEY, name TEXT, quantity INTEGER, price REAL, expirationDate TEXT, dateOfRecord TEXT)"
var initializeInventoryTable = "CREATE TABLE IF NOT EXISTS med_inventory (id INTEGER PRIMARY KEY, name TEXT, quantity INTEGER, expirationDate STRING)"

var addRecordQuery = "INSERT INTO med_record (name, quantity, price, expirationDate, dateOfRecord) VALUES (?,?,?,?,?) "
var truncateTableQuery = "DELETE FROM med_record; DELETE FROM med_inventory"

var updateInventoryQuery = "UPDATE med_inventory WHERE name = ? AND expirationDate = ? SET quantity = quantity + ?"
var checkInventoryQuery = "SELECT * FROM med_record where name = ? and expirationDate = ?"
var addInventoryQuery = "INSERT INTO med_inventory (name, quantity, expirationDate) VALUES (?, ?, ?)"

var Database *sql.DB

func GenerateDatabases(targetDb ...string) {
	var dbPath string
	if (len(targetDb) > 0) {
		dbPath = targetDb[0]
	} else {
		dbPath = databaseName
	}
	RunQuery(initializeInventoryTable, dbPath)
	RunQuery(initializeRecordTable, dbPath)
}

func RunQuery(query string, dbPath string) {
	statement, err := Database.Prepare(query)
	if (err != nil) {
		fmt.Print(err)
	}
	statement.Exec()
}

func TruncateTable() {
	statement, _ := Database.Prepare(truncateTableQuery)
	statement.Exec()
}

func AddRecordToDatabase(record recordType.Record, dbPath string) {
	statement, _ := Database.Prepare(addRecordQuery)
	statement.Exec(record.Name, record.Quantity, record.Price, record.ExpirationDate, record.DateOfRecord)
}

func UpdateInventoryInDatabase(product string, mhd string, change int) {
	statement, _ := Database.Prepare(updateInventoryQuery)
	statement.Exec(product, mhd, change)
}

func AddInventoryToDatabase(inv_record recordType.ItemInventory) {
	statement, _ := Database.Prepare(addInventoryQuery)
	statement.Exec(inv_record.Name, inv_record.Quantity, inv_record.ExpirationDate)
}

func RowsToRecord(rows *sql.Rows) []recordType.Record {
	var res []recordType.Record
	if (rows == nil) {
		fmt.Println("This is an empty record")
		return nil 
	}
	for rows.Next() {
		var r recordType.Record
		var i int 
		err := rows.Scan(&i, &r.Name, &r.Quantity, &r.Price, &r.ExpirationDate, &r.DateOfRecord)
		if (err != nil ) {
			fmt.Print(err)
		}
		res = append(res, r)
	}
	return res
}

func RowsToInventoryItems(rows *sql.Rows) []recordType.ItemInventory {
	var res []recordType.ItemInventory
	if (rows == nil) {
		fmt.Println("This is an empty record")
		return nil 
	}
	for rows.Next() {
		var r recordType.ItemInventory
		var i int 
		err := rows.Scan(&i, &r.Name, &r.Quantity, &r.ExpirationDate)
		if (err != nil ) {
			fmt.Print(err)
		}
		res = append(res, r)
	}
	return res
}

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