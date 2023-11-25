package main

type Character struct {
	Name  string `json:"Username"`
	Level int
}

func main() {
	var jsonString = `{
		"Username": "Rickyslash",
		"Level": 8
	}`

	jsonToStruct(jsonString)
	jsonToInterface(jsonString)
	jsonToObjectArray()
}
