package models

import (
	"errors"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Height    string   `json:"height"`
	Mass      string   `json:"mass"`
	HairColor string   `json:"hair_color"`
	SkinColor string   `json:"skin_color"`
	EyeColor  string   `json:"eye_color"`
	BirthYear string   `json:"birth_year"`
	Gender    string   `json:"gender"`
	Homeworld string   `json:"homeworld"`
	Films     []string `json:"films"`
	Species   []string `json:"species"`
	Vehicles  []string `json:"vehicles"`
	Starships []string `json:"starships"`
	//Created   time.Time `json:"created"`
	//Edited    time.Time `json:"edited"`
	Url string `json:"url"`
}

func NewPerson(name string) (*Person, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	return &Person{Name: name}, nil
}

func (p *Person) String() string {
	return fmt.Sprintf("Person{\nName: %s\n"+
		"Height: %s\n"+
		"Mass: %s\n"+
		"HairColor: %s\n"+
		"SkinColor: %s\n"+
		"EyeColor: %s\n"+
		"BirthYear: %s\n"+
		"Gender: %s\n"+
		"Homeworld: %s\n"+
		"Films: %s\n"+
		"Species: %s\n"+
		"Vehicles: %s\n"+
		"Starships: %s\n"+
		"Url: %s`}",
		p.Name,
		p.Height,
		p.Mass,
		p.HairColor,
		p.SkinColor,
		p.EyeColor,
		p.BirthYear,
		p.Gender,
		p.Homeworld,
		p.Films,
		p.Species,
		p.Vehicles,
		p.Starships,
		p.Url)
}
