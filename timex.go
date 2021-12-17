package main

import (
	"time"
)

// 判断是否是同一天
func IsSameDay(timeOne time.Time, timeTwo time.Time) bool {
	if timeOne.Year() != timeTwo.Year() {
		return false
	}
	return timeOne.YearDay() == timeTwo.YearDay()
}
