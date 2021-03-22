package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id, omitempty"`
	Name string `bson:"name" json:"name"`
	ParentID string	`bson:"parent_id" json:"parent_id"`
}
