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
	fmt.Printf("Starting server with configs: %v \n", config)

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
	for _, settings := range endpoints {
		endpoint := s.config.BasePath + settings.Endpoint
		log.Printf("Creating handler for endpoint %x.", endpoint)
		s.router.HandleFunc(endpoint, createResponseWithSettings(settings))
	}
}

func createResponseWithSettings(setting EndpointSettings) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		if setting.Method != r.Method {
			log.Printf("Provided method %x is not expected for this endpoint %x. Expected method: %x \n", r.Method, setting.Endpoint, setting.Method)
		}

		if setting.Request == string(body) {
			w.WriteHeader(setting.ResponseCode)
			fmt.Fprint(w, setting.ResponseBody)

		} else {
			fmt.Fprint(w, "Invalid request")
			log.Printf("Contract is violated. Expected: %x \n", setting.Request)
		}

		log.Printf("Request body: %s, Method: %s, Response: %s \n", string(body), r.Method, setting.ResponseBody)

	}
}
