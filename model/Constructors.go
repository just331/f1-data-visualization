package model

import "github.com/go-bongo/bongo"

// Constructors contains a list of every constructor team including name and nationality.
type Constructors struct {
	bongo.DocumentBase `bson:",inline"`
	ConstructorID      int    `json:"constructorid"`
	ConstructorRef     string `json:"constructorref"`
	Name               string `json:"name"`
	Nationality        string `json:"nationality"`
	URL                string `json:"url"`
}
