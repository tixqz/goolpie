package goolpie

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host      string                `yaml:"host"`
	Port      int                   `yaml:"port"`
	BasePath  string                `yaml:"basepath"`
	Endpoints map[string][]Endpoint `yaml:"endpoints"`
}

type Endpoint struct {
	Method       string `yaml:"method"`
	Request      string `yaml:"request"`
	ResponseCode int    `yaml:"responseCode"`
	ResponseBody string `yaml:"responseBody"`
}

func LoadConfigs(filepath string) Config {
	var conf Config

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Can't open file. Error: %w", err)
	}

	err = yaml.Unmarshal(file, &conf)
	if err != nil {
		log.Fatalf("Can't load configs. Error: %w", err)
	}

	return conf
}
