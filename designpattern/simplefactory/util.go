package simplefactory

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type defaultStructConfig struct {
	Name string `yaml :"name"`
}

var (
	structConfig defaultStructConfig
)

func init() {
	yamlFile, err := ioutil.ReadFile("struct.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, &structConfig)
	fmt.Println(structConfig)
	if err != nil {
		fmt.Println(err.Error())
	}
}
