package model

import "github.com/go-bongo/bongo"

// ConstructorStandings contains the standings for a constructor in a particular race
type ConstructorStandings struct {
	bongo.DocumentBase    `bson:",inline"`
	ConstructorStandingID int    `json:"constructorstandingid"`
	RaceID                int    `json:"raceid"`
	ConstructorID         int    `json:"constructorid"`
	Points                int    `json:"points"`
	Position              int    `json:"position"`
	PositionText          string `json:"positionstring"`
	Wins                  int    `json:"wins"`
}
