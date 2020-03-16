package utils

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Addr     string `yaml:"addr"`
		DbName   string `yaml:"dbName"`
	} `yaml:"database"`
}

func GetConfig() Config {
	data, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		panic(err)
	}

	var c Config
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}
	return c
}
