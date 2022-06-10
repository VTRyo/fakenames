package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Content struct {
	Generators    Generator `yaml:"generators"`
	Personalities []string  `yaml:"personalities"`
}

type Generator struct {
	Count string `yaml:"count"`
}

func FromFile(filePath string) *Content {
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}
	content := &Content{}
	err = yaml.Unmarshal(f, &content)
	if err != nil {
		fmt.Println(err)
	}
	return content
}
