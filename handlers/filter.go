package handlers

import "strconv"

func isSelected(ids []int, id int) bool {
	for _, selectedID := range ids {
		if selectedID == id {
			return true
		}
	}
	return false
}

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
