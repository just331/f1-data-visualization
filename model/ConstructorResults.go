package model

import "github.com/go-bongo/bongo"

// ConstructorResults contains details of the results for every race, including race, constructor, and awarded points
type ConstructorResults struct {
	bongo.DocumentBase  `bson:",inline"`
	ConstructorResultID int    `json:"constructorresultid"`
	RaceID              int    `json:"raceid"`
	ConstructorID       int    `json:"constructorid"`
	Points              int    `json:"points"`
	Status              string `json:"status"`
}
