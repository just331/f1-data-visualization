package model

import "github.com/go-bongo/bongo"

// Races contains details of every race, including year, date, time, circuit and round
type Races struct {
	bongo.DocumentBase `bson:",inline"`
	RaceID             int    `json:"raceid"`
	Year               int    `json:"year"`
	Round              int    `json:"round"`
	CircuitID          int    `json:"circuitid"`
	Name               string `json:"name"`
	Date               string `json:"date"`
	Time               string `json:"time"`
	URL                string `json:"url"`
}
