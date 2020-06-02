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
		if nil != err {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Id: ", *schema.Id)
			fmt.Println("Title: ", *schema.Title)
			fmt.Println("Description: ", *schema.Description)
			fmt.Println("Schema: ", *schema.Schema)
			fmt.Println("Type: ", *schema.Type)
			if nil != schema.Required {
				for _, item := range *schema.Required {
					fmt.Println(item)
				}
			}
		}
	}
}
