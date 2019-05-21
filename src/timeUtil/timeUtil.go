package timeUtil

import (
	"time"
)

var timeFormat = "02-01-2006"

func GetFormattedCurrentTime() string {
	return time.Now().Format(timeFormat)
}
