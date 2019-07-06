package dbconnector

import (
	"testing"
	"fmt"
)

func TestGenerateDatabase(t *testing.T) {
	dbconnector.GenerateDatabases()
	fmt.Println("Is this not great ?")
}

func TestAddRecordToDatabase(t *testing.T) {
	dbconnector.GenerateDatabases()
}