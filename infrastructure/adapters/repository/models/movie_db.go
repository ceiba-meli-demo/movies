package models

type MovieDb struct {
	ID       string `bson:"_id,omitempty"`
	Title    string
	Duration int64
	Synopsis string
}
