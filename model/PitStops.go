package model

// PitStops contains details of every pitstop in Formula 1, including the time of the pit stop, the duration, the race and driver
type PitStops struct {
	RaceId       int     `json:"raceId" bson:"raceId"`
	DriverId     int     `json:"driverId" bson:"driverId"`
	Stop         int     `json:"stop" bson:"stop"`
	Lap          int     `json:"lap" bson:"lap"`
	Time         string  `json:"time" bson:"time"`
	Duration     float32 `json:"duration" bson:"duration"`
	Milliseconds int     `json:"milliseconds" bson:"milliseconds"`
}
