package goolpie

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type StubServer struct {
	config Config
	router *http.ServeMux
}

func NewStubServer(config Config) *StubServer {
	fmt.Printf("Starting server with configs: %v", config)

	s := &StubServer{
		config,
		http.NewServeMux(),
	}
	s.createHandlers()
	return s
}

func (s *StubServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *StubServer) createHandlers() {
	endpoints := s.config.Endpoints
	for endpoint, settings := range endpoints {
		endpoint = s.config.BasePath + endpoint
		log.Printf("Creating handler for endpoint %x.", endpoint)
		for _, setting := range settings {
			s.router.HandleFunc(endpoint, createResponseWithSettings(endpoint, setting))
		}
	}
}

func createResponseWithSettings(path string, setting Endpoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Request body: %x", body)

		if setting.Method != r.Method {
			log.Fatalf("Provided method %x is not expected for this endpoint %x. Expected method: %x", r.Method, path, setting.Method)
		}
		fmt.Printf("Request method: %x", r.Method)

		if setting.Request == string(body) {
			w.WriteHeader(setting.ResponseCode)
			fmt.Fprint(w, setting.ResponseBody)
		} else {
			log.Fatalf("Contract is violated. Expected: %x", setting.Request)
		}

	}
}