package name

import (
	"encoding/json"

	"github.com/VTRyo/fakenames/helpers/config"
	"github.com/VTRyo/fakenames/helpers/jsonp"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:"name"`
}

func RandomNamesJson() NameApi {
	config := config.FromFile("config.yml")
	generateCount := config.Generators.Count

	url := "https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi?n=" + generateCount
	jsonStr := jsonp.Parse(url)
	var name NameApi
	json.Unmarshal([]byte(jsonStr), &name)
	return name
}
