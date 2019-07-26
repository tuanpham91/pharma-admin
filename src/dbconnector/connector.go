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
var truncateTableQuery = "DELETE FROM ?"

var updateInventoryQuery = "UPDATE med_inventory WHERE name = ? AND expirationDate = ? SET quantity = quantity + ?"
var checkInventoryQuery = "SELECT * FROM med_record where name = ? and expirationDate = ?"

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

func UpsertInventoryRecord(re recordType.ItemInventory, dbPath string) {
	
}

func CheckIfInventoryRecordExists(re recordType.ItemInventory, dbPath string) bool {
	statement, _ := Database.Prepare(checkInventoryQuery)
	rows, _ := statement.Query(re.Name, re.ExpirationDate)
	return (len(RowsToRecord(rows)) >= 1)
}

func TruncateTable(tableName string, dbPath string) {
	statement, err := Database.Prepare(truncateTableQuery)
	if (err != nil) {
		fmt.Print(err)
	}
	statement.Exec(tableName)
}

func AddRecordToDatabase(record recordType.Record, dbPath string) {
	statement, _ := Database.Prepare(addRecordQuery)
	statement.Exec(record.Name, record.Quantity, record.Price, record.ExpirationDate, record.DateOfRecord)
}

func SelectQueryFromDB(query string) *sql.Rows {
	rows, _ := Database.Query(query)
	return rows
}

func RowsToRecord(rows *sql.Rows) []recordType.Record {
	var res []recordType.Record
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

func SubtractARecord(r recordType.Record) {
	
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