package model

// LapTimes contains details the lap times for every race, including driver, lap number, position and time
type LapTimes struct {
	RaceId       int    `json:"raceId" bson:"raceId"`
	DriverId     int    `json:"driverId" bson:"driverId"`
	Lap          int    `json:"lap" bson:"lap"`
	Position     int    `json:"position" bson:"position"`
	Time         string `json:"time" bson:"time"`
	Milliseconds int    `json:"milliseconds" bson:"milliseconds"`
}
