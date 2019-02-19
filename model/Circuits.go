package model

// Circuit contains a list of every formula 1 circuit, including name, location and geographic data.
type Circuit struct {
	CircuitID  int    `json:"circuitid"`
	CircuitRef string `json:"circuitref"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Country    string `json:"country"`
	Latitude   string `json:"lat"`
	Longitude  string `json:"lng"`
	Altitude   string `json:"alt"`
	URL        string `json:"url"`
}
