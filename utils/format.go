package utils

import (
	"strconv"
	"time"
)

func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func StringToBool(s string) bool {
	if s == "True" || s == "Y" || s == "1" { // 1-SDI
		return true
	}
	return false
}

func StringToTime(s string) (time.Time, error){
	timeStamp, err := time.Parse("2006-01-02 15:04:05", s)
	return timeStamp, err
}