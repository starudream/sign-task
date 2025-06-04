package util

import (
	"time"
)

func GetToday(n ...int) time.Time {
	if len(n) == 0 {
		n = []int{0}
	}
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day()+n[0], 0, 0, 0, 0, now.Location())
}
