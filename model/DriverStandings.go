package model

// DriverStandings contains the standings for each driver in a race
type DriverStandings struct {
	DriverStandingId int    `json:"driverStandingId" bson:"driverStandingId"`
	RaceId           int    `json:"raceId" bson:"raceId"`
	DriverId         int    `json:"driverId" bson:"driverId"`
	Points           int    `json:"points" bson:"points"`
	Position         int    `json:"position" bson:"position"`
	PositionText     string `json:"positiontText" bson:"positiontText"`
	Wins             int    `json:"wins" bson:"wins"`
}
