package model

// ConstructorResult contains details of the results for every race, including race, constructor, and awarded points
type ConstructorResult struct {
	ConstructorResultID int    `json:"constructorresultid"`
	RaceID              int    `json:"raceid"`
	ConstructorID       int    `json:"constructorid"`
	Points              int    `json:"points"`
	Status              string `json:"status"`
}
