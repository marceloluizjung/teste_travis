package main

import (
	"encoding/json"

	"github.com/crgimenes/goconfig"
)

// Declare config struct

type mongoDB struct {
	Host string `cfgDefault:"example.com"`
	Port int    `cfgDefault:"999"`
}

type systemUser struct {
	Name     string `json:"name" cfg:"name"`
	Password string `json:"passwd" cfg:"passwd"`
}

type configTest struct {
	Domain  string
	User    systemUser `json:"user" cfg:"user"`
	MongoDB mongoDB    `json:"mongo" cfg:"mongo"`
}

func main() {

	config := configTest{}

	goconfig.PrefixEnv = "EXAMPLE"

	err := goconfig.Parse(&config)
	if err != nil {
		println(err)
		return
	}

	j, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		println(err)
		return
	}
	println(string(j))
}
