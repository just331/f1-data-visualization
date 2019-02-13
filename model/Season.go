package model

// Season contains a list of every season and corresponding Wikipedia link
type Season struct {
	Year int    `json:"year"`
	URL  string `json:"url"`
}
