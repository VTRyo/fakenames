package jsonp

import (
	"io"
	"net/http"
	"strings"
)

func Parse(url string) string {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	jsonpStr := string(body)
	startIndex := strings.Index(jsonpStr, "(") + 1
	endIndex := strings.Index(jsonpStr, ")")

	jsonStr := jsonpStr[startIndex:endIndex]

	return jsonStr
}
