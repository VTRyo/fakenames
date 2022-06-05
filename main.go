package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/VTRyo/fakenames/helpers/jsonp"
	"gopkg.in/yaml.v2"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:"name"`
}

func main() {
	url := "https://green.adam.ne.jp/roomazi/cgi-bin/randomname.cgi?n=" + readEnv()
	jsonStr := jsonp.Parse(url)
	var name NameApi
	json.Unmarshal([]byte(jsonStr), &name)

	for i := 0; i < len(name.Name); i++ {
		fmt.Printf("名前: %s(%s),\n", name.Name[i][0], name.Name[i][1])
	}

}

func readEnv() string {
	f, err := ioutil.ReadFile("./env.yml")
	if err != nil {
		fmt.Println(err)
	}

	yamlContent := &YamlContent{}
	err = yaml.Unmarshal(f, &yamlContent)
	if err != nil {
		fmt.Println(err)
	}
	//一旦Countしか返さないのでこれで…
	return yamlContent.Generators.Count
}

type YamlContent struct {
	Generators Generator `yaml:"generators"`
}

type Generator struct {
	Count string `yaml:"count"`
}
