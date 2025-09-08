package utils

import (
	"fmt"
	"strings"
)

func ParseMsg(msg string) (map[string]string, error) {
	results := map[string]string{}
	msg = strings.TrimSpace(msg)
	if msg == "" {
		return results, fmt.Errorf("msg信息为空")
	}

	pairs := []string{}
	flag := false // 区分postion中的逗号
	start := 0    // pair开始的下标
	for index, character := range msg {
		switch character {
		case '(':
			flag = true
		case ')':
			flag = false
		case ',':
			if !flag {
				pairs = append(pairs, msg[start:index]) // 不包含逗号
				start = index + 1
			}
		}
	}

	// 处理最后一个 pair
	if start < len(msg) {
		pairs = append(pairs, msg[start:])
	}

	for _, pair := range pairs {
		parts := strings.SplitN(pair, ": ", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])   // 去头尾空格
			value := strings.TrimSpace(parts[1]) // 去头尾空格
			results[key] = value
		}
	}

	// fmt.Println("results:", results)

	return results, nil
}

// 更改pos格式从(11, 127, 4014, 1)到111274014001
func ParsePosition(pos string) string {
	pos = strings.TrimPrefix(pos, "(")
	pos = strings.TrimSuffix(pos, ")")

	var position string
	parts := strings.Split(pos, ", ")

	for index, part := range parts {
		if index == 1 || index == 3 { // 第二位和最后一位保证占3bit，不足补0
			part = fmt.Sprintf("%03s", part)
		}
		position += part
	}

	return position
}

// 16条码垛巷道，得到对应LaneNumber
func GetLaneNumber(lane string) []int { // 修改
	var results = []int{}
	for index, char := range lane {
		if char == '1' {
			results = append(results, index+1)
		}
	}
	return results
}
