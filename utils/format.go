package utils

import (
	"fmt"
	"strconv"
	"strings"
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

func StringToTime(s string) (time.Time, error) {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return time.Time{}, fmt.Errorf("加载时区失败: %w", err)
	}

	timeStamp, err := time.ParseInLocation("2006-01-02 15:04:05", s, location)
	if err != nil {
		return time.Time{}, fmt.Errorf("时间解析失败: %w", err)
	}

	return timeStamp.In(location), nil
}

func StringToSlices(s string) []string {
	return strings.Split(s, " ")
}
