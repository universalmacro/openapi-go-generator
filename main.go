package main

import (
	"fmt"
	"os"

	"github.com/universalmacro/openapigogenerator/openapi"
	"gopkg.in/yaml.v3"
)

type Good struct {
	Good *[]string `yaml:"good"`
}

func main() {
	dat, _ := os.ReadFile("./openapi.yml")
	api := openapi.Openapi{}
	err := yaml.Unmarshal(dat, &api)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", api.File())
	fmt.Printf("%#v", api)

}
