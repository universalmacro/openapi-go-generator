package main

import (
	"fmt"
	"os"

	"github.com/universalmacro/openapi-go-generator/openapi"
	"gopkg.in/yaml.v3"
)

func main() {
	// argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	dat, err := os.ReadFile(argsWithoutProg[0])
	if err != nil {
		panic(err)
	}
	api := openapi.Openapi{}
	err = yaml.Unmarshal(dat, &api)
	if err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%#v", api.File(argsWithoutProg[1]))
	f, err := os.Create(argsWithoutProg[2])
	if err != nil {
		panic(err)
	}
	_, err = f.WriteString(s)
	if err != nil {
		panic(err)
	}
	f.Sync()
}
