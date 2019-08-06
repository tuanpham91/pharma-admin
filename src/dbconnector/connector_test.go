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
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeA", 10 , "1.1.2020" })	
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeA", 20 , "1.1.2022"})	
	result := CheckInventoryForAllDates("TypeA")
	t.Log("The Number of item in test :" + strconv.Itoa(result))
	if (result != 30) {
		t.Errorf(" The result was wrong")
	}
}

// Test if the function returns correctly when there is a drug in med_inventory and when there is not
func TestCheckingIfItemInInventory(t *testing.T) {
	TruncateTable()
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeB", 10, "1.1.2020"})
	result2 := CheckIfItemExistsInInventory("TypeB", "1.1.2020")
	if (result2 == false){
		t.Errorf("TestCheckingIfItemInInventory : Result is wrong, it should be : true")
	}
	result := CheckIfItemExistsInInventory("Tuan", "100")
	if (result == true) {
		t.Errorf("TestCheckingIfItemInInventory : Result is wrong, it should be : false")
	}
}

func TestCheckInventoryForItemTotalNumber(t *testing.T) {
	TruncateTable()
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeC", 10, "1.1.2020"})
	res := CheckInventoryForItemTotalNumber("TypeC","1.1.2020")
	if (res !=10) {
		t.Errorf("TestCheckInventoryForItemTotalNumber: The result is wrong : %d", res)
	}
}

func TestUpdateItemInvetory(t *testing.T) {
	TruncateTable()
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeD", 10, "1.1.2020"})
	UpdateInventoryInDatabase("TypeD","1.1.2020",20)
	res := CheckInventoryForItemTotalNumber("TypeD","1.1.2020")
	if (res !=30) {
		t.Errorf("TestUpdateItemInvetory: The result is wrong, why the hell should it be : %d",res)
	}
	
	TruncateTable()
	AddInventoryToDatabase(recordType.ItemInventory{1, "TypeG", 20, "1.1.2020"})
	UpdateInventoryInDatabase("TypeG","1.1.2020",-5)
	res2 := CheckInventoryForItemTotalNumber("TypeG","1.1.2020")
	if (res2 !=15) {
		t.Errorf("TestUpdateItemInvetory: The result is wrong, why the hell should it be : %d",res2)
	}
}

func TestHandleArrivingRecord (t *testing.T) {
	TruncateTable()
	record := recordType.Record{1, "Test", 10, 100000.00, "1.1.2022", "1.1.2020"}
	HandleRecord(record)
	result1 := CheckInventoryForAllDates(record.Name)
	if (result1 != 10) {
		t.Errorf("TestHandleArrivingRecord: The result is wrong, why the hell should it be : %d",result1)
	}
	record2 := recordType.Record{1, "Test", 12, 100000.00, "1.1.2023", "1.1.2020"}
	HandleRecord(record2)
	result2 := CheckInventoryForAllDates(record2.Name)
	if (result2 != 22) {
		t.Errorf("TestHandleArrivingRecord: The result is wrong, why the hell should it be : %d",result1)
	}
}