package model

// Constructor contains a list of every constructor team including name and nationality.
type Constructor struct {
	ConstructorID  int    `json:"constructorid"`
	ConstructorRef string `json:"constructorref"`
	Name           string `json:"name"`
	Nationality    string `json:"nationality"`
	URL            string `json:"url"`
}
