package dbconnector

// This class provides query builder which is neccessary in case some filter should be applied.
import "strings"
import "recordType"

type DBFilter struct {
	Attribute	string
	Value		string
}

func BaseQueryBuilder (tableName string, filters []DBFilter) {
	var sb strings.Builder
	sb.WriteString("SELECT * FROM " + tableName)
	if (len(filters) > 0) {
		sb.WriteString(" WHERE")
	}
	for _, f := range filters {
		sb.WriteString(f.Attribute + " " + f.Value)
	}
}

func GetRecordDataFromDBWithFilter(query string) []recordType.Record {
	var records = RowsToRecord(ExecuteQuery(query))
	return records
}

func GetInventoryDataFromDBWithFilter(query string) []recordType.ItemInventory {
	var records = RowsToInventoryItems(ExecuteQuery(query))
	return records
}