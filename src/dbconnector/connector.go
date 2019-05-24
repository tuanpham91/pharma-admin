package dbconnector
import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"recordType"
	"fmt"
)
var databaseName = "record"
var pathToDatabase = "./record.db"
var addRecordQuery = "INSERT INTO med_record (name, quantity, price, expirationDate, dateOfRecord) VALUES (?,?,?,?,?) "

func GenerateDatabase(name string) {
	database, _ := sql.Open("sqlite3", pathToDatabase)
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS med_record (id INTEGER PRIMARY KEY, name TEXT, quantity INTEGER, price REAL, expirationDate TEXT, dateOfRecord TEXT)")
	statement.Exec()
}

func AddRecordToDatabase(record recordType.Record) {
	database, _ := sql.Open("sqlite3", pathToDatabase)
	statement, _ := database.Prepare(addRecordQuery)
	statement.Exec(record.Name, record.Quantity, record.Price, record.ExpirationDate, record.DateOfRecord)
}

func SelectQueryFromDB(query string) *sql.Rows {
	database, _  := sql.Open("sqlite3", pathToDatabase)
	rows, _ := database.Query(query)
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