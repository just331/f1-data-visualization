package model

// Result contains details for the final results of a driver in a particular race
type Result struct {
	ResultID       int    `json:"resultid"`
	RaceID         int    `json:"raceid"`
	DriverID       int    `json:"driverid"`
	ConstructorID  int    `json:"constructorid"`
	Number         int    `json:"number"`
	Grid           int    `json:"grid"`
	Position       int    `json:"position"`
	PositionText   string `json:"positiontext"`
	PositionOrder  int    `json:"positionorder"`
	Points         int    `json:"points"`
	Laps           int    `json:"laps"`
	Time           string `json:"time"`
	Milliseconds   int    `json:"milliseconds"`
	FastestLap     int    `json:"fastestlap"`
	Rank           int    `json:"rank"`
	FastestLapTime string `json:"fastestlaptime"`
	StatusID       int    `json:"statusid"`
}
