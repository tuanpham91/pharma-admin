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

var initializeRecordTable = "CREATE TABLE IF NOT EXISTS med_record (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, quantity INTEGER, price REAL, expirationDate TEXT, dateOfRecord TEXT)"
var initializeInventoryTable = "CREATE TABLE IF NOT EXISTS med_inventory (id INTEGER PRIMARY KEY AUTOINCREMENT, name STRING, quantity INTEGER, expirationDate STRING)"

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
	var err error
	Database, err = sql.Open("sqlite3", dbPath)
	if (err != nil) {
		fmt.Println(err)
	}
	RunQuery(initializeInventoryTable)
	RunQuery(initializeRecordTable)
}

func RunQuery(query string) {
	statement, err := Database.Prepare(query)
	if (err != nil) {
		fmt.Print(err)
	}
	statement.Exec()
}

func ExecuteQuery(query string) *sql.Rows {
	rows, _ := Database.Query(query)
	return rows;
}

func TruncateTable() {
	statement, _ := Database.Prepare(truncateTableQuery)
	statement.Exec()
}

func AddRecordToDatabase(record recordType.Record) {
	statement, _ := Database.Prepare(addRecordQuery)
	statement.Exec(record.Name, record.Quantity, record.Price, record.ExpirationDate, record.DateOfRecord)
}

// A Function to handle every record that arrives
// The workflow works like this 
// A Record Arrives -> Check if already in inventory database with the same date -> If no then create a new one
// If yes then update the old one
// Record database should always be inserted.
func HandleRecord(record recordType.Record) {
	AddRecordToDatabase(record)
	existed := CheckIfItemExistsInInventory(record.Name, record.ExpirationDate)
	if (existed) {
		UpdateInventoryInDatabase(record.Name, record.ExpirationDate, record.Quantity)
	} else {
		AddInventoryToDatabase(recordType.ItemInventory{1, record.Name, record.Quantity, record.ExpirationDate})
	}
}

func RowsToRecord(rows *sql.Rows) []recordType.Record {
	var res []recordType.Record
	if (rows == nil) {
		fmt.Println("This is an empty record")
		return nil 
	}
	for rows.Next() {
		var r recordType.Record
		rows.Scan(r.Id, &r.Name, &r.Quantity, &r.Price, &r.ExpirationDate, &r.DateOfRecord)
		res = append(res, r)
	}
	return res
}
