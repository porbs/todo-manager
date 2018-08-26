package models

import "gopkg.in/mgo.v2/bson"

// Note model
type Note struct {
	ID bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string	 `json:"title" bson:"title"`
	Body string		 `json:"body" bson:"body"`
}
