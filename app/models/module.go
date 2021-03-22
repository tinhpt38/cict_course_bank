package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Module struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id, omitempty"`
	Name string `bson:"name" json:"name"`
	Questions []string `bson:"questions" json:"questions"`
	IsActive bool `bson:"is_active" json:"is_active"`
	IsExam bool `bson:"is_exam" json:"is_exam"`
	Level Level `bson:"leve" json:"level"`
	ParentID string `bson:"parent_id" json:"parent_id"`
}