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
		number, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		history = append(history, number)
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

func SaveCompare(w http.ResponseWriter, compare []int) {
	var ids []string
	for _, id := range compare {
		ids = append(ids, strconv.Itoa(id))
	}

	cookie := &http.Cookie{
		Name:  "compare",
		Value: strings.Join(ids, ","),
		Path:  "/",
	}
	http.SetCookie(w, cookie)
}

func GetCompare(r *http.Request) []int {
	cookie, err := r.Cookie("compare")
	if err != nil {
		return []int{}
	}

	var compare []int
	for _, idStr := range strings.Split(cookie.Value, ",") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			continue
		}
		compare = append(compare, id)
	}
	return compare
}
