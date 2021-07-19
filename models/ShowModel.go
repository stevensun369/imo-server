package models

type Show struct {
	Title string `json:"title,omitempty" bson:"title,omitempty"`
	TID   string `json:"tid,omitempty" bson:"tid,omitempty"`
}