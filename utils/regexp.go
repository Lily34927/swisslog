package utils

import (
	"fmt"
	"regexp"
)

func GetRegExpMap(re *regexp.Regexp, line string) (map[string]string, error) {
	result := map[string]string{}

	matches := re.FindStringSubmatch(line)
	if len(matches) == 0 {
		return result, fmt.Errorf("日志格式不匹配")
	}

	for i, name := range re.SubexpNames() {
		// fmt.Println("i:", i, ", name:", name) // 打印
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	return result, nil
}
