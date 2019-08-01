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
var initializeInventoryTable = "CREATE TABLE IF NOT EXISTS med_inventory (id INTEGER PRIMARY KEY, name STRING, quantity INTEGER, expirationDate STRING)"

var addRecordQuery = "INSERT INTO med_record (name, quantity, price, expirationDate, dateOfRecord) VALUES (?,?,?,?,?) "
var truncateTableQuery = "DELETE FROM med_record; DELETE FROM med_inventory"

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

func RowsToRecord(rows *sql.Rows) []recordType.Record {
	var res []recordType.Record
	if (rows == nil) {
		fmt.Println("This is an empty record")
		return nil 
	}
	for rows.Next() {
		var r recordType.Record
		var i int 
		rows.Scan(&i, &r.Name, &r.Quantity, &r.Price, &r.ExpirationDate, &r.DateOfRecord)
		res = append(res, r)
	}
	return res
}
