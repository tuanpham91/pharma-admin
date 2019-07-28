package dbconnector

import (
	"recordType"
	"testing"
	"strconv"
	"database/sql"
	"os"
	"fmt"
)

var pathToTestDatabase = "./test.db"

func TestMain(m *testing.M) {
	var err error
	Database, err = sql.Open("sqlite3", pathToTestDatabase)
	GenerateDatabases(pathToDatabase)
	fmt.Print(err)
	code := m.Run()	
	os.Remove(pathToTestDatabase)
	os.Exit(code)
}

func TestCalculatingInventory(t *testing.T) {
	TruncateTable()
	AddRecordToDatabase(recordType.Record{1, "Paracetamol", 10 , 109000, "1.1.2020", "17.7.2019"},"test.db")	
	AddRecordToDatabase(recordType.Record{1, "Paracetamol", 20 , 99000, "1.1.2020", "17.7.2019"},"test.db")	
	result := CheckInventory("Paracetamol","test.db")
	t.Log("The Number of item in test :" + strconv.Itoa(result))
	if (result != 300) {
		t.Errorf(" The result was wrong")
	}
}