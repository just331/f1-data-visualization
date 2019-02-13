package model

// Status contains a table of status codes and their status
type Status struct {
	StatusID int    `json:"statusid"`
	Status   string `json:"status"`
}
