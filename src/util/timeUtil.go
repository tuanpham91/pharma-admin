package util

import (
	"time"
)

var timeFormat = "02-01-2006"

func GetFormattedCurrentTime() string {
	return time.Now().Format(timeFormat)
}

func ConvertToFormattedDate(timeStamp int64) string {
	return time.Unix(timeStamp, 0).Format(timeFormat)
}
