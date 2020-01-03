package goolpie

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"net/http"
)

// StubServer represents simple structure of stub server by one path in config file.
type StubServer struct {
	Path struct {
		Request      string `yaml "request"`
		ResponseCode int    `yaml "responseCode"`
		ResponseBody string `yaml "responseBody"`
	} `yaml "path"`
}

// NewStubServer creates new StubServer instance
func NewStubServer() StubServer {
	return 
}

func (s *StubServer) StartServer()
