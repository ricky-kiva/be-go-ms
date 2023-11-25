package main

import (
	"encoding/json"
	"fmt"
)

func jsonToObjectArray() {
	var jsonString = `[
		{
			"Username": "Rickyslash",
			"Level": 8
		}, {
			"Username": "Geralt of Rivia",
			"Level": 80
		}
	]`

	var data []Character

	var err = json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Player 1 -> %s, Level %d\n", data[0].Name, data[0].Level)
	fmt.Printf("Player 2 -> %s, Level %d\n", data[1].Name, data[1].Level)
}
