package model

// ConstructorStandings contains the standings for a constructor in a particular race
type ConstructorStandings struct {
	ConstructorStandingId int    `json:"constructorStandingsId" bson:"constructorStandingsId"`
	RaceId                int    `json:"raceId" bson:"raceId"`
	ConstructorId         int    `json:"constructorId" bson:"constructorId"`
	Points                int    `json:"points" bson:"points"`
	Position              int    `json:"position" bson:"position"`
	PositionText          string `json:"positionText" bson:"positionText"`
	Wins                  int    `json:"wins" bson:"wins"`
}
