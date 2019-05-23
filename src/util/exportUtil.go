package util 

import (
	"os"
	"encoding/csv"
	"recordType"
)

func WriteRecordToCSVFile(r recordType.Record, destination string) {
	file, err := os.Create(destination)
}