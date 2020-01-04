package goolpie

import (
	"io"
	"net/http"
)

type StubServer struct {
	name string
	config StubConfig
	http.Handler
}

// StubServer represents simple structure of stub server by one path in config file.
type StubConfig struct {
	Path         string `yaml "path"`
	Method       string `yaml "method"`
	Request      string `yaml "request"`
	ResponseCode int    `yaml "responseCode"`
	ResponseBody string `yaml "responseBody"`
}

func NewStubServer() {
	s := new(StubServer)

	router = http.NewServeMux()
	router.Handle()
}

func 
