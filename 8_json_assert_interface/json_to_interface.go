package main

import (
	"encoding/json"
	"fmt"
)

func jsonToInterface(jsonString string) {
	var data map[string]interface{}

	var jsonData = []byte(jsonString)

	json.Unmarshal(jsonData, &data)

	fmt.Printf("Name: %s, Level: %.f\n", data["Username"], data["Level"])
}
