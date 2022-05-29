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
	jsonStr := string(body)[strings.Index(string(body), "(")+1 : strings.Index(string(body), ")")]
	return jsonStr
}
