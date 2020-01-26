package main

import (
	"fmt"
	"github.com/tixqz/goolpie"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	file, err := ioutil.ReadFile("C:/Users/Mikhail/go/src/github.com/tixqz/goolpie/example/example.yaml")
	if err != nil {
		log.Fatal(err)
	}

	stubApi := &goolpie.StubConfig{}

	err = yaml.Unmarshal(file, &stubApi)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %v", stubApi)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", stubApi.Port), nil))
}
