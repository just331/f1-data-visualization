package model

import "github.com/go-bongo/bongo"

// PitStops contains details of every pitstop in Formula 1, including the time of the pit stop, the duration, the race and driver
type PitStops struct {
	bongo.DocumentBase `bson:",inline"`
	RaceID             int     `json:"raceid"`
	DriverID           int     `json:"driverid"`
	Stop               int     `json:"stop"`
	Lap                int     `json:"lap"`
	Time               string  `json:"time"`
	Duration           float32 `json:"duration"`
	Milliseconds       int     `json:"milliseconds"`
}
