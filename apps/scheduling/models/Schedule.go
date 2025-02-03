package models

type Schedule struct {
	ID          string `bson:"_id,omitempty"`
	UserID      string `bson:"user_id"`
	Description string `bson:"description"`
	Date        string `bson:"date"`
	carrier     string `bson:"carrier"`
}
