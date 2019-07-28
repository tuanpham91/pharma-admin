package dbconnector

var checkIfItemExistsInventory = "SELECT * FROM med_record WHERE name = ? AND date = ?"

func CheckIfItemExistsInInventory(name string, date string) bool {
	statement, _ := Database.Prepare(checkIfItemExistsInventory)
	rows, _ := statement.Query(name, date)
	return (rows != nil)
}