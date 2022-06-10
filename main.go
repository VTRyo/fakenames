package main

import (
	"encoding/json"
	"fmt"

	"github.com/VTRyo/fakenames/helpers/config"
	"github.com/VTRyo/fakenames/helpers/jsonp"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:"name"`
}

func main() {
	config := config.FromFile("./config.yml")
	generateCount := config.Generators.Count
	personalities := config.Personalities
	url := "https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi?n=" + generateCount
	jsonStr := jsonp.Parse(url)
	var name NameApi
	json.Unmarshal([]byte(jsonStr), &name)

	for i := 0; i < len(name.Name); i++ {
		fmt.Printf("名前: %s(%s), 性格: %s,\n", name.Name[i][0], name.Name[i][1], personalities[i])
	}

}
