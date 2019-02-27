package model

// Circuits contains a list of every formula 1 circuit, including name, location and geographic data.
type Circuits struct {
	CircuitId  float64 `json:"circuitId" bson:"circuitId"`
	CircuitRef string  `json:"circuitRef" bson:"circuitRef"`
	Name       string  `json:"name" bson:"name"`
	Location   string  `json:"location" bson:"location"`
	Country    string  `json:"country" bson:"country"`
	Latitude   float64 `json:"lat" bson:"lat"`
	Longitude  float64 `json:"lng" bson:"lng"`
	Altitude   string  `json:"alt" bson:"alt"`
	URL        string  `json:"url" bson:"url"`
}
