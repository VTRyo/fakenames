package main

import (
	"encoding/json"
	"fmt"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:name`
}

var jsonstr = `
{
	"err": 0,
	"name": [
	[
	"鈴木 真紀",
	"すずき まき",
	"Suzuki Maki"
	],
	[
	"木村 久美子",
	"きむら くみこ",
	"Kimura Kumiko"
	],
	[
	"田島 好治",
	"たしま よしはる",
	"Tasima Yosiharu"
	]
	]
	}
`

func main() {
	var name NameApi
	json.Unmarshal([]byte(jsonstr), &name)
	for i := 0; i < len(name.Name[0]); i++ {
		fmt.Printf("名前: %s\n", name.Name[i][0])
	}
}
