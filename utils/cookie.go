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
