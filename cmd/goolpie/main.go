package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tixqz/goolpie/goolpie"
)

func main() {
	filePath := flag.String("filepath", "mock_settings.yml", "path to your local config")

	flag.Parse()

	conf, err := goolpie.LoadConfigs(openFile(*filePath))
	if err != nil {
		log.Fatalf("")
	}
	stubApi := goolpie.NewStubServer(conf)

	fmt.Printf("result: %v \n", http.ListenAndServe(fmt.Sprintf(":%d", conf.Port), stubApi))
}

func openFile(filepath string) []byte {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Can't open file. Error: %v", err)
	}

	return file
}
