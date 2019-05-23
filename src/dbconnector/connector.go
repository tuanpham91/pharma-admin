package dbconnector
import (
	"database/sql"
	"github.com/mattn/go-sqlite3"
)

func generateDatabase(name string) {
	database, _ := sql.Open("sqlite3", "./record.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
}