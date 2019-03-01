package model

// Results contains details for the final results of a driver in a particular race
type Results struct {
	ResultId        int     `json:"resultId" bson:"resultId"`
	RaceId          int     `json:"raceId" bson:"raceId"`
	DriverId        int     `json:"driverId" bson:"driverId"`
	ConstructorId   int     `json:"constructorId" bson:"constructorId"`
	Number          int     `json:"number" bson:"number"`
	Grid            int     `json:"grid" bson:"grid"`
	Position        int     `json:"position" bson:"position"`
	PositionText    string  `json:"positionText" bson:"positionText"`
	PositionOrder   int     `json:"positionOrder" bson:"positionOrder"`
	Points          int     `json:"points" bson:"points"`
	Laps            int     `json:"laps" bson:"laps"`
	Time            string  `json:"time" bson:"time"`
	Milliseconds    int     `json:"milliseconds" bson:"milliseconds"`
	FastestLap      int     `json:"fastestLap" bson:"fastestLap"`
	Rank            int     `json:"rank" bson:"rank"`
	FastestLapTime  string  `json:"fastestlLapTime" bson:"fastestLapTime"`
	FastestLapSpeed float64 `json:"fastestLapSpeed" bson:"fastestLapSpeed"`
	StatusId        int     `json:"statusId" bson:"statusId"`
}
