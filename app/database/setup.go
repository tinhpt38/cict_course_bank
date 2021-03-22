package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var UserCollection *mongo.Collection
var CourseCollection *mongo.Collection
var CategoryCollection *mongo.Collection

var ctx = context.TODO()

func InitDB(uri string) error {
	//clientOptions := options.Client().ApplyURI("mongodb://localhopst:27017")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx,clientOptions)
	if err != nil{
		return err
	}
	err = client.Ping(ctx,nil)
	if err != nil{
		return err
	}
	database := client.Database("cict-quiz-api")
	UserCollection = database.Collection("user")
	CourseCollection = database.Collection("course")
	CategoryCollection = database.Collection("category")
	return nil
}