package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

var UserCollection *mongo.Collection

var ctx = context.TODO()

func InitDB(uri string) error {
	//clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx,clientOptions)
	if err != nil{
		return err
	}
	err = client.Ping(ctx,nil)
	if err != nil{
		return err
	}
	UserCollection = client.Database("cict-quiz-api").Collection("user")
	return nil
}