package goolpie

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host      string             `yaml:"host"`
	Port      int                `yaml:"port"`
	BasePath  string             `yaml:"basepath"`
	Endpoints []EndpointSettings `yaml:"paths"`
}

type EndpointSettings struct {
	Endpoint     string `yaml:"endpoint"`
	Method       string `yaml:"method"`
	Request      string `yaml:"request"`
	ResponseCode int    `yaml:"responseCode"`
	ResponseBody string `yaml:"responseBody"`
}

// LoadConfigs returns unmarshaled configuration for server in the case of valid provided yaml config
func LoadConfigs(file []byte) (Config, error) {
	var conf Config

	if err := yaml.Unmarshal(file, &conf); err != nil {
		return conf, fmt.Errorf("Error: %v", err)
	}

	return conf, nil
}
