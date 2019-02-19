package model

// DriverStanding contains the standings for each driver in a race
type DriverStanding struct {
	DriverStandingID int    `json:"driverstandingid"`
	RaceID           int    `json:"raceid"`
	DriverID         int    `json:"driverid"`
	Points           int    `json:"points"`
	Position         int    `json:"position"`
	PositionText     string `json:"positionttext"`
	Wins             int    `json:"wins"`
}
