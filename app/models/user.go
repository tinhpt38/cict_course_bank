package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"

)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	FullName string `bson:"full_name" json:"full_name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Email string `bson:"email" json:"email"`
	RoleID string `bson:"role_id" json:"role_id"`
	TokenID string `bson:"token_id" json:"token_id"`
	Avatar string 	`bson:"avatar" json:"avatar"`
	CreateAt string `bson:"create_at" json:"create_at"`
	DeleteAt string `bson:"delete_at" json:"delete_at"`
	UpdateAt string `bson:"update_at" json:"update_at"`
}

func Hash(password string)(string,error){
	bytes, err := bcrypt.GenerateFromPassword( []byte(password),14)
	return string(bytes),err

}

func CheckHashAndPassword(password, hash string) error  {
	return bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
}