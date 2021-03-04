package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Roles struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`

}