package utils

import (
	"strconv"
	"strings"
)

func ParseIDs(strIDs []string) []int {
	var ids []int
	for _, strID := range strIDs {
		id, err := strconv.Atoi(strID)
		if err == nil {
			ids = append(ids, id)
		}
	}
	return ids
}

func ParseNumber(strNumber string) int {
	number, err := strconv.Atoi(strNumber)
	if err != nil {
		return 0
	}
	return number
}

func ContainsInt(ints []int, n int) bool {
	for _, i := range ints {
		if i == n {
			return true
		}
	}
	return false
}

func ContainsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
