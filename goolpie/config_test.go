package goolpie

import "testing"

func TestLoadConfig(t *testing.T) {
	t.Run("Config is valid YAML", func(t *testing.T) {
		okConfig := `
host: "http://localhost/"
port: 8800
basepath: /v1
paths:
 - endpoint: "/test"
   method: "POST"
   request: "foo"
   responseCode: 200
   responseBody: "bar"
`

		_, err := LoadConfigs([]byte(okConfig))

		if err != nil {
			t.Errorf("Can't load this YAML. %v", err)
		}

	})

	t.Run("Config is invalid YAML", func(t *testing.T) {
		notOkConfig := `
testhost
paths
	ndpoint: "/test"
	- request: "foo"
`

		_, err := LoadConfigs([]byte(notOkConfig))

		if err == nil {
			t.Error("This yaml shouldn't be loaded.")
		}

	})

}
