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

	persons, err := getPersons(client)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(persons)
	persons.Count = 222

	fmt.Println("print new count: ", persons.Count)
}

func getPersons(client *http.Client) (models.PersonsResponse, error) {
	url := "https://swapi.dev/api/people/"
	resp, err := client.Get(url)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not read response")
		return models.PersonsResponse{}, err
	}

	personResp := models.PersonsResponse{}
	err = json.Unmarshal(body, &personResp)

	return personResp, nil
}
