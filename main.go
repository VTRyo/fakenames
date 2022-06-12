package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/VTRyo/fakenames/apis/name"
	"github.com/VTRyo/fakenames/helpers/config"
)

type NameApi struct {
	Err  int        `json:"err"`
	Name [][]string `json:"name"`
}

func main() {
	config := config.FromFile("./config.yml")
	personalities := config.Personalities

	// personalitiesのリストをランダムに出力したい
	rand.Seed(time.Now().UnixNano())
	personalitiesLen := len(personalities)
	nameList := name.RandomNamesJson()

	for i := 0; i < len(nameList.Name); i++ {
		fmt.Printf("名前: %s(%s), 性格: %s,\n", nameList.Name[i][0], nameList.Name[i][1], personalities[rand.Intn(personalitiesLen)])
	}

}
