package model

// Qualifying contains the results of every qualifying session, including the race, driver, constructor, position, and times for Q1, Q2 and Q3.
type Qualifying struct {
	QualifyID     int    `json:"qualifyid"`
	RaceID        int    `json:"raceid"`
	DriverID      int    `json:"driverid"`
	ConstructorID int    `json:"constructorid"`
	Number        int    `json:"number"`
	Position      int    `json:"position"`
	Q1            string `json:"q1"`
	Q2            string `json:"q2"`
	Q3            string `json:"q3"`
}
