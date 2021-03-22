package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Answer struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	Code string `bson:"code" json:"code"`
	Answer string `bson:"answer" json:"answer"`
}
