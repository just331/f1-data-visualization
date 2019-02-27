package model

// ConstructorResults contains details of the results for every race, including race, constructor, and awarded Points
type ConstructorResults struct {
	ConstructorResultId int    `json:"constructorResultsId" bson:"constructorResultsId"`
	RaceId              int    `json:"raceId" bson:"raceId"`
	ConstructorId       int    `json:"constructorId" bson:"constructorId"`
	Points              int    `json:"points" bson:"points"`
	Status              string `json:"status" bson:"status"`
}
