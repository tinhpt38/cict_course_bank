package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	Question string `bson:"question" json:"question"`
	CorrectCodes []string	`bson:"correct_codes" json:"correct_codes"`
	Level Level `bson:"level" json:"level"`
	Answers []Answer `bson:"answers" json:"answers"`	
	ModuleID string	`bson:"module_id" json:"module_id"`
	Content Content  `bson:"content" json:"content"`

}

type Level int
const (
	High = 1
	Medium
	Low
)

type Type int
const (
	Text = 1
	Image
	Voice
)

type Content struct {
	Type Type 	`bson:"type" json:"type"`
	Content string `bson:"content" json:"content"`
}
