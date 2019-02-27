package model

// Status contains a table of status codes and their status
type Status struct {
	StatusId int    `json:"statusId" bson:"statusId"`
	Status   string `json:"status" bson:"status"`
}
