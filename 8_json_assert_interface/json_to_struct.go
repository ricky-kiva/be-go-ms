package main

import (
	"encoding/json"
	"fmt"
)

type Character struct {
	Name  string `json:"Username"`
	Level int
}

func jsonToStruct() {
	var jsonString = `{
		"Username": "Rickyslash",
		"Level": 8
	}`

	var jsonData = []byte(jsonString)

	var data Character

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error:", err.Error())
	}

	fmt.Printf("Name: %s, Level: %d\n", data.Name, data.Level)
}
