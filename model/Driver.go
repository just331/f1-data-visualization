package model

// Driver contains a list of every Formula 1 driver, including full name, dob, and nationality.
type Driver struct {
	DriverID    int    `json:"driverid"`
	DriverRef   string `json:"driverref"`
	Number      int    `json:"number"`
	Code        string `json:"code"`
	Forename    string `json:"forename"`
	Surname     string `json:"surname"`
	DOB         string `json:"dob"`
	Nationality string `json:"nationality"`
	URL         string `json:"url"`
}
