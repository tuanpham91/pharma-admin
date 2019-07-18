package dbconnector

import (
	"recordType"
	"testing"
	"strconv"
)

func TestGenerateDatabase(t *testing.T) {
	GenerateDatabases("test.db")
}

func TestCalculatingInventory(t *testing.T) {
	GenerateDatabases("test.db")
	AddRecordToDatabase(recordType.Record{1, "Paracetamol", 10 , 109000, "1.1.2020", "17.7.2019"},"test.db")	
	AddRecordToDatabase(recordType.Record{1, "Paracetamol", 20 , 99000, "1.1.2020", "17.7.2019"},"test.db")	
	result := CheckInventory("Paracetamol","test.db")
	t.Log("The Number of item in test :" + strconv.Itoa(result))
	if (result != 30) {
		t.Errorf(" The result was wrong")
	}
}