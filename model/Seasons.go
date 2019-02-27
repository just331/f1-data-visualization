package model

// Seasons contains a list of every season and corresponding Wikipedia link
type Seasons struct {
	Year int    `json:"year" bson:"year"`
	URL  string `json:"url" bson:"url"`
}
