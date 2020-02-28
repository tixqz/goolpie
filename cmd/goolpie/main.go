package main

import (
	"flag"
	"fmt"

	"github.com/tixqz/goolpie"
)

func main() {
	filePath := flag.String("filepath", "./example/example.yml", "path to your local config")

	flag.Parse()

	conf := goolpie.LoadConfigs(*filePath)
	stubApi := goolpie.NewStubServer(conf)
	fmt.Printf("result: %v", stubApi)
}
