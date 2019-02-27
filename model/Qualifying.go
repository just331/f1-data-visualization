package model

// Qualifying contains the results of every qualifying session, including the race, driver, constructor, position, and times for Q1, Q2 and Q3.
type Qualifying struct {
	QualifyId     int    `json:"qualifyId" bson:"qualifyId"`
	RaceId        int    `json:"raceId" bson:"raceId"`
	DriverId      int    `json:"driverId" bson:"driverId"`
	ConstructorId int    `json:"constructorId" bson:"constructorId"`
	Number        int    `json:"number" bson:"number"`
	Position      int    `json:"position" bson:"position"`
	Q1            string `json:"q1" bson:"q1"`
	Q2            string `json:"q2" bson:"q2"`
	Q3            string `json:"q3" bson:"q3"`
}
