package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"swapi/models"
)

func main() {
	person, err := models.NewPerson("")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(person)

	url := "https://swapi.dev/api/people/1/"

	client := &http.Client{}

	response, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	err = json.Unmarshal(body, &person)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	fmt.Println(person)

}
