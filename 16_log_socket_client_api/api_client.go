package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func apiClient() {
	var client = &http.Client{}

	request, err := http.NewRequest("GET", "http://localhost:8080/aot", nil)
	if err != nil {
		fmt.Println("Error making request!")
		return
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error drawing response!")
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading body!")
		return
	}

	fmt.Println(string(body))
}
