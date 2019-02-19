package model

import "github.com/go-bongo/bongo"

// LapTimes contains details the lap times for every race, including driver, lap number, position and time
type LapTimes struct {
	bongo.DocumentBase `bson:",inline"`
	RaceID             int    `json:"raceid"`
	DriverID           int    `json:"driverid"`
	Lap                int    `json:"lap"`
	Position           int    `json:"position"`
	Time               string `json:"time"`
	Milliseconds       int    `json:"milliseconds"`
}
