package utils

import (
	"fmt"
	"time"
)

var StartTime time.Time

func Uptime() string {
	uptime := time.Since(StartTime)
	days := int(uptime / (24 * time.Hour))
	hours := int((uptime % (24 * time.Hour)) / time.Hour)
	minutes := int((uptime % time.Hour) / time.Minute)
	seconds := int((uptime % time.Minute) / time.Second)
	formattedDuration := fmt.Sprintf("%ddays-%dhours-%dminutes-%dseconds", days, hours, minutes, seconds)
	return formattedDuration
}

func init() {
	StartTime = time.Now()
}
