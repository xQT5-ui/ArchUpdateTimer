package lib

import "time"

func GetDayOfWeek() time.Weekday {
	nowTime := time.Now()
	return nowTime.Weekday()
}
