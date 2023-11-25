package main

import (
	"encoding/json"
	"fmt"
)

func objectArrayToJson() {
	var characterObject = []Character{
		{"Rickyslash", 8},
		{"Geralt of Rivia", 80},
	}

	var jsonData, err = json.Marshal(characterObject)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
