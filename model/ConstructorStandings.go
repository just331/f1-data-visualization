package model

// ConstructorStanding contains the standings for a constructor in a particular race
type ConstructorStanding struct {
	ConstructorStandingID int    `json:"constructorstandingid"`
	RaceID                int    `json:"raceid"`
	ConstructorID         int    `json:"constructorid"`
	Points                int    `json:"points"`
	Position              int    `json:"position"`
	PositionText          string `json:"positionstring"`
	Wins                  int    `json:"wins"`
}
