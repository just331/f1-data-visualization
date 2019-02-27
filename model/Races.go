package model

// Races contains details of every race, including year, date, time, circuit and round
type Races struct {
	RaceId    int    `json:"raceId" bson:"raceId"`
	Year      int    `json:"year" bson:"year"`
	Round     int    `json:"round" bson:"round"`
	CircuitId int    `json:"circuitId" bson:"circuitId"`
	Name      string `json:"name" bson:"name"`
	Date      string `json:"date" bson:"date"`
	Time      string `json:"time" bson:"time"`
	URL       string `json:"url" bson:"url"`
}
