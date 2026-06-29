package utils

import (
	"net/http"
	"strconv"
	"strings"
)

type Cookie struct {
	Name  string
	Value string
}

func SaveHistory(w http.ResponseWriter, history []int) {
	var ids []string
	for _, id := range history {
		ids = append(ids, strconv.Itoa(id))
	}
	historyString := strings.Join(ids, ",")
	cookie := &http.Cookie{
		Name:  "recent",
		Value: historyString,
	}
	http.SetCookie(w, cookie)
}

func GetHistory(r *http.Request) []int {
	cookie, err := r.Cookie("recent")
	if err != nil {
		return []int{}
	}
	ids := strings.Split(cookie.Value, ",")

	var history []int
	for _, idStr := range ids {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}

		history = append(history, id)
	}
	return history
}

func RemoveID(history []int, carID int) []int {
	var result []int
	for _, id := range history {
		if id != carID {
			result = append(result, id)
		}
	}
	return result
}
