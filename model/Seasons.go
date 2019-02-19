package model

import "github.com/go-bongo/bongo"

// Seasons contains a list of every season and corresponding Wikipedia link
type Seasons struct {
	bongo.DocumentBase `bson:",inline"`
	Year               int    `json:"year"`
	URL                string `json:"url"`
}
