package helper

import (
	"strings"
)

var scheduleDays = map[string]int{
	"none":      0,
	"monday":    1,
	"tuesday":   2,
	"wednesday": 4,
	"thursday":  8,
	"friday":    16,
	"saturday":  32,
	"sunday":    64,
	"all":       127,
}

func CalcScheduleDays(days []string) int {
	var result int
	for _, v := range days {
		result += scheduleDays[strings.ToLower(v)]
	}

	return result
}
