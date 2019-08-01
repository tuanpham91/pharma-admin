package dbconnector

var checkIfItemExistsInventory = "SELECT * FROM med_record WHERE name = ? AND date = ?"

func CheckIfItemExistsInInventory(name string, date string) bool {
	rows, _ := Database.Query(checkIfItemExistsInventory, name, date)
	return (rows != nil)
}