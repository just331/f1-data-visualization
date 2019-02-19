package model

import "github.com/go-bongo/bongo"

// Circuits contains a list of every formula 1 circuit, including name, location and geographic data.
type Circuits struct {
	bongo.DocumentBase `bson:",inline"`
	CircuitID          int    `json:"circuitid"`
	CircuitRef         string `json:"circuitref"`
	Name               string `json:"name"`
	Location           string `json:"location"`
	Country            string `json:"country"`
	Latitude           string `json:"lat"`
	Longitude          string `json:"lng"`
	Altitude           string `json:"alt"`
	URL                string `json:"url"`
}
