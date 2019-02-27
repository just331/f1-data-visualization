package model

// Constructors contains a list of every constructor team including name and nationality.
type Constructors struct {
	ConstructorId  int    `json:"constructorId" bson:"constructorId"`
	ConstructorRef string `json:"constructorRef" bson:"constructorRef"`
	Name           string `json:"name" bson:"name"`
	Nationality    string `json:"nationality" bson:"nationality"`
	URL            string `json:"url" bson:"url"`
}
