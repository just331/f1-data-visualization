package model

// Drivers contains a list of every Formula 1 driver, including full name, dob, and nationality.
type Drivers struct {
	DriverId    int    `json:"driverId" bson:"driverId"`
	DriverRef   string `json:"driverRef" bson:"driverRef"`
	Number      int    `json:"number" bson:"number"`
	Code        string `json:"code" bson:"code"`
	Forename    string `json:"forename" bson:"forename"`
	Surname     string `json:"surname" bson:"surname"`
	DOB         string `json:"dob" bson:"dob"`
	Nationality string `json:"nationality" bson:"nationality"`
	URL         string `json:"url" bson:"url"`
}
