package main

import (
	"encoding/json"
	"fmt"
	"github.com/thebagchi/openapi3-generator/jsonschema"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("specifications/openapi3.json")
	if nil == err {
		schema := jsonschema.Schema{}
		err := json.Unmarshal(content, &schema)
		if nil == err {
			// Generate Code For OpenAPI Specification
		} else {
			fmt.Println("Error: ", err)
		}
	}
}
