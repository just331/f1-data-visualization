package model

// LapTimes contains details the lap times for every race, including driver, lap number, position and time
type LapTimes struct {
	RaceID       int    `json:"raceid"`
	DriverID     int    `json:"driverid"`
	Lap          int    `json:"lap"`
	Position     int    `json:"position"`
	Time         string `json:"time"`
	Milliseconds int    `json:"milliseconds"`
}
