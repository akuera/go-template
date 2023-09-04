package models

type Domain struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
	// Add more fields
}
