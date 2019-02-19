package model

import "github.com/go-bongo/bongo"

// Status contains a table of status codes and their status
type Status struct {
	bongo.DocumentBase `bson:",inline"`
	StatusID           int    `json:"statusid"`
	Status             string `json:"status"`
}
