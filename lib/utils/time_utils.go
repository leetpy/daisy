package utils

import "time"

func GetCurrentYear() int {
	now := time.Now()
	return now.Year()
}
